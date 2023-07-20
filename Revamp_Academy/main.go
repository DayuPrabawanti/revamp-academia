package main

import (
	"log"
	"os"

	"codeid.revampacademy/config"
	"codeid.revampacademy/servers"
	server "codeid.revampacademy/servers"

	_ "github.com/lib/pq"
)

func main() {

	log.Println("starting db_revamp")

	log.Println("Initializing configuration")
	config := config.InitConfig(getConfigFileName())

	log.Println("Initializing database...")
	dbHandler := server.InitDatabase(config)
	log.Println(dbHandler)

	log.Println("Initializing HTTP Server")
	httpServer := servers.InitHttpServer(config, dbHandler)

	httpServer.Start()

}

func getConfigFileName() string {
	env := os.Getenv("ENV")

	if env != "" {
		return "db_revamp" + env
	}

	return "db_revamp"
}
