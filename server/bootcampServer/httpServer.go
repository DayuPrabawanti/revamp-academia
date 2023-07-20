package bootcampServer

import (
	"database/sql"
	"log"

	"codeid.revampacademy/controllers/bootcampController"
	"codeid.revampacademy/repositories/bootcampRepository"
	"codeid.revampacademy/services/bootcampService"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config            *viper.Viper
	router            *gin.Engine
	ControllerManager bootcampController.ControllerManager
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {

	repositoryManager := bootcampRepository.NewRepositoryManager(dbHandler)
	serviceManager := bootcampService.NewServiceManager(repositoryManager)
	controllerManager := bootcampController.NewControllerManager(serviceManager)
	//create object router only one
	router := gin.Default()

	InitRouter(router, controllerManager)

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
