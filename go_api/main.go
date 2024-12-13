package main

import (
	"log"

	"API/internal"
)

func main() {
	// Log, and get the token from the server
	token, err := internal.PerformLogin("http://localhost:9000/login", "admin", "password")
	if err != nil {
		log.Fatalf("Erreur lors de la connexion : %v", err)
	}

	// Start the server
	go internal.StartServer()

	// Open the browser
	if err := internal.OpenBrowser("http://localhost:8080/welcome"); err != nil {
		log.Fatalf("Erreur lors de l'ouverture du navigateur : %v", err)
	}

	// Keep the program alive to consult the website
	select {}
}
