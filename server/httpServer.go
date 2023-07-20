package server

import (
	"database/sql"
	"log"

	"codeid.revampacademy/controller"
	"codeid.revampacademy/repositories"
	"codeid.revampacademy/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config           *viper.Viper
	router           *gin.Engine
	controllerMockup *controller.ControllerMock
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	repositoryMock := repositories.NewRepositoryMock(dbHandler)
	serviceMock := service.NewServiceMock(repositoryMock)
	controllerMock := controller.NewRepositoryMock(serviceMock)

	router := gin.Default()

	return HttpServer{
		config:           config,
		router:           router,
		controllerMockup: controllerMock,
	}
}

// running gin http server
func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("Http.server_address"))

	if err != nil {
		log.Fatalf("Error while starting HTTp server : %v", err)
	}
}
