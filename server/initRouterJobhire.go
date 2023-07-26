package server

import (
	"codeid.revampacademy/controllers/jobhireController"
	"github.com/gin-gonic/gin"
)

func InitRouterJobhire(router *gin.Engine, controllerManager *jobhireController.ControllerManager) *gin.Engine {

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
