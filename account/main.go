package main

import (
	"account/app"
	"log"
)

func main() {
	log.Println("Account server is running")

	app := app.App{}
	app.Initialize("amqp://guest:guest@localhost:5672/", "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai")

	// need defer connections here, because in other case - thay close after Initialize end
	defer app.MQConnection.Close()
	defer app.MQChannel.Close()
	app.Run()
}
