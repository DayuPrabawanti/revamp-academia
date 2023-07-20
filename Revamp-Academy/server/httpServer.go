package server

import (
	// "codeid.revampacademy/config"
	"database/sql"
	"log"

	"codeid.revampacademy/controller"
	"codeid.revampacademy/repositories"
	"codeid.revampacademy/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config *viper.Viper
	router *gin.Engine
	// categoryController *controller.CategoryController
	// masterController *controller.MasterController
	jobHireController *controller.JobHire
}

func InitHttpServer(cfg *viper.Viper, dbHandler *sql.DB) HttpServer {

	//set router from gin first
	router := gin.Default()

	categoryRepo := repositories.NewCategoryRepo(dbHandler)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryControl := controller.NewCategoryController(categoryService)

	//buat router Endpoint
	router.GET("/listJobCategory", categoryControl.GetListCategoryControl)

	masterRepo := repositories.NewMasterRepo(dbHandler)
	masterService := service.NewMasterService(masterRepo)
	masterController := controller.NewMasterController(masterService)

	//make endpoint route
	router.GET("/listaddress", masterController.GetListAddressControl)
	router.GET("listcity", masterController.GetListCityControl)

	jobRepo := repositories.NewJobPostRepo(dbHandler)
	jobService := service.NewJobService(jobRepo)
	jobController := controller.NewJobControll(jobService)

	router.GET("/jobs", jobController.GetJobPostMergeControl)
	router.GET("/dumpJobs", jobController.GetJobPostControl)

	return HttpServer{
		config: cfg,
		router: router,
		// router:             JobHireEndpointRoute(cfg, dbHandler),
		jobHireController: jobController,
	}
}

// running for gin server
func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("Error While starting HTTP Server : %v", err)
	}
}
