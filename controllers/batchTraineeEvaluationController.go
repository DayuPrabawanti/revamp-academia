package controllers

import (
	"net/http"

	"codeid.revampacademy/services"
	"github.com/gin-gonic/gin"
)

type BatchTraineeEvaluationController struct {
	batchTraineeEvaluationService *services.BatchTraineeEvaluationService
}

// declare constructor
func NewBatchTraineeEvaluationController(batchTraineeEvaluationService *services.BatchTraineeEvaluationService) *BatchTraineeEvaluationController {
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
