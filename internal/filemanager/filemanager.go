package filemanager

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type App struct {
	router         *gin.Engine
	DB             *sql.DB
	ServiceStorage string
}

func (app *App) Initialize(database_path, storage_service_url string) error {
	db, err := sql.Open("postgres", database_path)
	if err != nil {
		return err
	}
	app.DB = db

	app.ServiceStorage = storage_service_url

	app.router = gin.Default()
	return nil
}
