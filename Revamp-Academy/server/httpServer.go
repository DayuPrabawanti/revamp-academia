package server

import (
	// "codeid.revampacademy/config"
	"database/sql"
	"log"

	// "codeid.revampacademy/controller"
	"codeid.revampacademy/controller/jobhireController"
	"codeid.revampacademy/service/jobhireService"

	// "codeid.revampacademy/repositories"
	"codeid.revampacademy/repositories/jobhireRepositories"
	// "codeid.revampacademy/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config *viper.Viper
	router *gin.Engine
	// categoryController *controller.CategoryController
	// masterController *controller.MasterController
	jobHireController *jobhireController.ControllerManager
}

func InitHttpServer(cfg *viper.Viper, dbHandler *sql.DB) HttpServer {

	//set router from gin first
	repositoryManager := jobhireRepositories.NewRepositoryManager(dbHandler)
	serviceManager := jobhireService.NewServiceManager(repositoryManager)
	controllerManager := jobhireController.NewControllerManager(serviceManager)

	router := InitRouter(controllerManager)

	return HttpServer{
		config:            cfg,
		router:            router,
		jobHireController: controllerManager,
	}
}

// running for gin server
func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("Error While starting HTTP Server : %v", err)
	}
}
