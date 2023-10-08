package main

import (
	"os"
	"strconv"

	"github.com/shaninalex/homefilestorage/api"
	"github.com/shaninalex/homefilestorage/pkg/database"
	"github.com/shaninalex/homefilestorage/pkg/filemanager"
)

var (
	DATABASE_CONNECTION = os.Getenv("DATABASE_CONNECTION")
	FILEMANAGER_PATH    = os.Getenv("FILEMANAGER_PATH")
	PORT                = os.Getenv("PORT")
)

func main() {
	fm := filemanager.Initialize(FILEMANAGER_PATH)
	db, err := database.CreateConnection(DATABASE_CONNECTION)
	if err != nil {
		panic(err)
	}
	defer db.DB.Close()

	port, _ := strconv.Atoi(PORT)
	app, err := api.CreateApi(fm, db)
	if err != nil {
		panic(err)
	}
	app.Run(port)
}
