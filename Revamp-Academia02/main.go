package main

import (
	"log"
	"os"

	// "time"

	"codeid.revamptwo/config"

	"codeid.revamptwo/server"

	_ "github.com/lib/pq"
)

func getConfigFileName() string {
	env := os.Getenv("ENV")

	if env != "" {
		return "revamp" + env
	}
	// == file revamp.toml
	return "revamp"
}

func main() {
	log.Println("Starting revamp RestAPI")

	log.Println("Initializing Configuration")
	config := config.InitConfig(getConfigFileName())

	log.Println("Initializing Database")
	dbHandler := server.InitDatabase(config)

	log.Println("initializing HTTP Server")
	httpServer := server.InitHttpServer(config, dbHandler)
	httpServer.Start()
}
