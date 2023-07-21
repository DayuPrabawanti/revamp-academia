package bootcampController

import (
	"net/http"

	"codeid.revampacademy/services/bootcampService"
	"github.com/gin-gonic/gin"
)

type BatchTraineeController struct {
	batchTraineeService *bootcampService.BatchTraineeService
}

// declare constructor
func NewBatchTraineeController(batchTraineeService *bootcampService.BatchTraineeService) *BatchTraineeController {
	return &BatchTraineeController{
		batchTraineeService: batchTraineeService,
	}
}

// method
func (batchTraineeController BatchTraineeController) GetListBatchTrainee(ctx *gin.Context) {
	response, responseErr := batchTraineeController.batchTraineeService.GetListBatchTrainee(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}
