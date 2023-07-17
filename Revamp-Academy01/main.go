package main

import (
	"log"
	"os"

	"codeid.revamptwo/config"
	"codeid.revamptwo/servers/userssv"
	_ "github.com/lib/pq"
)

func getConfigFileName() string {
	env := os.Getenv("ENV")

	if env != "" {
		return "db_revamp" + env
	}
	// == file northwind.toml
	return "db_revamp"
}

func main() {
	log.Println("Starting db_revamp restapi")

	log.Println("Initializing configuration")
	config := config.InitConfig(getConfigFileName())

	log.Println("Initializing database")
	dbHandler := userssv.InitDatabase(config)
	// log.Println(dbHandler)

	log.Println("Initializing HTTP Server")
	httpServer := userssv.InitHttpServer(config, dbHandler) 

	httpServer.Start()
}