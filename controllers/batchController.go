package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"codeid.revampacademy/repositories/dbContext"
	"codeid.revampacademy/services"
	"github.com/gin-gonic/gin"
)

type BatchController struct {
	batchService *services.BatchService
}

// declare constructor
func NewBatchController(batchService *services.BatchService) *BatchController {
	return &BatchController{
		batchService: batchService,
	}
}

// create method
func (batchController BatchController) GetListBatch(ctx *gin.Context) {
	response, responseErr := batchController.batchService.GetListBatch(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
	// ctx.JSON(http.StatusOK, "Hello gin Framework")
}

// func (batchController BatchController) GetBatch(ctx *gin.Context) {

// 	batchId, err := strconv.Atoi(ctx.Param("id"))

// 	if err != nil {
// 		log.Println("Error while reading paramater id", err)
// 		ctx.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}

// 	response, responseErr := batchController.batchService.GetBatch(ctx, int64(batchId))
// 	if responseErr != nil {

// 		ctx.JSON(responseErr.Status, responseErr)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, response)
// }

func (batchController BatchController) GetBatch(ctx *gin.Context) {
	id := ctx.Query("id") // Mengambil nilai query parameter id dari URL

	batchId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Error while parsing id", err)
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

func (batchController BatchController) SearchBatch(ctx *gin.Context) {
	batchIDParam := ctx.DefaultQuery("batch", "")
	statusParam := ctx.DefaultQuery("status", "")

	batchID, _ := strconv.Atoi(batchIDParam)

	// Call the batchService.SearchBatch function to perform the search operation.
	batches, err := batchController.batchService.SearchBatch(ctx, int32(batchID), statusParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to search batches",
		})
		return
	}

	ctx.JSON(http.StatusOK, batches)
}
