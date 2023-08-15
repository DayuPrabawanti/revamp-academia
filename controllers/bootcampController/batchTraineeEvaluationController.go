package bootcampController

import (
	"net/http"

	"codeid.revampacademy/services/bootcampService"
	"github.com/gin-gonic/gin"
)

type BatchTraineeEvaluationController struct {
	batchTraineeEvaluationService *bootcampService.BatchTraineeEvaluationService
}

// declare constructor
func NewBatchTraineeEvaluationController(batchTraineeEvaluationService *bootcampService.BatchTraineeEvaluationService) *BatchTraineeEvaluationController {
	return &BatchTraineeEvaluationController{
		batchTraineeEvaluationService: batchTraineeEvaluationService,
	}
}

// method
func (batchTraineeEvaluationController BatchTraineeEvaluationController) GetListBatchTraineeEvaluation(ctx *gin.Context) {
	response, responseErr := batchTraineeEvaluationController.batchTraineeEvaluationService.GetListBatchTraineeEvaluation(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}
