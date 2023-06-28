package main

import (
	"log"
	"os"

	"github.com/shaninalex/homefilestorage/api"
	"github.com/shaninalex/homefilestorage/internal/database"
	"github.com/shaninalex/homefilestorage/internal/filemanager"
)

var (
	DATABASE_CONNECTION = os.Getenv("DATABASE_CONNECTION")
	FILEMANAGER_PATH    = os.Getenv("FILEMANAGER_PATH")
)

func main() {

	// TODO: get storage path from env
	fm := filemanager.Initialize(FILEMANAGER_PATH)
	database, err := database.CreateConnection(DATABASE_CONNECTION)
	if err != nil {
		log.Println(err)
	}

	app, err := api.CreateApplication(fm, database)
	if err != nil {
		log.Println(err)
	}
	app.Run(8000)
}
