package server

import (
	"codeid.revampacademy/controllers/hrController"
	"github.com/gin-gonic/gin"
)

func InitRouterHR(routers *gin.Engine, controllerMgr *hrController.ControllerManager) *gin.Engine {

	// Register routes here
	departmentHistoryRoute := routers.Group("/departmenthistory")
	{
		// routers endpoint/url http category
		departmentHistoryRoute.GET("", controllerMgr.DepartmentHistoryController.GetListDepartmentHistory)
		departmentHistoryRoute.GET("/:id", controllerMgr.DepartmentHistoryController.GetDepartmentHistory)
		departmentHistoryRoute.POST("", controllerMgr.DepartmentHistoryController.CreateDepartmentHistory)
		departmentHistoryRoute.PUT("/:id", controllerMgr.DepartmentHistoryController.UpdateDepartmentHistory)
		departmentHistoryRoute.DELETE("/:id", controllerMgr.DepartmentHistoryController.DeleteDepartmenHistory)
	}

	departmentRoute := routers.Group("/department")
	{
		// routers endpoint/url http category
		departmentRoute.GET("", controllerMgr.DepartmentController.GetListDepartment)
		departmentRoute.GET("/:id", controllerMgr.DepartmentController.GetDepartment)
		departmentRoute.POST("", controllerMgr.DepartmentController.CreateDepartment)
		departmentRoute.PUT("/:id", controllerMgr.DepartmentController.UpdateDepartment)
		departmentRoute.DELETE("/:id", controllerMgr.DepartmentController.DeleteDepartment)
	}

	employeeRouter := routers.Group("/employee")
	{
		// routers endpoint/url http category
		employeeRouter.GET("", controllerMgr.EmployeeController.GetListEmployee)
		employeeRouter.GET("/:id", controllerMgr.EmployeeController.GetEmployee)
		employeeRouter.POST("", controllerMgr.EmployeeController.CreateEmployee)
		employeeRouter.PUT("/:id", controllerMgr.EmployeeController.UpdateEmployee)
		employeeRouter.DELETE("/:id", controllerMgr.EmployeeController.DeleteEmployee)
		employeeRouter.POST("/users", controllerMgr.EmployeeController.CreateUser)
	}

	payHistoryRoute := routers.Group("/payhistory")
	{
		// routers endpoint/url http category
		payHistoryRoute.GET("", controllerMgr.PayHistoryController.GetListPayHistory)
		payHistoryRoute.GET("/:id", controllerMgr.PayHistoryController.GetPayHistory)
		payHistoryRoute.POST("", controllerMgr.PayHistoryController.CreatePayHistory)
		payHistoryRoute.PUT("/:id", controllerMgr.PayHistoryController.UpdatePayHistory)
		payHistoryRoute.DELETE("/:id", controllerMgr.PayHistoryController.DeletePayHistory)
	}

	talentDetailRoute := routers.Group("/api/hr")
	{
		// routers endpoint/url http category
		talentDetailRoute.GET("/talentdetail", controllerMgr.TalentsDetailMockupController.GetListTalentDetailMockupDetail)
		talentDetailRoute.GET("/talentdetail/:id", controllerMgr.TalentsDetailMockupController.GetTalentDetail)
		talentDetailRoute.GET("/talentdetail/search", controllerMgr.TalentsDetailMockupController.SearchTalentDetail)
		talentDetailRoute.GET("/talentdetail/switchStatus", controllerMgr.TalentsDetailMockupController.UpdateSwitch)

	}

	talentRoute := routers.Group("/api/hr")
	{
		// routers endpoint/url http category
		talentRoute.GET("/talent", controllerMgr.TalentsMockupController.GetListTalentMockup)
		talentRoute.GET("/talent/search", controllerMgr.TalentsMockupController.SearchTalent)
		talentRoute.GET("/talent/paging", controllerMgr.TalentsMockupController.PagingTalent)
		talentRoute.GET("/talent/view", controllerMgr.TalentsMockupController.GetBatch)
		talentRoute.GET("/talent/batchid", controllerMgr.TalentsMockupController.UpdateBatch)

	}

	clientContractRoute := routers.Group("/api/hr")
	{
		// routers endpoint/url http category
		clientContractRoute.GET("clientcontract", controllerMgr.ClientContractController.GetListClientContract)
		clientContractRoute.GET("/clientcontract/:id", controllerMgr.ClientContractController.GetClientContract)
		clientContractRoute.POST("clientcontract/create", controllerMgr.ClientContractController.CreateClientContract)
		clientContractRoute.PUT("/clientcontract/:id", controllerMgr.ClientContractController.UpdateClientContract)
		clientContractRoute.DELETE("clientcontract/:id", controllerMgr.ClientContractController.DeleteClientContract)
	}

	employeesRoute := routers.Group("/api/hr")
	{
		// routers endpoint/url http category
		employeesRoute.GET("/employees", controllerMgr.EmployeeMockupController.ListEmployeeMockup)
		employeesRoute.POST("/employees/create", controllerMgr.EmployeeMockupController.CreateEmployeeMockup)
		employeesRoute.PUT("/employees/update/:id", controllerMgr.EmployeeMockupController.UpdateEmployeeMockup)
		employeesRoute.GET("/employees/search", controllerMgr.EmployeeMockupController.SearchEmployee)

	}

	return routers
}
