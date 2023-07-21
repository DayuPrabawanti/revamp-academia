package server

import (
	"codeid.revampacademy/controller/jobhireController"
	"github.com/gin-gonic/gin"
)

func InitRouter(controllerManager *jobhireController.ControllerManager) *gin.Engine {
	//set router from gin
	router := gin.Default()

	//Membuat router Endpoint
	jobRoute := router.Group("/jobs")
	{
		jobRoute.GET("/listJobCategory", controllerManager.GetListCategoryControl)
		jobRoute.GET("", controllerManager.GetJobPostMergeControl)
		jobRoute.GET("/dumpJobs", controllerManager.GetJobPostControl)
	}

	masterRoute := router.Group("/masterdata")
	{
		masterRoute.GET("/listaddress", controllerManager.GetListAddressControl)
		masterRoute.GET("/listcity", controllerManager.GetListCityControl)
	}

	return router
}
