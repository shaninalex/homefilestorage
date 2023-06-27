package main

import (
	"log"

	"github.com/shaninalex/homefilestorage/api"
	"github.com/shaninalex/homefilestorage/internal/filemanager"
)

func main() {

	// TODO: get storage path from env
	fm := filemanager.Initialize("/tmp/")

	app, err := api.CreateApplication(fm)
	if err != nil {
		log.Println(err)
	}
	app.Run(8000)
}
