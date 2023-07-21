package server

import (
	"database/sql"
	"log"

	"codeid.revampacademy/controller/salesController"
	"codeid.revampacademy/repositories/salesRepositories"
	"codeid.revampacademy/services/salesServices"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config *viper.Viper
	router *gin.Engine

	ControllerManager salesController.ControllerManager
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {

	repositoryManager := salesRepositories.NewRepositoryManager(dbHandler)

	serviceManager := salesServices.NewServiceManager(repositoryManager)

	controllerManager := salesController.NewControllerManager(serviceManager)

	// router := gin.Default()
	router := InitRouter(controllerManager)

	return HttpServer{
		config:            config,
		router:            router,
		ControllerManager: *controllerManager,
	}
}

// running gin httpserver
func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("Error while starting HTTP Server : %v", err)
	}
}
