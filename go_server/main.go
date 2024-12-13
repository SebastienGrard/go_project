package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"monitoring/internal"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Initiate the SQL database
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal("Erreur d'ouverture de la base de données", err)
	}
	defer db.Close()

	internal.CreateUsersTable(db)
	internal.AddUser(db)

	router := mux.NewRouter()

	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		internal.LoginHandler(w, r, db)
	}).Methods("POST")

	router.HandleFunc("/welcome", internal.WelcomeHandler).Methods("GET")

	fmt.Println("Serveur démarré sur http://localhost:9000")
	log.Fatal(http.ListenAndServe(":9000", router))
}
