package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type App struct {
	router *gin.Engine
}

func CreateApplication() (*App, error) {
	var app App

	app.router = gin.Default()
	return &app, nil
}

func (app *App) initializeRoutes() {
	app.router.GET("/health", app.AppHealth)
}

func (app *App) Run(port int) {
	app.router.Run(fmt.Sprintf(":%d", port))
}

func (app *App) AppHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": true})
}
