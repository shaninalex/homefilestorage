package app

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/shaninalex/homefilestorage/filemanger/app/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	router         *gin.Engine
	DB             *gorm.DB
	MQConnection   *amqp.Connection
	MQChannel      *amqp.Channel
	MQQueue        *amqp.Queue
	ServiceAccount string
	ServiceStorage string
}

func (app *App) Initialize(rabbitmq_connection, database_path, account_service_url, storage_service_url string) error {
	db, err := gorm.Open(postgres.Open(database_path), &gorm.Config{})
	if err != nil {
		return err
	}
	db.AutoMigrate(&database.Folder{}, &database.File{})
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
	app.ServiceAccount = account_service_url
	app.ServiceStorage = storage_service_url

	app.router = gin.Default()
	app.initializeRoutes()
	return nil
}
func (app *App) initializeRoutes() {
	store := persistence.NewInMemoryStore(time.Minute)
	app.router.GET("/_health", Health)

	actionsGroupUser := app.router.Group("/user")
	{
		actionsGroupUser.GET("/:user_id/file/:id", cache.CachePage(store, time.Minute, app.GetFile))
		actionsGroupUser.POST("/:user_id/save-file", app.SaveFile)
	}
}

func (app *App) Run(port string) {
	portInt, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}
	app.router.Run(fmt.Sprintf(":%d", portInt))
}
