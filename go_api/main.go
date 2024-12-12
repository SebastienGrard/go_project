package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

// Structure de la requête de connexion
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Clé secrète utilisée pour signer et vérifier le JWT
var jwtKey = []byte("secret_key")

// Cache pour la page Welcome
type Cache struct {
	data      string
	createdAt time.Time
	sync.Mutex
}

var welcomeCache = Cache{}

func main() {
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/login", LoginHandler).Methods("POST")
	router.HandleFunc("/welcome", WelcomeHandler).Methods("GET")

	fmt.Println("Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Handler de connexion
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginRequest LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Requête invalide"})
		return
	}

	// Simuler la validation de l'utilisateur (à remplacer par une BDD)
	if loginRequest.Username == "admin" && loginRequest.Password == "password" {
		jwtToken, err := GenerateJWT(loginRequest.Username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Erreur interne du serveur"})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"token": jwtToken})
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(map[string]string{"error": "Nom d'utilisateur ou mot de passe incorrect"})
}

// Handler de la page de bienvenue
func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Token manquant"})
		return
	}

	// Supprimer le préfixe "Bearer " de l'en-tête Authorization
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("méthode de signature inattendue")
		}
		return jwtKey, nil
	})
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Token invalide"})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if _, ok := claims["sub"].(string); ok {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")

			// Gérer le cache (valide pendant 2 minutes)
			welcomeCache.Lock()
			defer welcomeCache.Unlock()

			if time.Since(welcomeCache.createdAt) > 2*time.Minute || welcomeCache.data == "" {
				fmt.Println("Mise à jour du cache de la page Welcome")
				http.ServeFile(w, r, "./templates/welcome.html")
				cacheContent := loadHTMLFile("./templates/welcome.html")
				welcomeCache.data = cacheContent
				welcomeCache.createdAt = time.Now()
			} else {
				fmt.Println("Utilisation du cache pour la page Welcome")
				w.Write([]byte(welcomeCache.data))
			}
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

// Fonction pour charger un fichier HTML dans une variable
func loadHTMLFile(filePath string) string {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Erreur de chargement du fichier %s: %v", filePath, err)
		return ""
	}
	return string(content)
}
