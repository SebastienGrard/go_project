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

	"github.com/yuin/goldmark" // Bibliothèque pour le format Markdown
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
	url := "http://localhost:9000/login"

	loginRequest := LoginRequest{
		Username: "admin",
		Password: "password",
	}

	jsonData, err := json.Marshal(loginRequest)
	if err != nil {
		log.Fatalf("Erreur lors de l'encodage JSON: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Erreur lors de l'envoi de la requête: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Erreur lors de la lecture de la réponse: %v", err)
	}

	var loginResponse LoginResponse
	if err := json.Unmarshal(body, &loginResponse); err != nil {
		log.Fatalf("Erreur lors du décodage de la réponse JSON: %v", err)
	}

	if loginResponse.Token != "" {
		fmt.Println("Connexion réussie, token :", loginResponse.Token)

		go func() {
			http.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
				// Récupération du contenu du README depuis GitHub
				readmeContent, err := getReadmeContent()
				if err != nil {
					http.Error(w, "Erreur lors de la récupération du contenu", http.StatusInternalServerError)
					return
				}

				// Conversion du Markdown en HTML
				htmlContent, err := convertMarkdownToHTML(readmeContent)
				if err != nil {
					http.Error(w, "Erreur lors de la conversion du Markdown", http.StatusInternalServerError)
					return
				}

				// Charger le template et injecter le contenu HTML
				tmpl, err := template.New("welcome").Parse(`
					<!DOCTYPE html>
					<html lang="fr">
					<head>
						<meta charset="UTF-8">
						<title>Welcome</title>
					</head>
					<body>
						<h1></h1>
						<div>{{.}}</div>
					</body>
					</html>
				`)
				if err != nil {
					http.Error(w, "Erreur lors du chargement du template", http.StatusInternalServerError)
					return
				}

				err = tmpl.Execute(w, template.HTML(htmlContent)) // Attention: template.HTML pour éviter l'échappement
				if err != nil {
					http.Error(w, "Erreur lors de l'exécution du template", http.StatusInternalServerError)
					return
				}
			})

			log.Println("Serveur Web démarré sur http://localhost:8080")
			err := http.ListenAndServe(":8080", nil)
			if err != nil {
				log.Fatalf("Erreur lors du démarrage du serveur Web : %v", err)
			}
		}()

		err := openBrowser("http://localhost:8080/welcome")
		if err != nil {
			log.Fatalf("Erreur lors de l'ouverture du navigateur : %v", err)
		}

		select {}
	} else {
		fmt.Println("Échec de la récupération du token")
	}
}

// Fonction pour récupérer le contenu du README depuis GitHub
func getReadmeContent() (string, error) {
	url := "https://raw.githubusercontent.com/SebastienGrard/go_project/testing/go_api/README.md"
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("erreur lors de la récupération du README: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("erreur HTTP: %v", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("erreur lors de la lecture du README: %v", err)
	}

	return string(body), nil
}

// Fonction pour convertir du Markdown en HTML
func convertMarkdownToHTML(markdownContent string) (string, error) {
	var buf bytes.Buffer
	md := goldmark.New()
	if err := md.Convert([]byte(markdownContent), &buf); err != nil {
		return "", fmt.Errorf("erreur lors de la conversion du Markdown: %v", err)
	}
	return buf.String(), nil
}

// Fonction pour ouvrir un navigateur
func openBrowser(url string) error {
	var cmd *exec.Cmd
	switch os := runtime.GOOS; os {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	default:
		return fmt.Errorf("système d'exploitation non supporté : %v", os)
	}
	return cmd.Start()
}
