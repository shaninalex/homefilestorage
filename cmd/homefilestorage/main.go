package main

import (
	"log"
	"os"
	"strconv"

	"github.com/shaninalex/homefilestorage/api"
	"github.com/shaninalex/homefilestorage/internal/database"
	"github.com/shaninalex/homefilestorage/internal/filemanager"
)

var (
	DATABASE_CONNECTION = os.Getenv("DATABASE_CONNECTION")
	FILEMANAGER_PATH    = os.Getenv("FILEMANAGER_PATH")
	PORT                = os.Getenv("PORT")
)

func main() {
	fm := filemanager.Initialize(FILEMANAGER_PATH)
	database, err := database.CreateConnection(DATABASE_CONNECTION)
	if err != nil {
		log.Println(err)
	}

	port, _ := strconv.Atoi(PORT)
	app, err := api.CreateApi(fm, database)
	if err != nil {
		log.Println(err)
	}
	app.Run(port)
}
