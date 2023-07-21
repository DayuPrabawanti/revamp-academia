package server

import (
	"codeid.revampacademy/controller/salesController"
	"github.com/gin-gonic/gin"
)

func InitRouter(controllerMgr *salesController.ControllerManager) *gin.Engine {
	router := gin.Default()

	initRouter := router.Group("/sales")
	{
		// router endpoint

		// initRouter.GET("/", controllerMgr.CartItemsController.GetListCartItems)
		initRouter.GET("/cart/:id", controllerMgr.CartItemsController.Getcart_items)

		initRouter.GET("/search", controllerMgr.ControllerMock.GetMockup1)

		// initRouter.GET("/:id", controllerMgr.FintechController.GetPaymentFintech)

		initRouter.GET("/", controllerMgr.SpecialOfferController.GetListSpecialOffer)
		initRouter.GET("/:id", controllerMgr.SpecialOfferController.GetSpecial_offer)
		// initRouter.POST("/", controllerMgr.SpecialOfferController.CreateSpecialOffer)

		initRouter.GET("/edu/:id", controllerMgr.EducationController.GetEducation)
		// initRouter.POST("/", controllerMgr.EducationController.CreateEducation)

	}

	return router

}
