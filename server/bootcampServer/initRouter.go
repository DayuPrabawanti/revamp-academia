package bootcampServer

import (
	"codeid.revampacademy/controllers/bootcampController"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine, controllerMgr *bootcampController.ControllerManager) *gin.Engine {

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

		// evaluation
		batchRoute.GET("/batch/evaluation/view", controllerMgr.BootcampBatchEvaluationController.GetListBootcampBatchEvaluation)

		//candidate
		batchRoute.GET("/candidate/gabung", controllerMgr.GabungController.GetListGabung)
		batchRoute.PUT("/candidate/switchstatus/userentityid/:id", controllerMgr.ProgAppllyController.UpdateProgApply)
		batchRoute.GET("/candidate/filterby/status/apply", controllerMgr.ProgAppllyController.GetlistProgApplyStatus)
		batchRoute.GET("/candidate/filterby/status/filtering", controllerMgr.ProgAppllyController.GetlistProgApplyfiltering)
		batchRoute.GET("/candidate/filterby/status/contract", controllerMgr.ProgAppllyController.GetlistProgApplycontract)
		batchRoute.GET("/candidate/filterby/status/failed", controllerMgr.ProgAppllyController.GetlistProgApplyfailed)
		batchRoute.GET("/candidate/filterby/status/idle", controllerMgr.ProgAppllyController.GetlistProgApplyidle)
		batchRoute.PUT("/candidate/updatereview/:id", controllerMgr.ProgAppllyController.UpdatePrapReview)
		batchRoute.GET("/candidate/getstatus/:id", controllerMgr.ProgAppllyController.GetStatus)
		batchRoute.GET("/candidate/status/:id", controllerMgr.ProgAppllyController.GetProgApply)
		batchRoute.GET("/candidate/filterby", controllerMgr.ProgAppllyController.GetlistModifiedDate)

	}
	return router
}
