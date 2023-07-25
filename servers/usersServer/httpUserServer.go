package usersServer

import (
	"database/sql"
	"log"

	"codeid.revampacademy/controllers/usersController"
	"codeid.revampacademy/repositories/usersRepository"
	"codeid.revampacademy/services/usersService"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpUserServer struct {
	config            *viper.Viper
	router            *gin.Engine
	ControllerManager usersController.ControllerManager
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpUserServer {

	repositoryManager := usersRepository.NewRepositoryManager(dbHandler)
	serviceManager := usersService.NewServiceManager(repositoryManager)
	controllerManager := usersController.NewControllerManager(serviceManager)

	// create object router nly one
	router := gin.Default()
	InitRouter(router, controllerManager)

	return HttpUserServer{
		config:            config,
		router:            router,
		ControllerManager: *controllerManager,
	}
}

// Running gin httpserver
func (hs HttpUserServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("Error While Starting HTTP Server : %v", err)
	}
}
