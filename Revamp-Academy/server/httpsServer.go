package server

import (
	"database/sql"
	"log"

	"codeid.revampacademy/controller"
	"codeid.revampacademy/repositories"
	"codeid.revampacademy/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config                 *viper.Viper
	router                 *gin.Engine
	CartItemsController    *controller.CartItemsController
	controllerMockup       *controller.ControllerMock
	specialOfferController *controller.SpecialOfferController
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {

	cartItemsRepository := repositories.NewCartItemsRepository(dbHandler)

	cartItemsService := services.NewCartItemsRepository(cartItemsRepository)

	cartItemsController := controller.NewCartItemsController(cartItemsService)

	repositoryMock := repositories.NewRepositoryMock(dbHandler)

	serviceMock := services.NewServiceMock(repositoryMock)

	controllerMock := controller.NewRepositoryMock(serviceMock)

	specialOfferRepository := repositories.NewSpecialOfferRepository(dbHandler)

	specialOfferService := services.NewSpecialOfferService(specialOfferRepository)

	specialOfferController := controller.NewSpecialController(specialOfferService)

	router := gin.Default()
	// router endpoint

	//cartitems
	// router.GET("/cartItems", cartItemsController.GetListCartItems)
	router.GET("/cartItems/:id", cartItemsController.Getcart_items)
	// router.POST("/cartItems", cartItemsController.CreatecartItems)

	// router.PUT("/cartItems/:id", cartItemsController.UpdatecartItems)
	// router.DELETE("/cartItems/:id", cartItemsController.DeletecartItems)

	//curriculum.prog_entity
	router.GET("/api/bootcamp/search", controllerMock.GetMockup1)

	//specialoffer
	router.GET("/specialOffer", specialOfferController.GetListSpecialOffer)
	router.GET("/specialOffer/:id", specialOfferController.GetSpecial_offer)
	router.POST("/specialOffer", specialOfferController.CreateSpecialOffer)

	return HttpServer{
		config:                 config,
		router:                 router,
		CartItemsController:    cartItemsController,
		controllerMockup:       controllerMock,
		specialOfferController: specialOfferController,
	}
}

// running gin httpserver
func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("Error while starting HTTP Server : %v", err)
	}
}
