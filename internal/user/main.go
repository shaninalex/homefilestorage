package main

import (
	"os"

	"github.com/shaninalex/homefilestorage/user/app"
)

func main() {
	app := app.App{}
	app.Initialize(
		os.Getenv("KRATOS_PATH"),
	)

	app.Run(os.Getenv("PORT"))
}
