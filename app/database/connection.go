package database

import (
	"database/sql"
	"log"
)

type DBConfig struct {
	DB_ENGINE string
	DB_PATH   string
}

// public
func CreateDatabaseConnection(conf *DBConfig) *sql.DB {

	db, err := sql.Open(conf.DB_ENGINE, conf.DB_PATH)
	if err != nil {
		log.Fatalf("Cant connect to database: %s", err)
		return nil
	}

	return db
}
