package server

import (
	salescontrollers "codeid.revampacademy/controllers/salesControllers"
	"github.com/gin-gonic/gin"
)

func InitRouterSales(router *gin.Engine, controllerMgr *salescontrollers.ControllerManager) *gin.Engine {

	cateRouter := router.Group("/sales")

	{
		// router endpoint
		//api mockup1
		cateRouter.GET("/search", controllerMgr.ControllerMock.GetMockup1)
		cateRouter.GET("/orderby", controllerMgr.ControllerMock.GetListProgram)
		//api mockup2
		cateRouter.GET("/cari/:id", controllerMgr.ControllerMock.GetMockupId)
		cateRouter.GET("/api/search", controllerMgr.ControllerMock.GetListProgramMock)
		//api mockup3
		cateRouter.POST("/save", controllerMgr.ControlMock3.CreateMergeUsers)
		cateRouter.GET("/applyRegular/:id", controllerMgr.ControlMock3.GetUsers)

	}
	return router
}
