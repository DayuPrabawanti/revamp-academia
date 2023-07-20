package curriculumServer

import (
	"database/sql"
	"log"

	controllers "codeid.revampacademy/controllers/curriculumControllers"
	repo "codeid.revampacademy/repositories/curriculum"
	services "codeid.revampacademy/services/curriculumServices"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config *viper.Viper
	router *gin.Engine
	//progentityController *controllers.ProgEntityController
	controllerManager controllers.ControllerManager
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	repositoryManager := repo.NewRepositoryManager(dbHandler)
	servicesManager := services.NewServiceManager(repositoryManager)
	controllerManager := controllers.NewControllerManager(servicesManager)

	router := gin.Default()
	InitRouter(router, controllerManager)

	//router.POST("/creategabungmockup", progentityController.CreateGabungbyMockup)
	//router.PUT("/updategabung/:id", progentityController.UpdateGabung)

	return HttpServer{
		config:            config,
		router:            router,
		controllerManager: *controllerManager,
	}
}

func (hs HttpServer) GetStart() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("Eror di: %v", err)
	}
}
