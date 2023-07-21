package curriculumServer

import (
	"database/sql"
	"log"

	controllers "codeid.revampacademy/controllers/curriculumControllers"
	"codeid.revampacademy/controllers/hrController"
	"codeid.revampacademy/controllers/jobhireController"
	"codeid.revampacademy/controllers/paymentControllers"
	repo "codeid.revampacademy/repositories/curriculumRepositories"
	"codeid.revampacademy/repositories/hrRepository"
	"codeid.revampacademy/repositories/jobhireRepositories"
	"codeid.revampacademy/repositories/paymentRepositories"
	services "codeid.revampacademy/services/curriculumServices"
	"codeid.revampacademy/services/hrService"
	"codeid.revampacademy/services/jobhireService"
	"codeid.revampacademy/services/paymentServices"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config *viper.Viper
	router *gin.Engine
	//progentityController *controllers.ProgEntityController
	controllerManager         controllers.ControllerManager
	hrcontrollerManager       hrController.ControllerManager
	paymentControllersManager paymentControllers.ControllersManager
	jobhirecontrollerManager  jobhireController.ControllerManager
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	repositoryManager := repo.NewRepositoryManager(dbHandler)
	servicesManager := services.NewServiceManager(repositoryManager)
	controllerManager := controllers.NewControllerManager(servicesManager)

	hrrepositoryManager := hrRepository.NewRepositoryManager(dbHandler)
	hrserviceManager := hrService.NewServiceManager(hrrepositoryManager)
	hrcontrollerManager := hrController.NewControllerManager(hrserviceManager)

	paymentrepositoriesManager := paymentRepositories.NewRepositoriesManager(dbHandler)
	paymentservicesManager := paymentServices.NewServicesManager(paymentrepositoriesManager)
	paymentcontrollersManager := paymentControllers.NewControllersManager(paymentservicesManager)

	jobhirerepositoryManager := jobhireRepositories.NewRepositoryManager(dbHandler)
	jobhireserviceManager := jobhireService.NewServiceManager(jobhirerepositoryManager)
	jobhirecontrollerManager := jobhireController.NewControllerManager(jobhireserviceManager)

	router := gin.Default()
	InitRouter(router, controllerManager)
	InitRouterHR(router, hrcontrollerManager)
	InitRouterPayment(router, paymentcontrollersManager)
	InitRouterJobhire(router, jobhirecontrollerManager)

	//router.POST("/creategabungmockup", progentityController.CreateGabungbyMockup)
	//router.PUT("/updategabung/:id", progentityController.UpdateGabung)

	return HttpServer{
		config:                    config,
		router:                    router,
		controllerManager:         *controllerManager,
		hrcontrollerManager:       *hrcontrollerManager,
		paymentControllersManager: *paymentcontrollersManager,
		jobhirecontrollerManager:  *jobhirecontrollerManager,
	}
}

func (hs HttpServer) GetStart() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("Eror di: %v", err)
	}
}
