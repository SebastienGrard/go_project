package internal

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Start the server, and start the HTML function
func StartServer() {
	http.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		if err := RenderWelcomePage(w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Println("Serveur Web démarré sur http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Erreur lors du démarrage du serveur Web : %v", err)
	}
}

// Get the README content, parsing in HTML, and put it in the template
func RenderWelcomePage(w http.ResponseWriter) error {
	readmeContent, err := GetReadmeContent()
	if err != nil {
		return fmt.Errorf("erreur lors de la récupération du contenu : %v", err)
	}

	htmlContent, err := ConvertMarkdownToHTML(readmeContent)
	if err != nil {
		return fmt.Errorf("erreur lors de la conversion du Markdown : %v", err)
	}

	tmpl, err := template.New("welcome").Parse(`
		<!DOCTYPE html>
		<html lang="fr">
		<head>
			<meta charset="UTF-8">
			<title>Welcome</title>
		</head>
		<body>
			<h1>Bienvenue</h1>
			<div>{{.}}</div>
		</body>
		</html>
	`)
	if err != nil {
		return fmt.Errorf("erreur lors du chargement du template : %v", err)
	}

	if err := tmpl.Execute(w, template.HTML(htmlContent)); err != nil {
		return fmt.Errorf("erreur lors de l'exécution du template : %v", err)
	}

	return nil
}
