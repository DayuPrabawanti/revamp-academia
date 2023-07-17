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
	config          *viper.Viper
	router          *gin.Engine
	salesController *controller.SalesController
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	salesRepository := repositories.NewSalesRepository(dbHandler)
	salesService := service.NewSalesService(salesRepository)

	salesController := controller.NewSalesRepository(salesService)

	router := gin.Default()

	// router endpoint
	router.GET("/SpecialOffer", salesController.GetListSales)

	return HttpServer{
		config:          config,
		router:          router,
		salesController: salesController,
	}
}

// running gin http server
func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("Http.server_address"))

	if err != nil {
		log.Fatalf("Error while starting HTTp server : %v", err)
	}
}
