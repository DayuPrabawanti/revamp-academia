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
	controllerMockup *controller.ControllerMock
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	salesRepository := repositories.NewSalesRepository(dbHandler)
	salesService := service.NewSalesService(salesRepository)
	salesController1 := controller.NewSalesRepository(salesService)

	repositoryMock := repositories.NewRepositoryMock(dbHandler)
	serviceMock := service.NewServiceMock(repositoryMock)
	controllerMock := controller.NewRepositoryMock(serviceMock)

	router := gin.Default()

	// router endpoint
	router.GET("/SpecialOffer", salesController1.GetListSales)
	router.GET("/CartItem/:id", salesController1.GetListCart_item)
	router.GET("/api/bootcamp/search", controllerMock.GetMockup1)
	router.GET("/api/bootcamp/progentityid/:id", controllerMock.GetMockupId)

	return HttpServer{
		config:           config,
		router:           router,
		salesController:  salesController1,
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
