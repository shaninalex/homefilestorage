package main

import (
	"database/sql"
	"errors"
	"flag"
	"log"

	"github.com/shaninalex/homefilestorage/api"
	"github.com/shaninalex/homefilestorage/pkg/config"
	"github.com/shaninalex/homefilestorage/pkg/database"
)

func main() {

	configPath := flag.String("config", "", "Config path")
	flag.Parse()

	if *configPath == "" {
		panic(errors.New("config path is required"))
	}

	log.Println("Read config...")
	config, err := config.ParseConfig(*configPath)
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

	log.Println("Start server...")
	app.Run(config.GIN.Port)
}
