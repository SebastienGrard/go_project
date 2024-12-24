package dbconfig

import (
	"crud_project/helper"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Postgres goalng driver
	"github.com/rs/zerolog/log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbName   = "test"
)

func DatabaseConnection() *sql.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := sql.Open("postgres", sqlInfo)
	helper.PanicError(err)

	err = db.Ping()
	helper.PanicError(err)

	log.Info().Msg("Connected to database !")

	return db
}
