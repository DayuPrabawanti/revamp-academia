package bootcampController

import (
	"net/http"
	"strconv"

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
	batchId := ctx.Param("batchid")
	// Convert string batchId to int32
	batchIdInt, err := strconv.ParseInt(batchId, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid batch ID"})
		return
	}

	response, responseErr := bootcampBatchEvaluationController.bootcampBatchEvaluationService.GetListBootcampBatchEvaluation(ctx, int32(batchIdInt))

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// func (bootcampBatchEvaluationController BootcampBatchEvaluationController) GetListBootcampBatchEvaluation(ctx *gin.Context) {
// 	batchIdStr := ctx.Param("batchId") // get batchId from request params

// 	// Convert batchId from string to int
// 	batchId, err := strconv.Atoi(batchIdStr)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid batch ID"})
// 		return
// 	}

// 	response, responseErr := bootcampBatchEvaluationController.bootcampBatchEvaluationService.GetListBootcampBatchEvaluation(ctx, batchId)

// 	if responseErr != nil {
// 		ctx.JSON(responseErr.Status, responseErr)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, response)
// }
