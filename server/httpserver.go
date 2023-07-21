package server

import (
	"database/sql"
	"log"

	"codeid.revampacademy/controllers/bootcampController"
	controllers "codeid.revampacademy/controllers/curriculumControllers"
	"codeid.revampacademy/controllers/hrController"
	"codeid.revampacademy/controllers/jobhireController"
	"codeid.revampacademy/controllers/paymentControllers"
	"codeid.revampacademy/controllers/salesControllers"
	"codeid.revampacademy/controllers/usersController"
	"codeid.revampacademy/repositories/bootcampRepository"
	repo "codeid.revampacademy/repositories/curriculumRepositories"
	"codeid.revampacademy/repositories/hrRepository"
	"codeid.revampacademy/repositories/jobhireRepositories"
	"codeid.revampacademy/repositories/paymentRepositories"
	"codeid.revampacademy/repositories/salesRepositories"
	"codeid.revampacademy/repositories/usersRepository"
	"codeid.revampacademy/services/bootcampService"
	services "codeid.revampacademy/services/curriculumServices"
	"codeid.revampacademy/services/hrService"
	"codeid.revampacademy/services/jobhireService"
	"codeid.revampacademy/services/paymentServices"
	"codeid.revampacademy/services/salesService"
	"codeid.revampacademy/services/usersService"
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
	bootcampcontrollerManager bootcampController.ControllerManager
	salescontrollerManager    salesControllers.ControllerManager
	usercontrollerManager     usersController.ControllerManager
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

	bootcamprepositoryManager := bootcampRepository.NewRepositoryManager(dbHandler)
	bootcampserviceManager := bootcampService.NewServiceManager(bootcamprepositoryManager)
	bootcampcontrollerManager := bootcampController.NewControllerManager(bootcampserviceManager)

	salesrepositoryManager := salesRepositories.NewRepositoryManager(dbHandler)
	salesserviceManager := salesService.NewServiceManager(salesrepositoryManager)
	salescontrollerManager := salesControllers.NewControllerManager(salesserviceManager)

	userrepositoryManager := usersRepository.NewRepositoryManager(dbHandler)
	userserviceManager := usersService.NewServiceManager(userrepositoryManager)
	usercontrollerManager := usersController.NewControllerManager(userserviceManager)

	router := gin.Default()
	InitRouterCurriculum(router, controllerManager)
	InitRouterHR(router, hrcontrollerManager)
	InitRouterPayment(router, paymentcontrollersManager)
	InitRouterJobhire(router, jobhirecontrollerManager)
	InitRouterBootcamp(router, bootcampcontrollerManager)
	InitRouterSales(router, salescontrollerManager)
	InitRouterUser(router, usercontrollerManager)

	//router.POST("/creategabungmockup", progentityController.CreateGabungbyMockup)
	//router.PUT("/updategabung/:id", progentityController.UpdateGabung)

	return HttpServer{
		config:                    config,
		router:                    router,
		controllerManager:         *controllerManager,
		hrcontrollerManager:       *hrcontrollerManager,
		paymentControllersManager: *paymentcontrollersManager,
		jobhirecontrollerManager:  *jobhirecontrollerManager,
		bootcampcontrollerManager: *bootcampcontrollerManager,
		salescontrollerManager:    *salescontrollerManager,
		usercontrollerManager:     *usercontrollerManager,
	}
}

func (hs HttpServer) GetStart() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("Eror di: %v", err)
	}
}
