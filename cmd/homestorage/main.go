package main

import (
	"database/sql"
	"errors"
	"flag"
	"log"

	"github.com/shaninalex/homefilestorage/pkg/config"
	"github.com/shaninalex/homefilestorage/pkg/database"
	"github.com/shaninalex/homefilestorage/web"
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

	webapp, err := web.InitializeWebApp(config, repo)
	if err != nil {
		panic(err)
	}

	webapp.Run(config.Web.Port)
}
