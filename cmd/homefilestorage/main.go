package main

import (
	"log"

	"github.com/shaninalex/homefilestorage/api"
)

func main() {
	app, err := api.CreateApplication()
	if err != nil {
		log.Println(err)
	}
	app.Run(8000)
}
