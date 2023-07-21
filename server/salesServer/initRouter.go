package salesserver

import (
	salescontrollers "codeid.revampacademy/controller/salesControllers"
	"github.com/gin-gonic/gin"
)

func InitRouter(controllerMgr *salescontrollers.ControllerManager) *gin.Engine {
	router := gin.Default()

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
		cateRouter.GET("/get", controllerMgr.ControlMock3.GetListGroup)

	}
	return router
}
