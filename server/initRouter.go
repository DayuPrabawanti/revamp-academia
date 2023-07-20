package server

import (
	"codeid.revampacademy/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(controllerMgr *controller.ControllerMock) {
	router := gin.Default()

	cateRouter := router.Group("/Sales/")

	{
		// router endpoint
		cateRouter.GET("/api/bootcamp/search", controllerMgr.GetMockup1)
		cateRouter.GET("/api/bootcamp/progentityid/:id", controllerMgr.GetMockupId)
		cateRouter.GET("/contoh", controllerMgr.GetListProgram)
	}
}
