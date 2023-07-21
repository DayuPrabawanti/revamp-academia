package main

import (
	"log"
	"os"

	"codeid.revampacademy/config"
	server "codeid.revampacademy/server"

	_ "github.com/lib/pq"
)

func main() {

	log.Println("Startting MiniProjectRevamp rest API")
	log.Println("initializing configuraiton")

	config := config.InitConfig(getConfigFileName())

	log.Println("Inisializing database")
	dbHandler := server.InitDatabase(config)
	// log.Println(dbHandler)

	// //test insert
	// ctx := context.Background() //bikin goroutine

	// queries := repositories.New(dbHandler)
	log.Println("Inisializing HTTP Server")
	httpServer := server.InitHttpServer(config, dbHandler)

	httpServer.GetStart()

}

func getConfigFileName() string {
	env := os.Getenv("ENV")

	if env != "" {
		return "db_revamp" + env
	}

	return "db_revamp"
}
