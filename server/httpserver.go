package server

import (
	"database/sql"
	"log"

	"codeid.revampacademy/controllers"
	"codeid.revampacademy/repositories"
	"codeid.revampacademy/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config                   *viper.Viper
	router                   *gin.Engine
	departmentController     *controllers.DepartmentController
	employeeController       *controllers.EmployeeController
	clientContractController *controllers.ClientContractController
	talentDetailController   *controllers.TalentsDetailMockupController
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	departmentRepository := repositories.NewDepartmentRepository(dbHandler)
	departmentService := services.NewDepartmentService(departmentRepository)
	departmentController := controllers.NewDepartmentController(departmentService)

	employeeRepository := repositories.NewEmployeeRepository(dbHandler)
	employeeService := services.NewEmployeeService(employeeRepository)
	employeeController := controllers.NewEmployeeController(employeeService)

	clientContractRepository := repositories.NewClientContractRepository(dbHandler)
	clientContractService := services.NewClientContractService(clientContractRepository)
	clientContractController := controllers.NewClientContractController(clientContractService)

	talentDetailRepository := repositories.NewTalentDetailMockupRepository(dbHandler)
	talentDetailService := services.NewTalentDetailMockupService(talentDetailRepository)
	talentDetailController := controllers.NewTalentDetailMockupController(talentDetailService)

	router := gin.Default()

	//router endpoint table department
	router.GET("/department", departmentController.GetListDepartment)
	router.GET("/department/:id", departmentController.GetDepartment)
	router.POST("/department", departmentController.CreateDepartment)
	router.PUT("/department/:id", departmentController.UpdateDepartment)
	router.DELETE("/department/:id", departmentController.DeleteDepartment)

	//router endpoint table employee
	router.GET("/employee", employeeController.GetListEmployee)
	router.GET("/employee/:id", employeeController.GetEmployee)
	router.POST("/employee", employeeController.CreateEmployee)
	router.PUT("/employee/:id", employeeController.UpdateEmployee)
	router.DELETE("/employee/:id", employeeController.DeleteEmployee)

	//router endpoint table client contract
	router.GET("/clientContract", clientContractController.GetListClientContract)
	router.GET("/clientContract/:id", clientContractController.GetClientContract)
	router.POST("/clientContract", clientContractController.CreateClientContract)
	router.PUT("/clientContract/:id", clientContractController.UpdateClientContract)
	router.DELETE("/clientContract/:id", clientContractController.DeleteClientContract)

	// router endpoint table TalentDetail
	router.GET("/talentdetail", talentDetailController.GetListTalentDetailMockuptDetail)

	return HttpServer{
		config:                   config,
		router:                   router,
		departmentController:     departmentController,
		employeeController:       employeeController,
		clientContractController: clientContractController,
		talentDetailController:   talentDetailController,
	}
}

// create method for running gin httpserver
func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http server address"))
	if err != nil {
		log.Fatal("Error while starting HTTP Server : %v", err)
	}
}
