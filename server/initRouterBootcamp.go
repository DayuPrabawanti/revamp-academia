package server

import (
	"codeid.revampacademy/controllers/bootcampController"
	"github.com/gin-gonic/gin"
)

func InitRouterBootcamp(router *gin.Engine, controllerMgr *bootcampController.ControllerManager) *gin.Engine {

	batchRoute := router.Group("/api/bootcamp")
	{
		//router endpoint
		batchRoute.GET("/batch/", controllerMgr.BatchController.GetListBatch)
		batchRoute.GET("/batch/view", controllerMgr.BatchController.GetBatch) // Menggunakan query parameter id
		batchRoute.POST("/batch/create", controllerMgr.BatchController.CreateBatch)

		batchRoute.PUT("/batch/update/batchid", controllerMgr.BatchController.UpdateBatch)
		batchRoute.DELETE("/batch/delete/:id", controllerMgr.BatchController.DeleteBatch)

		batchRoute.GET("/batch/search", controllerMgr.BatchController.SearchBatch)
		batchRoute.GET("/batch/paging", controllerMgr.BatchController.PagingBatch)
	}
	return router
}
