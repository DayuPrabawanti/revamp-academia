package server

import (
	"database/sql"
	"log"

	controllers "codeid.revampacademy/controllers/paymentControllers"
	repositories "codeid.revampacademy/repositories/paymentRepositories"
	services "codeid.revampacademy/services/paymentServices"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config             *viper.Viper
	router             *gin.Engine
	ControllersManager controllers.ControllersManager
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	router := gin.Default()

	repositoriesManager := repositories.NewRepositoriesManager(dbHandler)

	servicesManager := services.NewServicesManager(repositoriesManager)

	controllersManager := controllers.NewControllersManager(servicesManager)

	return HttpServer{
		config:             config,
		router:             router,
		ControllersManager: *controllersManager,
	}
}

// Running gin HttpServer
func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))

	if err != nil {
		log.Fatalf("Error while starting HTTP Server: %v", err)
	}
}
