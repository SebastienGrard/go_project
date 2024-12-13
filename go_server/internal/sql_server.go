package internal

import (
	"database/sql"
	"log"
	"sync"
)

// Mutex pour la gestion concurrente de la base de données
var DbMutex sync.Mutex

func CreateUsersTable(db *sql.DB) {
	DbMutex.Lock()
	defer DbMutex.Unlock()

	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Erreur de création de la table users", err)
	}
}

// Add a user in the database. Check if he's already in the database, if it's not, then add in it.
func AddUser(db *sql.DB) {
	DbMutex.Lock()
	defer DbMutex.Unlock()

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", "admin").Scan(&count)
	if err != nil {
		log.Println("Erreur lors de la vérification de l'utilisateur :", err)
		return
	}

	if count > 0 {
		log.Println("L'utilisateur 'admin' existe déjà.")
		return
	}

	_, err = db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", "admin", "password")
	if err != nil {
		log.Println("Erreur lors de l'ajout de l'utilisateur de test :", err)
	} else {
		log.Println("Utilisateur ajouté avec succès")
	}
}
