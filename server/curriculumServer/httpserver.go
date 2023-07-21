package curriculumServer

import (
	"database/sql"
	"log"

	controllers "codeid.revampacademy/controllers/curriculumControllers"
	"codeid.revampacademy/controllers/hrController"
	repo "codeid.revampacademy/repositories/curriculumRepositories"
	"codeid.revampacademy/repositories/hrRepository"
	services "codeid.revampacademy/services/curriculumServices"
	"codeid.revampacademy/services/hrService"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config *viper.Viper
	router *gin.Engine
	//progentityController *controllers.ProgEntityController
	controllerManager   controllers.ControllerManager
	hrcontrollerManager hrController.ControllerManager
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	repositoryManager := repo.NewRepositoryManager(dbHandler)
	servicesManager := services.NewServiceManager(repositoryManager)
	controllerManager := controllers.NewControllerManager(servicesManager)

	hrrepositoryManager := hrRepository.NewRepositoryManager(dbHandler)
	hrserviceManager := hrService.NewServiceManager(hrrepositoryManager)
	hrcontrollerManager := hrController.NewControllerManager(hrserviceManager)

	router := gin.Default()
	InitRouter(router, controllerManager)
	InitRouterHR(router, hrcontrollerManager)

	//router.POST("/creategabungmockup", progentityController.CreateGabungbyMockup)
	//router.PUT("/updategabung/:id", progentityController.UpdateGabung)

	return HttpServer{
		config:              config,
		router:              router,
		controllerManager:   *controllerManager,
		hrcontrollerManager: *hrcontrollerManager,
	}
}

func (hs HttpServer) GetStart() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("Eror di: %v", err)
	}
}
