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
		return "revamp_academy" + env
	}
	// == file revamp_academy.toml
	return "revamp_academy"
}

func main() {
	log.Println("Starting revamp_academy restapi")

	log.Println("Initializing configuration")
	config := config.InitConfig(getConfigFileName())
	log.Println("Initializing database")
	dbHandler := servers.InitDatabase(config)

	log.Println("Initializing HTTP Server")
	httpUserServer := usersServer.InitHttpServer(config, dbHandler)

	 httpUserServer.Start()

	//  test
}