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

type BootcampBatchEvaluationController struct {
	bootcampBatchEvaluationService *bootcampService.BootcampBatchEvaluationService
}

// declare constructor
func NewBootcampBatchEvaluationController(bootcampBatchEvaluationService *bootcampService.BootcampBatchEvaluationService) *BootcampBatchEvaluationController {
	return &BootcampBatchEvaluationController{
		bootcampBatchEvaluationService: bootcampBatchEvaluationService,
	}
}

func (bootcampBatchEvaluationController BootcampBatchEvaluationController) GetBootcampBatchEvaluation(ctx *gin.Context) {
	id := ctx.Query("batchid") // Mengambil nilai query parameter id dari URL

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "batchid is required"})
		return
	}

	batchID, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Error while parsing id", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid batchid"})
		return
	}

	response, responseErr := bootcampBatchEvaluationController.bootcampBatchEvaluationService.GetBootcampBatchEvaluation(ctx, int32(batchID))
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (bootcampBatchEvaluationController BootcampBatchEvaluationController) GetBatchTraineeReview(ctx *gin.Context) {
	id := ctx.Query("userentityid") // Mengambil nilai query parameter id dari URL

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user entity id is required"})
		return
	}

	userentityid, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Error while parsing id", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user entity id"})
		return
	}

	response, responseErr := bootcampBatchEvaluationController.bootcampBatchEvaluationService.GetBatchTraineeReview(ctx, int32(userentityid))
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (bootcampBatchEvaluationController BootcampBatchEvaluationController) CreateBatchTraineeReview(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create batch trainee request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var batchTraineeReview dbContext.CreateBatchTraineeReviewParams
	err = json.Unmarshal(body, &batchTraineeReview)
	if err != nil {
		log.Println("Error while unmarshaling create batch trainee request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := bootcampBatchEvaluationController.bootcampBatchEvaluationService.CreateBatchTraineeReview(ctx, &batchTraineeReview)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (bootcampBatchEvaluationController BootcampBatchEvaluationController) UpdateBatchTraineeReview(ctx *gin.Context) {
	userentityid, err := strconv.Atoi(ctx.Query("userentityid")) // atoi mengubah string ke integer

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

	var batchTraineeReview dbContext.UpdateBatchTraineeReviewParams
	err = json.Unmarshal(body, &batchTraineeReview)
	if err != nil {
		log.Println("Error while unmarshaling update batch request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := bootcampBatchEvaluationController.bootcampBatchEvaluationService.UpdateBatchTraineeReview(ctx, &batchTraineeReview, int64(userentityid))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (bootcampBatchEvaluationController BootcampBatchEvaluationController) DeleteBatchTraineeReview(ctx *gin.Context) {

	batrId, err := strconv.Atoi(ctx.Param("id")) // atoi mengubah string ke integer

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseErr := bootcampBatchEvaluationController.bootcampBatchEvaluationService.DeleteBatchTraineeReview(ctx, int64(batrId))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (bootcampBatchEvaluationController BootcampBatchEvaluationController) BootcampBatchTraineeReview(ctx *gin.Context) {
	userEntityIDStr := ctx.DefaultQuery("userentityid", "0")
	userEntityID, err := strconv.Atoi(userEntityIDStr)
	if err != nil || userEntityID <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing userentityid"})
		return
	}

	BatrStatus := ctx.PostForm("BatrStatus")
	BatrReview := ctx.PostForm("BatrReview")

	updatedReview, responseErr := bootcampBatchEvaluationController.bootcampBatchEvaluationService.BootcampBatchTraineeReview(ctx, int32(userEntityID), BatrStatus, BatrReview)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, updatedReview)
}

// 	batchId, err := strconv.Atoi(ctx.Param("id")) // atoi mengubah string ke integer

// 	if err != nil {
// 		log.Println("Error while reading paramater id", err)
// 		ctx.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}

// 	response, responseErr := bootcampBatchEvaluationController.bootcampBatchEvaluationService.GetBootcampBatchEvaluation(ctx, int32(batchId))
// 	if responseErr != nil {

// 		ctx.JSON(responseErr.Status, responseErr)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, response)
// }

// func (bootcampBatchEvaluationController BootcampBatchEvaluationController) GetListBootcampBatchEvaluation(ctx *gin.Context) {
// 	batchId := ctx.Param("batchid")
// 	// Convert string batchId to int32
// 	batchIdInt, err := strconv.ParseInt(batchId, 10, 32)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid batch ID"})
// 		return
// 	}

// 	response, responseErr := bootcampBatchEvaluationController.bootcampBatchEvaluationService.GetListBootcampBatchEvaluation(ctx, int32(batchIdInt))

// 	if responseErr != nil {
// 		ctx.JSON(responseErr.Status, responseErr)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, response)
// }

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
