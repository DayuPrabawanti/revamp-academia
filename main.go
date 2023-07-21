package main

import (
	"log"
	"os"

	"codeid.revampacademy/config"
	"codeid.revampacademy/server/bootcampServer"

	_ "github.com/lib/pq"
)

func main() {

	log.Println("starting db_revamp")

	log.Println("Initializing configuration")
	config := config.InitConfig(getConfigFileName())

	log.Println("Initializing database...")
	dbHandler := bootcampServer.InitDatabase(config)
	log.Println(dbHandler)

	log.Println("Initializing HTTP Server!")
	httpServer := bootcampServer.InitHttpServer(config, dbHandler)

	httpServer.Start()
}

func getConfigFileName() string {
	env := os.Getenv("ENV")

	if env != "" {
		return "db_revamp" + env
	}

	return "db_revamp"
}
