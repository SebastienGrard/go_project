package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

var jwtKey = []byte("secret_key")

func main() {
	// Démarrage de l'API dans une goroutine
	go startAPI()

	// Délai pour s'assurer que l'API est en ligne avant de tenter de se connecter
	fmt.Println("⏳ Démarrage de l'API...")
	time.Sleep(3 * time.Second) // Attendre 3 secondes pour permettre au serveur de démarrer

	// Simuler une connexion au serveur
	username := "admin"
	password := "password"
	token, err := login(username, password)
	if err != nil {
		log.Fatalf("❌ Erreur de connexion : %v", err)
	}
	fmt.Println("✅ Connexion réussie, Token :", token)

	fmt.Println("Ouvrir la page web")
	openBrowser("http://localhost:9000/welcome")

	// Ne pas fermer le programme immédiatement, attendre une entrée de l'utilisateur
	// Cela permet au serveur de continuer à fonctionner
	select {}
}

// Démarre l'API en local
func startAPI() {
	router := mux.NewRouter()

	router.HandleFunc("/login", LoginHandler).Methods("POST")
	router.HandleFunc("/welcome", WelcomeHandler).Methods("GET")

	fmt.Println("📡 Serveur API démarré sur le port 9000")
	err := http.ListenAndServe(":9000", router)
	if err != nil {
		log.Fatalf("Erreur lors du démarrage du serveur API : %v", err)
	}
}

// Handler de connexion
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Requête invalide"})
		return
	}

	// Simulation de la validation des identifiants
	if loginRequest.Username == "admin" && loginRequest.Password == "password" {
		token, err := GenerateJWT(loginRequest.Username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Erreur lors de la génération du token"})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"token": token})
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Identifiants incorrects"})
	}
}

// Handler de la page de bienvenue
func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html lang="fr">
		<head>
			<meta charset="UTF-8">
			<title>Bienvenue</title>
		</head>
		<body>
			<h1>Bienvenue sur le tableau de bord !</h1>
			<p>Vous êtes maintenant connecté avec succès.</p>
		</body>
		</html>
	`)
}

// Fonction de connexion à l'API
func login(username, password string) (string, error) {
	loginURL := "http://localhost:9000/login" // Modifié pour correspondre au port correct
	loginPayload := map[string]string{
		"username": username,
		"password": password,
	}
	payloadBytes, err := json.Marshal(loginPayload)
	if err != nil {
		return "", fmt.Errorf("erreur de création de la requête JSON : %v", err)
	}

	resp, err := http.Post(loginURL, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", fmt.Errorf("erreur de connexion au serveur : %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("erreur de lecture de la réponse : %v", err)
	}

	var responseMap map[string]string
	err = json.Unmarshal(body, &responseMap)
	if err != nil {
		return "", fmt.Errorf("erreur de décodage JSON : %v", err)
	}

	if token, exists := responseMap["token"]; exists {
		return token, nil
	}
	return "", fmt.Errorf("identifiants incorrects")
}

func GenerateJWT(username string) (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Expiration de 24 heures
		Subject:   username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// Fonction pour ouvrir la page web dans le navigateur par défaut
func openBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	}
	if err != nil {
		log.Printf("Erreur lors de l'ouverture du navigateur : %v", err)
	}
}
