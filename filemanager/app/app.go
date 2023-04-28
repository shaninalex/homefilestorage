package app

import (
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

type App struct {
	router       *gin.Engine
	DB           *gorm.DB
	MQConnection *amqp.Connection
	MQChannel    *amqp.Channel
	MQQueue      *amqp.Queue
}

func (app *App) Initialize(rabbitmq_connection, database_path string) error { return nil }
func (app *App) initializeRoutes()                                          {}
func (app *App) Run(port string)                                            {}
