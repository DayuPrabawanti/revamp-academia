package bootcampController

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"codeid.revampacademy/repositories/bootcampRepository/dbContext"
	"codeid.revampacademy/services/bootcampService"
	"github.com/gin-gonic/gin"
)

type BatchController struct {
	batchService *bootcampService.BatchService
}

// declare constructor
func NewBatchController(batchService *bootcampService.BatchService) *BatchController {
	return &BatchController{
		batchService: batchService,
	}
}

// method
func (batchController BatchController) GetListBatch(ctx *gin.Context) {
	response, responseErr := batchController.batchService.GetListBatch(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
	// ctx.JSON(http.StatusOK, "Hello Gin Framework")
}

func (batchController BatchController) GetBatch(ctx *gin.Context) {

	batchId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := batchController.batchService.GetBatch(ctx, int64(batchId))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (batchController BatchController) CreateBatch(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create batch request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var batch dbContext.CreateBatchParams
	err = json.Unmarshal(body, &batch)
	if err != nil {
		log.Println("Error while unmarshaling create batch request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := batchController.batchService.CreateBatch(ctx, &batch)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (batchController BatchController) UpdateBatch(ctx *gin.Context) {

	batchId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update batch request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var batch dbContext.CreateBatchParams
	err = json.Unmarshal(body, &batch)
	if err != nil {
		log.Println("Error while unmarshaling update batch request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := batchController.batchService.UpdateBatch(ctx, &batch, int64(batchId))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (batchController BatchController) DeleteBatch(ctx *gin.Context) {

	batchId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseErr := batchController.batchService.DeleteBatch(ctx, int64(batchId))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}
