package main

import (
	"log"

	"github.com/vincentyeungg/golang-menu-creator/api"
	"github.com/vincentyeungg/golang-menu-creator/config"
)

func main() {
	// load environment variables for app from current directory
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// start server
	server := api.SetupServer()
	server.Start(config.ServerAddress)
}
