package app

import (
	"account/app/models"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gorm.io/driver/postgres"

	amqp "github.com/rabbitmq/amqp091-go"
)

// TODO: add rabbitmq exchange instead. For example:
// - "loggin" - for sending logging messages
// - "manage" - background tasks like update avatar link from filestorage service, or schedule payments etc...
type App struct {
	router       *gin.Engine
	DB           *gorm.DB
	MQConnection *amqp.Connection
	MQChannel    *amqp.Channel
	MQQueue      *amqp.Queue
}

func (app *App) Initialize(rabbitmq_connection, database_path string) error {
	db, err := gorm.Open(postgres.Open(database_path), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&models.User{})
	app.DB = db

	// Connect with RabbitMQ
	mq_connection, err := connectToRabbitMQ(rabbitmq_connection)
	if err != nil {
		return err
	}
	ch, err := mq_connection.Channel()
	failOnError(err, "Failed to open a channel")
	q, err := ch.QueueDeclare(
		"register", // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	failOnError(err, "Failed to declare a queue")

	app.MQConnection = mq_connection
	app.MQChannel = ch
	app.MQQueue = &q

	app.router = gin.Default()
	app.initializeRoutes()

	return nil
}

func (app *App) initializeRoutes() {
	store := persistence.NewInMemoryStore(time.Minute)
	app.router.GET("/ping", Ping)
	app.router.POST("/account", app.CreateUser)
	app.router.GET("/account/:id", cache.CachePage(store, time.Minute, app.GetUser))
	app.router.PATCH("/account/:id", app.UpdateUser)
	app.router.DELETE("/account/:id", app.UpdateUser)
}

func (app *App) Run(port string) {
	portInt, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}
	app.router.Run(fmt.Sprintf(":%d", portInt))
}
