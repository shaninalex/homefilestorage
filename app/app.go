package app

import (
	"homestorage/app/database"

	"gorm.io/gorm"
)

type App struct {
	DB *gorm.DB
}

func Run(conf *Config) {
	database.CreateDatabaseConnection(conf.Database)
}
