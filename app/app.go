package app

import (
	"homestorage/app/database"

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

}
