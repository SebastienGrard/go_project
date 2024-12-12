package main

import (
	"fmt"
	"log"

	"API/internal"
)

func main() {
	// Étape 1 : Se connecter et obtenir le token
	token, err := internal.PerformLogin("http://localhost:9000/login", "admin", "password")
	if err != nil {
		log.Fatalf("Erreur lors de la connexion : %v", err)
	}

	fmt.Println("Connexion réussie, token :", token)

	// Étape 2 : Lancer le serveur
	go internal.StartServer()

	// Étape 3 : Ouvrir le navigateur
	if err := internal.OpenBrowser("http://localhost:8080/welcome"); err != nil {
		log.Fatalf("Erreur lors de l'ouverture du navigateur : %v", err)
	}

	// Maintenir le programme actif
	select {}
}
