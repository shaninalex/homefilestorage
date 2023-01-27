package app

import (
	"database/sql"
	"homestorage/app/database"
	"homestorage/app/restapi"
)

type App struct {
	DB     *sql.DB
	config *Config
}

var (
	app = &App{}
)

func Run(conf *Config) {
	db_connection := database.CreateDatabaseConnection(conf.Database)
	app.DB = db_connection
	app.config = conf

	// Run http server
	restapi.Server(app.DB, app.config.Application.PORT)
}
