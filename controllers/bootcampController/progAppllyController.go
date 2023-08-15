package bootcampController

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"codeid.revampacademy/models/features"
	"codeid.revampacademy/repositories/bootcampRepository/dbContext"
	"codeid.revampacademy/services/bootcampService"
	"github.com/gin-gonic/gin"
)

type ProgAppllyController struct {
	ProgApplyService *bootcampService.ProgAppllyService
}

func NewProgApplyController(ProgApplyService *bootcampService.ProgAppllyService) *ProgAppllyController {
	return &ProgAppllyController{
		ProgApplyService: ProgApplyService,
	}
}

func (ProgApplyController ProgAppllyController) GetListProgApply(ctx *gin.Context) {
	response, responseErr := ProgApplyController.ProgApplyService.GetListProgApply(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (ProgApplyController ProgAppllyController) UpdateProgApply(ctx *gin.Context) {

	switchUp, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update Program Apply request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var switchPassed dbContext.ProgApply
	err = json.Unmarshal(body, &switchPassed)
	if err != nil {
		log.Println("Error while unmarshaling update Program Apply request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := ProgApplyController.ProgApplyService.UpdateProgApply(ctx, &switchPassed, int64(switchUp))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (ProgAppllyController ProgAppllyController) GetProgApply(ctx *gin.Context) {

	update, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := ProgAppllyController.ProgApplyService.GetProgApply(ctx, int64(update))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (ProgAppllyController ProgAppllyController) GetTestScore(ctx *gin.Context) {

	prapStatus, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := ProgAppllyController.ProgApplyService.GetTestScore(ctx, int64(prapStatus))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (ProgAppllyController ProgAppllyController) GetReview(ctx *gin.Context) {

	prapStatus, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := ProgAppllyController.ProgApplyService.GetReview(ctx, int64(prapStatus))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (ProgAppllyController ProgAppllyController) GetStatus(ctx *gin.Context) {

	prapStatus, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := ProgAppllyController.ProgApplyService.GetStatus(ctx, int64(prapStatus))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (ProgApplyController ProgAppllyController) GetlistProgApplyStatus(ctx *gin.Context) {
	status := ctx.Query("Apply")
	response, responseErr := ProgApplyController.ProgApplyService.GetlistProgApplyStatus(ctx, status)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (ProgApplyController ProgAppllyController) GetlistProgApplyfiltering(ctx *gin.Context) {
	status := ctx.Query("Filtering")
	response, responseErr := ProgApplyController.ProgApplyService.GetlistProgApplyfiltering(ctx, status)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (ProgApplyController ProgAppllyController) GetlistProgApplycontract(ctx *gin.Context) {
	status := ctx.Query("Contract")
	response, responseErr := ProgApplyController.ProgApplyService.GetlistProgApplycontract(ctx, status)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (ProgApplyController ProgAppllyController) GetlistProgApplyfailed(ctx *gin.Context) {
	status := ctx.Query("Failed")
	response, responseErr := ProgApplyController.ProgApplyService.GetlistProgApplyfailed(ctx, status)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (ProgApplyController ProgAppllyController) GetlistProgApplyidle(ctx *gin.Context) {
	status := ctx.Query("Idle")
	response, responseErr := ProgApplyController.ProgApplyService.GetlistProgApplyidle(ctx, status)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (ProgApplyController ProgAppllyController) UpdateTestScore(ctx *gin.Context) {

	prapStatus, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update Program Apply request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var ProgramApply dbContext.UpdateStatus
	err = json.Unmarshal(body, &ProgramApply)
	if err != nil {
		log.Println("Error while unmarshaling update Program Apply request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := ProgApplyController.ProgApplyService.UpdateTestScore(ctx, &ProgramApply, int32(prapStatus))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (ProgApplyController ProgAppllyController) UpdatePrapStatus(ctx *gin.Context) {

	prapStatus, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update Program Apply request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var ProgramApply dbContext.UpdateStatus
	err = json.Unmarshal(body, &ProgramApply)
	if err != nil {
		log.Println("Error while unmarshaling update Program Apply request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := ProgApplyController.ProgApplyService.UpdatePrapStatus(ctx, &ProgramApply, int32(prapStatus))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (ProgApplyController ProgAppllyController) UpdatePrapReview(ctx *gin.Context) {

	prapStatus, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update Program Apply request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var ProgramApply dbContext.UpdateStatus
	err = json.Unmarshal(body, &ProgramApply)
	if err != nil {
		log.Println("Error while unmarshaling update Program Apply request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := ProgApplyController.ProgApplyService.UpdatePrapReview(ctx, &ProgramApply, int32(prapStatus))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (ProgApplyController ProgAppllyController) GetlistModifiedDate(ctx *gin.Context) {
	month, _ := strconv.Atoi(ctx.DefaultQuery("month", "7"))
	year, _ := strconv.Atoi(ctx.DefaultQuery("year", "2023"))

	metadata := features.Metadata{
		Month: month,
		Year:  year,
	}

	response, responseErr := ProgApplyController.ProgApplyService.GetlistModifiedDate(ctx, &metadata)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
