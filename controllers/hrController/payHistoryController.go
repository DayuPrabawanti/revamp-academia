package hrController

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"codeid.revampacademy/repositories/hrRepository/dbContext"
	"codeid.revampacademy/services/hrService"
	"github.com/gin-gonic/gin"
)

type PayHistoryController struct {
	payHistoryService *hrService.PayHistoryService
}

// declare constructor
func NewPayHistoryController(payHistoryService *hrService.PayHistoryService) *PayHistoryController {
	return &PayHistoryController{
		payHistoryService: payHistoryService,
	}
}

// create method
func (payHistoryController PayHistoryController) GetListPayHistory(ctx *gin.Context) {

	response, responseErr := payHistoryController.payHistoryService.GetListPayHistory(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
	// ctx.JSON(http.StatusOK, "Hello gin framework")
}

func (payHistoryController PayHistoryController) GetPayHistory(ctx *gin.Context) {

	ephi_entity_id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := payHistoryController.payHistoryService.GetPayHistory(ctx, int64(ephi_entity_id))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (payHistoryController PayHistoryController) CreatePayHistory(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create pay history request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var payHistory dbContext.CreatePayHistoryParams
	err = json.Unmarshal(body, &payHistory)
	if err != nil {
		log.Println("Error while unmarshaling create pay history request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := payHistoryController.payHistoryService.CreatePayHistory(ctx, &payHistory)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (payHistoryController PayHistoryController) UpdatePayHistory(ctx *gin.Context) {

	ephi_entity_id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update pay history request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var payHistory dbContext.UpdatePayHistoryParams
	err = json.Unmarshal(body, &payHistory)
	if err != nil {
		log.Println("Error while unmarshaling update pay history request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := payHistoryController.payHistoryService.UpdatePayHistory(ctx, &payHistory, int64(ephi_entity_id))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (payHistoryController PayHistoryController) DeletePayHistory(ctx *gin.Context) {

	ephi_entity_id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseErr := payHistoryController.payHistoryService.DeletePayHistory(ctx, int64(ephi_entity_id))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}
