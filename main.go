package main

import (
	"/internal/handlers"
	"PROJET/internal/db"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	db.InitDB()

	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/dashboard", handlers.DashboardHandler).Methods("GET")

	fmt.Println("Serveur démarré sur le port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
