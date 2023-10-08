package main

import (
	"database/sql"
	"log"

	"github.com/shaninalex/homefilestorage/api"
	"github.com/shaninalex/homefilestorage/pkg/config"
	"github.com/shaninalex/homefilestorage/pkg/database"
)

func main() {

	log.Println("Read config...")
	// get config path from comand line arguments:
	// hfsapp --config=/home/user/.local/share/hfsapp/config.toml
	config, err := config.ParseConfig("/home/user/.local/share/hfsapp/config.toml")
	if err != nil {
		panic(err)
	}

	log.Println("Initialize database...")
	db, err := sql.Open("sqlite3", config.DB.Path)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := database.InitSQLiteRepository(db)

	log.Println("Create api...")
	app, err := api.CreateApi(repo, config)
	if err != nil {
		panic(err)
	}

	app.Run(config.GIN.Port)
}
