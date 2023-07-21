package hrServer

import (
	"codeid.revampacademy/controllers/hrController"
	"github.com/gin-gonic/gin"
)

func InitRouter(routers *gin.Engine, controllerMgr *hrController.ControllerManager) *gin.Engine {
	// router := gin.Default()

	departmentRoute := routers.Group("/department")
	{
		//router endpoint table department
		departmentRoute.GET("/", controllerMgr.DepartmentController.GetListDepartment)
		departmentRoute.GET("/:id", controllerMgr.DepartmentController.GetDepartment)
		departmentRoute.POST("/", controllerMgr.DepartmentController.CreateDepartment)
		departmentRoute.PUT("/:id", controllerMgr.DepartmentController.UpdateDepartment)
		departmentRoute.DELETE("/:id", controllerMgr.DepartmentController.DeleteDepartment)
	}

	departmentHistoryRoute := routers.Group("/departmenthistory")
	{
		//router endpoint table employee department history
		departmentHistoryRoute.GET("/", controllerMgr.DepartmentHistoryController.GetListDepartmentHistory)
		departmentHistoryRoute.GET("/:id", controllerMgr.DepartmentHistoryController.GetDepartmentHistory)
		departmentHistoryRoute.POST("/", controllerMgr.DepartmentHistoryController.CreateDepartmentHistory)
		departmentHistoryRoute.PUT("/:id", controllerMgr.DepartmentHistoryController.UpdateDepartmentHistory)
		departmentHistoryRoute.DELETE("/:id", controllerMgr.DepartmentHistoryController.DeleteDepartmenHistory)
	}

	clientContractRoute := routers.Group("/clientcontract")
	{
		//router endpoint table client contract
		clientContractRoute.GET("/", controllerMgr.ClientContractController.GetListClientContract)
		clientContractRoute.GET("/:id", controllerMgr.ClientContractController.GetClientContract)
		clientContractRoute.POST("/", controllerMgr.ClientContractController.CreateClientContract)
		clientContractRoute.PUT("/:id", controllerMgr.ClientContractController.UpdateClientContract)
		clientContractRoute.DELETE("/:id", controllerMgr.ClientContractController.DeleteClientContract)
	}

	employeeRoute := routers.Group("/employee")
	{
		//router endpoint table employee
		employeeRoute.GET("/", controllerMgr.EmployeeController.GetListEmployee)
		employeeRoute.GET("/:id", controllerMgr.EmployeeController.GetEmployee)
		employeeRoute.POST("/", controllerMgr.EmployeeController.CreateEmployee)
		employeeRoute.PUT("/:id", controllerMgr.EmployeeController.UpdateEmployee)
		employeeRoute.DELETE("/:id", controllerMgr.EmployeeController.DeleteEmployee)
	}

	payHistoryRoute := routers.Group("/payhistory")
	{
		//router endpoint table employee pay history
		payHistoryRoute.GET("/", controllerMgr.PayHistoryController.GetListPayHistory)
		payHistoryRoute.GET("/:id", controllerMgr.PayHistoryController.GetPayHistory)
		payHistoryRoute.POST("/", controllerMgr.PayHistoryController.CreatePayHistory)
		payHistoryRoute.PUT("/:id", controllerMgr.PayHistoryController.UpdatePayHistory)
		payHistoryRoute.DELETE("/:id", controllerMgr.PayHistoryController.DeletePayHistory)
	}

	talentRoute := routers.Group("/talent")
	{
		//router endpoint table talent
		talentRoute.GET("/", controllerMgr.TalentsMockupController.GetListTalentMockup)
	}

	talentDetailRoute := routers.Group("/talentdetail")
	{
		// router endpoint table talent detail
		talentDetailRoute.GET("/", controllerMgr.TalentsDetailMockupController.GetListTalentDetailMockupDetail)
	}

	return routers
}
