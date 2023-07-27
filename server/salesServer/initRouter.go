package salesserver

import (
	salescontrollers "codeid.revampacademy/controller/salesControllers"
	"github.com/gin-gonic/gin"
)

func InitRouter(controllerMgr *salescontrollers.ControllerManager) *gin.Engine {
	router := gin.Default()

	salesRouter := router.Group("/sales")

	{
		// router endpoint
		//api mockup1
		salesRouter.GET("/search", controllerMgr.ControllerMock.GetMockup1)
		salesRouter.GET("/orderby", controllerMgr.ControllerMock.GetListProgram)
		//api mockup2
		salesRouter.GET("/:id", controllerMgr.ControllerMock.GetMockupId)
		salesRouter.GET("/api/search", controllerMgr.ControllerMock.GetListProgramMock)
		//api mockup3
		salesRouter.GET("/applyRegular/:id", controllerMgr.ControlMock3.GetUsers)
		salesRouter.GET("/get", controllerMgr.ControlMock3.GetListGroup)
		salesRouter.GET("/getBatch", controllerMgr.ControlMock3.GetListApplyProgress)
		salesRouter.POST("/media", controllerMgr.ControlMock3.CreateMedian)
		salesRouter.POST("/users", controllerMgr.ControlMock3.CreateUsers)
		salesRouter.POST("/education", controllerMgr.ControlMock3.CreateEducation)
		salesRouter.POST("/save", controllerMgr.ControlMock3.CreateMergeUsers)

		//api mockup6
		salesRouter.GET("/cart/:id", controllerMgr.ControlMock6.GetUserIdShoppingCart1)
		salesRouter.GET("/discount", controllerMgr.ControlMock6.ListSalesOrderControl)
		salesRouter.GET("/fintech/verify/account", controllerMgr.ControlMock6.GetAccountNumbersControl)

		//api mockup7
		salesRouter.GET("/cart/userEntity/:id", controllerMgr.ControlMock7.GetUsersIdShoopingCart2Control)
		salesRouter.POST("/OrderDetail", controllerMgr.ControlMock7.CreateOrderDetailControl)
		salesRouter.DELETE("/Cancel/:id", controllerMgr.ControlMock7.CancelOrderDetailControl)
		salesRouter.GET("/accounts", controllerMgr.ControlMock7.GetAccountNumbersMock7Control)

		//api mockup8
		salesRouter.GET("/summaryOrder", controllerMgr.ControlMockup8.GetIdSummaryOrderMock8Control)
	}
	return router
}
