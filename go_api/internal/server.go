package internal

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		// Récupération et rendu de la page welcome
		if err := RenderWelcomePage(w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Println("Serveur Web démarré sur http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Erreur lors du démarrage du serveur Web : %v", err)
	}
}

// renderWelcomePage génère et affiche le contenu Markdown dans la page d'accueil
func RenderWelcomePage(w http.ResponseWriter) error {
	// Récupération du contenu du README
	readmeContent, err := GetReadmeContent()
	if err != nil {
		return fmt.Errorf("erreur lors de la récupération du contenu : %v", err)
	}

	// Conversion du Markdown en HTML
	htmlContent, err := ConvertMarkdownToHTML(readmeContent)
	if err != nil {
		return fmt.Errorf("erreur lors de la conversion du Markdown : %v", err)
	}

	// Chargement et exécution du template
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

	// Injection et affichage du contenu HTML dans le template
	if err := tmpl.Execute(w, template.HTML(htmlContent)); err != nil {
		return fmt.Errorf("erreur lors de l'exécution du template : %v", err)
	}

	return nil
}
