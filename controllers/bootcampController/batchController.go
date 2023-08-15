package bootcampController

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"codeid.revampacademy/models"
	"codeid.revampacademy/models/features"
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

// create method
func (batchController BatchController) GetListBatch(ctx *gin.Context) {
	//add metadata to hold data from query parameter, use defaultquery
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "0"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "3"))
	batch := ctx.DefaultQuery("batch", "")
	status := ctx.DefaultQuery("status", "")

	metadata := features.Metadata{
		Page:     page,
		PageSize: pageSize,
		Batch:    batch,
		Status:   status,
	}

	response, responseErr := batchController.batchService.GetListBatch(ctx, &metadata)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (batchController *BatchController) GetBatch(ctx *gin.Context) {
	id := ctx.Query("batchid")

	batchID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid batch id",
		})
		return
	}

	batch, responseErr := batchController.batchService.GetBatchWithMembers(ctx, int64(batchID))
	if responseErr != nil {
		ctx.JSON(responseErr.Status, gin.H{
			"message": responseErr.Message,
		})
		return
	}

	ctx.JSON(http.StatusOK, batch)
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

func (batchController BatchController) CreateInstructorPrograms(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create batch request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var batch dbContext.CreateInstructorProgramsParams
	err = json.Unmarshal(body, &batch)
	if err != nil {
		log.Println("Error while unmarshaling create batch request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := batchController.batchService.CreateInstructorPrograms(ctx, &batch)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (batchController BatchController) CreateBatchTrainee(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create batch request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var batch dbContext.CreateBatchTraineeParams
	err = json.Unmarshal(body, &batch)
	if err != nil {
		log.Println("Error while unmarshaling create batch request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := batchController.batchService.CreateBatchTrainee(ctx, &batch)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (batchController BatchController) UpdateBatch(ctx *gin.Context) {
	id := ctx.Query("id") // Mengambil nilai query parameter id dari URL

	batchId, err := strconv.Atoi(id)

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

	response := batchController.batchService.UpdateBatch(ctx, &batch, int32(batchId))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (batchController BatchController) UpdateInstructorPrograms(ctx *gin.Context) {
	id := ctx.Query("id") // Mengambil nilai query parameter id dari URL

	batchId, err := strconv.Atoi(id)

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update instructor programs request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var instructorPrograms dbContext.CreateInstructorProgramsParams
	err = json.Unmarshal(body, &instructorPrograms)
	if err != nil {
		log.Println("Error while unmarshaling update instructor programs request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := batchController.batchService.UpdateInstructorPrograms(ctx, &instructorPrograms, int32(batchId))
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

func (batchController BatchController) DeleteInstructorPrograms(ctx *gin.Context) {

	batchId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseErr := batchController.batchService.DeleteInstructorPrograms(ctx, int64(batchId))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (batchController BatchController) DeleteBatchTrainee(ctx *gin.Context) {

	batchId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseErr := batchController.batchService.DeleteBatchTrainee(ctx, int64(batchId))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (batchController BatchController) DeleteBatchTrainee2(ctx *gin.Context) {
	batrTraineeEntityID, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid trainee entity ID",
		})
		return
	}

	batrBatchID, err := strconv.Atoi(ctx.Query("batch"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid batch ID",
		})
		return
	}

	responseErr := batchController.batchService.DeleteBatchTrainee2(ctx, int64(batrTraineeEntityID), int64(batrBatchID))
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Batch trainee data has been deleted",
	})
}

func (batchController BatchController) SearchBatch(ctx *gin.Context) {
	batchName := ctx.DefaultQuery("batch", "")
	status := ctx.DefaultQuery("status", "")

	batches, responseErr := batchController.batchService.SearchBatch(ctx, batchName, status)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, batches)
}

func (batchController BatchController) CreateBatchInstructorTrainee(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create batch request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var batch models.CreateBatchInstructorTraineeDto
	err = json.Unmarshal(body, &batch)
	if err != nil {
		log.Println("Error while unmarshaling create batch request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := batchController.batchService.CreateBatchInstructorTraineeDto(ctx, &batch)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (batchController BatchController) DeleteBatchTransaction(ctx *gin.Context) {
	batchIdStr := ctx.Query("batchid")
	batchId, err := strconv.Atoi(batchIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid batchID",
		})
		return
	}

	responseErr := batchController.batchService.DeleteBatchTransaction(ctx, int64(batchId))
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Batch data, instructor programs and batch trainee have been deleted",
	})
}
