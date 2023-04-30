package main

import (
	"os"

	"github.com/shaninalex/homefilestorage/filemanger/app"
)

func main() {
	app := app.App{}
	err := app.Initialize(
		os.Getenv("RABBITMQ_URL"),
		os.Getenv("DATABASE_URL"),
		os.Getenv("ACCOUNT_SERVICE"),
		os.Getenv("STORAGE_SERVICE"),
	)
	if err != nil {
		panic(err)
	}

	// need defer connections here, because in other case - thay close after Initialize end
	defer app.MQConnection.Close()
	defer app.MQChannel.Close()
	app.Run(os.Getenv("PORT"))
}
