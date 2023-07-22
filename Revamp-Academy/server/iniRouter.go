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
		// initRouter.GET("/cart/:id", controllerMgr.CartItemsController.Getcart_items)

		//api mockup1 localhost:8886/sales/search?nama=Golang localhost:8886/sales/api/search?orderby=online
		initRouter.GET("/search", controllerMgr.ControllerMock.GetMockup1)
		initRouter.GET("/orderby", controllerMgr.ControllerMock.GetListProgram)

		//api mockup2 localhost:8886/sales/cari/2
		initRouter.GET("/cari/:id", controllerMgr.ControllerMock.GetMockupId)
		initRouter.GET("/api/search", controllerMgr.ControllerMock.GetListProgramMock)

		//api mockup3 localhost:8886/sales/applyRegular/1
		initRouter.POST("/save", controllerMgr.ControlMock3.CreateMergeUsers)
		initRouter.GET("/applyRegular/:id", controllerMgr.ControlMock3.GetUsers)
		initRouter.GET("/get", controllerMgr.ControlMock3.GetListGroup)

		// initRouter.GET("/:id", controllerMgr.FintechController.GetPaymentFintech)

		// initRouter.GET("/", controllerMgr.SpecialOfferController.GetListSpecialOffer)
		// initRouter.GET("/:id", controllerMgr.SpecialOfferController.GetSpecial_offer)
		// initRouter.POST("/", controllerMgr.SpecialOfferController.CreateSpecialOffer)

		// initRouter.GET("/edu/:id", controllerMgr.EducationController.GetEducation)
		// initRouter.POST("/", controllerMgr.EducationController.CreateEducation)

	}

	return router

}
