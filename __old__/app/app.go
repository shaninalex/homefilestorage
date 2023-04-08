package app

import (
	"homestorage/app/database"
	"homestorage/app/restapi"
	"log"
)

type App struct {
	DB     *database.DatabaseRepository
	config *Config
}

var (
	app = &App{}
)

func Run(conf *Config) {
	db, err := database.CreateDatabaseConnection(conf.Database)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Migrate()
	if err != nil {
		log.Fatal(err)
	}
	app.DB = db
	app.config = conf

	// Run http server
	restapi.Server(app.DB, conf.FileStorage.SYSTEM_PATH, app.config.Application.PORT)
}
