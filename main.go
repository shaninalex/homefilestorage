package main

import (
	"flag"
	"homestorage/app"
	"log"
)

func main() {

	configPath := flag.String("config", "", "Application toml configuration file path")

	flag.Parse()

	if *configPath == "" {
		log.Fatal("Config file did not provided! Use `-config=<config_path_>` to set configuration file")
	}

	config := app.GetConfig(*configPath)
	app.Run(&config)
}
