package hrServer

import (
	"database/sql"
	"log"

	"codeid.revampacademy/controllers/hrController"
	"codeid.revampacademy/repositories/hrRepository"
	"codeid.revampacademy/services/hrService"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config            *viper.Viper
	router            *gin.Engine
	ControllerManager hrController.ControllerManager
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	repositoryManager := hrRepository.NewRepositoryManager(dbHandler)
	serviceManager := hrService.NewServiceManager(repositoryManager)
	controllerManager := hrController.NewControllerManager(serviceManager)

	router := gin.Default()
	InitRouter(router, controllerManager)

	return HttpServer{
		config:            config,
		router:            router,
		ControllerManager: *controllerManager,
	}
}

// create method for running gin httpserver
func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http server address"))
	if err != nil {
		log.Fatal("Error while starting HTTP Server : %v", err)
	}
}
