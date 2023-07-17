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
	salesController  *controller.SalesController
	controllerMockup *controller.ControllerMock1
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	salesRepository := repositories.NewSalesRepository(dbHandler)
	salesService := service.NewSalesService(salesRepository)

	salesController1 := controller.NewSalesRepository(salesService)

	repositoryMock1 := repositories.NewRepositoryMock1(dbHandler)
	serviceMock1 := service.NewServiceMock1(repositoryMock1)

	controllerMock1 := controller.NewRepositoryMock1(serviceMock1)

	router := gin.Default()

	// router endpoint
	router.GET("/SpecialOffer", salesController1.GetListSales)
	router.GET("/CartItem/:id", salesController1.GetListCart_item)
	router.GET("/MockUp/:nama", controllerMock1.GetMockup1)

	return HttpServer{
		config:           config,
		router:           router,
		salesController:  salesController1,
		controllerMockup: controllerMock1,
	}
}

// running gin http server
func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("Http.server_address"))

	if err != nil {
		log.Fatalf("Error while starting HTTp server : %v", err)
	}
}
