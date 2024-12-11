package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	connStr := "user=postgres password=yourpassword dbname=monitoring sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Erreur de connexion à la base de données:", err)
	}
}
