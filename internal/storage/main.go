package main

import (
	"os"

	"github.com/shaninalex/homefilestorage/storage/app"
)

func main() {
	app := app.App{}
	app.Initialize(
		os.Getenv("RABBITMQ_URL"),
		os.Getenv("STORAGE_PATH"),
	)

	// need defer connections here, because in other case - thay close after Initialize end
	defer app.MQConnection.Close()
	defer app.MQChannel.Close()
	app.Run(os.Getenv("PORT"))
}
