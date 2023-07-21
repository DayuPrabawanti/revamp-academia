package bootcampController

import (
	"net/http"

	"codeid.revampacademy/services/bootcampService"
	"github.com/gin-gonic/gin"
)

type BootcampBatchEvaluationController struct {
	bootcampBatchEvaluationService *bootcampService.BootcampBatchEvaluationService
}

// declare constructor
func NewBootcampBatchEvaluationController(bootcampBatchEvaluationService *bootcampService.BootcampBatchEvaluationService) *BootcampBatchEvaluationController {
	return &BootcampBatchEvaluationController{
		bootcampBatchEvaluationService: bootcampBatchEvaluationService,
	}
}

func (bootcampBatchEvaluationController BootcampBatchEvaluationController) GetListBootcampBatchEvaluation(ctx *gin.Context) {
	response, responseErr := bootcampBatchEvaluationController.bootcampBatchEvaluationService.GetListBootcampBatchEvaluation(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
