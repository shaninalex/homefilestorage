package main

import (
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
	KRATOS_PATH         = os.Getenv("KRATOS_PATH")
)

func main() {
	fm := filemanager.Initialize(FILEMANAGER_PATH)
	db, err := database.CreateConnection(DATABASE_CONNECTION)
	if err != nil {
		panic(err)
	}
	defer db.DB.Close()

	port, _ := strconv.Atoi(PORT)
	app, err := api.CreateApi(fm, db, KRATOS_PATH)
	if err != nil {
		panic(err)
	}
	app.Run(port)
}
