package main

import (
	"account/app"
	"log"
)

func main() {
	log.Println("Account server is running")

	app := app.App{}
	app.Initialize("amqp://guest:guest@localhost:5672/")

	// need defer connections here, because in other case - thay close after Initialize end
	defer app.MQConnection.Close()
	defer app.MQChannel.Close()
	app.Run()
}
