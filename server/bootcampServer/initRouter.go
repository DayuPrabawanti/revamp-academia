package bootcampServer

import (
	"codeid.revampacademy/controllers/bootcampController"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine, controllerMgr *bootcampController.ControllerManager) *gin.Engine {

	batchRoute := router.Group("/api/bootcamp")
	{
		//router endpoint
		//BATCH
		batchRoute.GET("/batch/", controllerMgr.BatchController.GetListBatch)
		batchRoute.GET("/batch/view", controllerMgr.BatchController.GetBatch) // Menggunakan query parameter id
		batchRoute.POST("/batch/create", controllerMgr.BatchController.CreateBatchInstructorTrainee)

		batchRoute.PUT("/batch/update/batchid", controllerMgr.BatchController.UpdateBatch)
		batchRoute.PUT("/batch/updateInstructor/batchid", controllerMgr.BatchController.UpdateInstructorPrograms)

		batchRoute.DELETE("/batch/delete/:id", controllerMgr.BatchController.DeleteBatch)
		batchRoute.DELETE("/batch/delete", controllerMgr.BatchController.DeleteBatchTransaction)
		batchRoute.DELETE("/batch/deleteTrainee", controllerMgr.BatchController.DeleteBatchTrainee2)

		batchRoute.GET("/batch/search", controllerMgr.BatchController.SearchBatch)

		// evaluation
		batchRoute.GET("/batch/evaluation/view", controllerMgr.BootcampBatchEvaluationController.GetBootcampBatchEvaluation)

		batchRoute.GET("/batch/evaluation/review", controllerMgr.BootcampBatchEvaluationController.GetBatchTraineeReview)
		batchRoute.POST("/batch/evaluation/review/create", controllerMgr.BootcampBatchEvaluationController.CreateBatchTraineeReview)
		batchRoute.PUT("/batch/evaluation/review/update", controllerMgr.BootcampBatchEvaluationController.UpdateBatchTraineeReview)
		batchRoute.DELETE("/batch/evaluation/review/delete/:id", controllerMgr.BootcampBatchEvaluationController.DeleteBatchTraineeReview)

		batchRoute.PUT("/batch/evaluation/review/", controllerMgr.BootcampBatchEvaluationController.BootcampBatchTraineeReview)

		batchRoute.GET("/batch/evaluation/scoring", controllerMgr.EvaluationCandidateController.GetEvaluationCandidate)
		batchRoute.POST("/batch/evaluation/scoring/create", controllerMgr.EvaluationCandidateController.CreateEvaluationCandidate)
		batchRoute.PUT("/batch/evaluation/scoring/", controllerMgr.EvaluationCandidateController.UpdateEvaluationCandidate)

	}
	return router
}
