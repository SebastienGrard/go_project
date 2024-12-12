package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

// Structure de la requête de connexion
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Clé secrète utilisée pour signer et vérifier le JWT
var jwtKey = []byte("secret_key")

// Mutex pour la gestion concurrente de la base de données
var dbMutex sync.Mutex

func main() {
	// Initialisation de la base de données SQLite
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal("Erreur d'ouverture de la base de données", err)
	}
	defer db.Close()

	createUsersTable(db)

	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		LoginHandler(w, r, db)
	}).Methods("POST")

	router.HandleFunc("/welcome", WelcomeHandler).Methods("GET")

	fmt.Println("Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Création de la table "users" si elle n'existe pas déjà
func createUsersTable(db *sql.DB) {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Erreur de création de la table users", err)
	}
}

// Handler de connexion
func LoginHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var loginRequest LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Requête invalide"})
		return
	}

	dbMutex.Lock()
	defer dbMutex.Unlock()

	var storedPassword string
	err := db.QueryRow("SELECT password FROM users WHERE username = ?", loginRequest.Username).Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Nom d'utilisateur ou mot de passe incorrect"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Erreur interne du serveur"})
		}
		return
	}

	if loginRequest.Password != storedPassword {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Nom d'utilisateur ou mot de passe incorrect"})
		return
	}

	jwtToken, err := GenerateJWT(loginRequest.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Erreur interne du serveur"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": jwtToken})
}

// Handler de la page de bienvenue
func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Token manquant"})
		return
	}

	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Méthode de signature inattendue")
		}
		return jwtKey, nil
	})
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Token invalide"})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if username, ok := claims["sub"].(string); ok {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"message": "Bienvenue " + username})
			return
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(map[string]string{"error": "Token non valide"})
}

// Fonction de génération de JWT
func GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Expire dans 24 heures
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
