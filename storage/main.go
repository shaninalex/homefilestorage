package main

import (
	"os"
)

func main() {
	app := App{}
	app.Initialize(
		os.Getenv("RABBITMQ_URL"),
	)

	// need defer connections here, because in other case - thay close after Initialize end
	defer app.MQConnection.Close()
	defer app.MQChannel.Close()
	app.Run(os.Getenv("PORT"))
}
