package app

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	amqp "github.com/rabbitmq/amqp091-go"
)

type App struct {
	router         *gin.Engine
	DB             *sql.DB
	MQConnection   *amqp.Connection
	MQChannel      *amqp.Channel
	MQQueue        *amqp.Queue
	ServiceAccount string
	ServiceStorage string
}

func (app *App) Initialize(rabbitmq_connection, database_path, storage_service_url string) error {
	db, err := sql.Open("postgres", database_path)
	if err != nil {
		return err
	}
	app.DB = db

	// Connect with RabbitMQ
	mq_connection, err := connectToRabbitMQ(rabbitmq_connection)
	if err != nil {
		return err
	}
	ch, err := mq_connection.Channel()
	failOnError(err, "Failed to open a channel")
	q, err := ch.QueueDeclare(
		"storage", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	app.MQConnection = mq_connection
	app.MQChannel = ch
	app.MQQueue = &q
	app.ServiceStorage = storage_service_url

	app.router = gin.Default()
	app.initializeRoutes()
	return nil
}

func (app *App) initializeRoutes() {
	// store := persistence.NewInMemoryStore(time.Minute)
	app.router.GET("/_health", Health)

	actionsGroupUser := app.router.Group("/files")
	{
		actionsGroupUser.GET("/list", app.GetFiles)
		// actionsGroupUser.GET("/:id", app.SingleFile)
		// actionsGroupUser.POST("/save", app.SaveFile)
		// actionsGroupUser.GET("/data/:file_id", app.FileData)
	}
}

func (app *App) Run(port string) {
	portInt, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}
	app.router.Run(fmt.Sprintf(":%d", portInt))
}
