package app

import (
	"homestorage/app/database"
	"homestorage/app/restapi"

	"gorm.io/gorm"
)

type App struct {
	DB *gorm.DB
}

var (
	app = &App{}
)

func Run(conf *Config) {
	db_connection := database.CreateDatabaseConnection(conf.Database)
	app.DB = db_connection

	// Run http server
	// TODO: it should be running in goroutine with other app.
	restapi.Server(app.DB)
}
