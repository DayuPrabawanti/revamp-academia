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
		batchRoute.POST("/batch/create", controllerMgr.BatchController.CreateBatchInstructorTrainee)

		batchRoute.PUT("/batch/update/batchid", controllerMgr.BatchController.UpdateBatch)
		batchRoute.PUT("/batch/updateInstructor/batchid", controllerMgr.BatchController.UpdateInstructorPrograms)

		batchRoute.DELETE("/batch/delete/:id", controllerMgr.BatchController.DeleteBatch)
		batchRoute.DELETE("/batch/delete", controllerMgr.BatchController.DeleteBatchTransaction)
		batchRoute.DELETE("/batch/deleteTrainee", controllerMgr.BatchController.DeleteBatchTrainee2)

		batchRoute.GET("/batch/search", controllerMgr.BatchController.SearchBatch)
	}
	return router
}
