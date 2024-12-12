package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

// Structure de la requête de connexion
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Structure de la réponse avec le token
type LoginResponse struct {
	Token string `json:"token"`
}

func main() {
	// URL du serveur de connexion
	url := "http://localhost:9000/login"

	// Création de la requête de connexion
	loginRequest := LoginRequest{
		Username: "admin",
		Password: "password",
	}

	// Encodage de la requête en JSON
	jsonData, err := json.Marshal(loginRequest)
	if err != nil {
		log.Fatalf("Erreur lors de l'encodage JSON: %v", err)
	}

	// Envoi de la requête POST pour se connecter
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Erreur lors de l'envoi de la requête: %v", err)
	}
	defer resp.Body.Close()

	// Lire la réponse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Erreur lors de la lecture de la réponse: %v", err)
	}

	// Vérifier si la réponse contient le token
	var loginResponse LoginResponse
	if err := json.Unmarshal(body, &loginResponse); err != nil {
		log.Fatalf("Erreur lors du décodage de la réponse JSON: %v", err)
	}

	// Vérifier si le token est présent
	if loginResponse.Token != "" {
		fmt.Println("OK:", loginResponse.Token)

		// Démarrer un serveur HTTP pour afficher la page welcome.html
		go func() {
			http.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
				// Charger et exécuter le template HTML
				tmpl, err := template.ParseFiles("templates/welcome.html")
				if err != nil {
					http.Error(w, "Erreur lors du chargement de la page", http.StatusInternalServerError)
					return
				}
				// Rendre le template
				err = tmpl.Execute(w, nil)
				if err != nil {
					http.Error(w, "Erreur lors de l'exécution du template", http.StatusInternalServerError)
				}
			})

			// Démarrer le serveur HTTP pour afficher la page
			log.Println("Serveur Web démarré sur http://localhost:8080")
			err := http.ListenAndServe(":8080", nil)
			if err != nil {
				log.Fatalf("Erreur lors du démarrage du serveur Web : %v", err)
			}
		}()

		// Ouvrir le navigateur pour afficher la page welcome
		err := openBrowser("http://localhost:8080/welcome")
		if err != nil {
			log.Fatalf("Erreur lors de l'ouverture du navigateur : %v", err)
		}

		// Le serveur principal continue de fonctionner
		fmt.Println("Serveur de connexion en cours d'exécution...")
		select {} // Maintenir le programme en vie
	} else {
		fmt.Println("Échec de la récupération du token")
	}
}

// Fonction pour ouvrir un navigateur avec une URL donnée
func openBrowser(url string) error {
	var cmd *exec.Cmd
	// Selon le système d'exploitation
	switch os := runtime.GOOS; os {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		return fmt.Errorf("système d'exploitation non supporté : %v", os)
	}

	// Exécuter la commande
	return cmd.Start()
}
