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
	config               *viper.Viper
	router               *gin.Engine
	departmentController *controllers.DepartmentController
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	departmentRepository := repositories.NewDepartmentRepository(dbHandler)

	departmentService := services.NewDepartmentService(departmentRepository)

	departmentController := controllers.NewDepartmentController(departmentService)

	router := gin.Default()

	//router endpoint table department
	router.GET("/department", departmentController.GetListDepartment)
	router.GET("/department/:id", departmentController.GetDepartment)
	router.POST("/department", departmentController.CreateDepartment)
	router.PUT("/department/:id", departmentController.UpdateDepartment)
	router.DELETE("/department/:id", departmentController.DeleteDepartment)

	return HttpServer{
		config:               config,
		router:               router,
		departmentController: departmentController,
	}
}

// create method for running gin httpserver
func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http server address"))
	if err != nil {
		log.Fatal("Error while starting HTTP Server : %v", err)
	}
}
