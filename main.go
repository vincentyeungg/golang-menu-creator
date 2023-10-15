package main

import (
	"database/sql"
	"log"

	"github.com/vincentyeungg/golang-menu-creator/api"
	"github.com/vincentyeungg/golang-menu-creator/config"
	db "github.com/vincentyeungg/golang-menu-creator/db/sqlc"

	_ "github.com/lib/pq"
)

func main() {
	// load environment variables for app from current directory
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// setup db connection
	conn, err := sql.Open(config.DbDriver, config.DbSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.New(conn)
	
	// start server
	server := api.SetupServer(store)
	server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
