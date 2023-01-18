package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBConfig struct {
	DEBUG      bool
	DEBUG_NAME string
	HOST       string
	PORT       int
	NAME       string
	USER       string
	PASS       string
}

// public
func CreateDatabaseConnection(conf *DBConfig) *gorm.DB {

	if conf.DEBUG {

		db_name := "database.db"
		if conf.DEBUG_NAME != "" {
			db_name = conf.DEBUG_NAME
		}

		fmt.Printf("Using SQLite3: %s\n", db_name)

		db, err := gorm.Open(sqlite.Open(db_name), &gorm.Config{})

		if err != nil {
			log.Fatalf("Cant connect to database: %s", err)
			return nil
		}
		return db

	} else {
		fmt.Println(buildConnectionUrl(conf))
		db, err := gorm.Open(postgres.Open(buildConnectionUrl(conf)), &gorm.Config{SkipDefaultTransaction: true})
		if err != nil {
			log.Fatalf("Cant connect to database: %s", err)
			return nil
		}

		return db
	}
}

// Internal functions
func buildConnectionUrl(config *DBConfig) string {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.HOST,
		config.USER,
		config.PASS,
		config.NAME,
		config.PORT,
	)
	return dsn
}
