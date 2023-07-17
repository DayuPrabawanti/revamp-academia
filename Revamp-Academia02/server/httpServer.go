package server

import (
	"database/sql"
	"log"

	"codeid.revamptwo/controllers/hrsCr"
	"codeid.revamptwo/repositories/hrs"
	"codeid.revamptwo/services/hrsSc"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config                 *viper.Viper
	router                 *gin.Engine
	departmentController   *hrsCr.DepartmentController
	talentDetailController *hrsCr.TalentsDetailMockupController
	employeeController     *hrsCr.EmployeeController
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	// DEPARTMENT
	departmentRepository := hrs.NewDepartmentRepository(dbHandler)
	departmentServices := hrsSc.NewDepartmentService(departmentRepository)
	departmentController := hrsCr.NewDepartmentController(departmentServices)

	// EMPLOYEE
	employeeRepository := hrs.NewEmployeeRepository(dbHandler)
	employeeService := hrsSc.NewEmployeeService(employeeRepository)
	employeeController := hrsCr.NewEmployeeController(employeeService)

	// TALENT DETAIL
	talentDetailRepository := hrs.NewTalentDetailMockupRepository(dbHandler)
	talentDetailService := hrsSc.NewTalentDetailMockupService(talentDetailRepository)
	talentDetailController := hrsCr.NewTalentDetailMockupController(talentDetailService)

	router := gin.Default()
	// router endpoint/url http Department
	router.GET("/department", departmentController.GetListDepartment)
	router.GET("/department/:id", departmentController.GetDepartment)
	router.POST("/department", departmentController.CreateDepartment)
	router.PUT("/department/:id", departmentController.UpdateDepartment)
	router.DELETE("/department/:id", departmentController.DeleteDepartment)

	// router endpoint/url http Employee
	router.GET("/employee", employeeController.GetListEmployee)
	router.GET("/employee/:id", employeeController.GetEmployee)
	router.POST("/employee", employeeController.CreateEmployee)
	router.PUT("/employee/:id", employeeController.UpdateEmployee)
	router.DELETE("/employee/:id", employeeController.DeleteEmployee)

	// router endpoint/url http TalentDetail
	router.GET("/talentdetail", talentDetailController.GetListTalentDetailMockuptDetail)

	// router

	return HttpServer{
		config:                 config,
		router:                 router,
		departmentController:   departmentController,
		talentDetailController: talentDetailController,
		employeeController:     employeeController,
	}
}

// method running gin httpserver
func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("Error while starting HTTP Server : %v ", err)
	}
}
