package main

import (
	"log"
	"os"

	"codeid.revampacademy/config"
	"codeid.revampacademy/servers"
	"codeid.revampacademy/servers/usersServer"
	_ "github.com/lib/pq"
)

func getConfigFileName() string {
	env := os.Getenv("ENV")

	if env != "" {
		return "revamp_academia" + env
	}
	// == file revamp_academia.toml
	return "revamp_academia"
}

func main() {
	log.Println("Starting revamp_academia restapi")

	log.Println("Initializing configuration")
	config := config.InitConfig(getConfigFileName())
	log.Println("Initializing database")
	dbHandler := servers.InitDatabase(config)

	log.Println("Initializing HTTP Server")
	httpUserServer := usersServer.InitHttpServer(config, dbHandler)

	httpUserServer.Start()

	//  test
}
