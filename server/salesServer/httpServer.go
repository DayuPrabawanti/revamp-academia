package salesserver

import (
	"database/sql"
	"log"

	salescontrollers "codeid.revampacademy/controller/salesControllers"
	salesrepositories "codeid.revampacademy/repositories/salesRepositories"
	salesservice "codeid.revampacademy/service/salesService"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config            *viper.Viper
	router            *gin.Engine
	ControllerManager salescontrollers.ControllerManager
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	repositoryMock := salesrepositories.NewRepositoryManager(dbHandler)
	serviceMock := salesservice.NewServiceManager(repositoryMock)
	controllerMock := salescontrollers.NewControllerManager(serviceMock)

	router := InitRouter(controllerMock)

	return HttpServer{
		config:            config,
		router:            router,
		ControllerManager: *controllerMock,
	}
}

// running gin http server
func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("Http.server_address"))

	if err != nil {
		log.Fatalf("Error while starting HTTp server : %v", err)
	}
}
