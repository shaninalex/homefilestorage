package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

type App struct {
	router       *gin.Engine
	storage      *FileStorage
	MQConnection *amqp.Connection
	MQChannel    *amqp.Channel
	MQQueue      *amqp.Queue
}

func (app *App) Initialize(rabbitmq_connection, storage_path string) error {
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

	app.router = gin.Default()
	app.initializeRoutes()

	storage, err := CreateFileStorage(storage_path)
	if err != nil {
		return err
	}
	app.storage = storage

	return nil
}

func (app *App) initializeRoutes() {
	app.router.POST("/save", app.SaveFile)
	app.router.GET("/files/:y/:m/:d/:filename", app.RetrieveFile)
}

func (app *App) Run(port string) {
	portInt, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}
	app.router.Run(fmt.Sprintf(":%d", portInt))
}

func (app *App) SaveFile(c *gin.Context) {

	file, handler, err := c.Request.FormFile("file")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	defer file.Close()

	dFile := &File{
		Size: int(handler.Size),
		Name: handler.Filename,
	}

	dFile, err = app.storage.SaveFileToStorage(file, handler, dFile)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	message, err := json.Marshal(gin.H{"message": "New file uploaded", "data": dFile})
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	Publish(app.MQQueue.Name, string(message), app.MQChannel, app.MQQueue)
	c.JSON(http.StatusOK, dFile)
}

func (app *App) RetrieveFile(c *gin.Context) {
	y := c.Params.ByName("y")
	m := c.Params.ByName("m")
	d := c.Params.ByName("d")
	filename := c.Params.ByName("filename")
	file_path := fmt.Sprintf("%s/%s/%s/%s/%s", app.storage.storage, y, m, d, filename)
	c.FileAttachment(file_path, filename)
}
