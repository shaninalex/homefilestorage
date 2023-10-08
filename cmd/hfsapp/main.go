package main

import (
	"database/sql"
	"os"
	"strconv"

	"github.com/shaninalex/homefilestorage/api"
	"github.com/shaninalex/homefilestorage/pkg/database"
)

var (
	DATABASE_CONNECTION = os.Getenv("DATABASE_CONNECTION")
	PORT                = os.Getenv("PORT")
)

func main() {
	db, err := sql.Open("sqlite3", DATABASE_CONNECTION)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := database.InitSQLiteRepository(db)
	port, _ := strconv.Atoi(PORT)
	app, err := api.CreateApi(repo)
	if err != nil {
		panic(err)
	}
	app.Run(port)
}
