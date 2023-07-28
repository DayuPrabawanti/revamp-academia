package paymentControllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"codeid.revampacademy/repositories/paymentRepositories/dbContext"
	services "codeid.revampacademy/services/paymentServices"
	"github.com/gin-gonic/gin"
)

type PaymentFintechController struct {
	paymentFintechService *services.PaymentFintechService
}

// Declare constructor
func NewPaymentFintechController(paymentFintechService *services.PaymentFintechService) *PaymentFintechController {
	return &PaymentFintechController{
		paymentFintechService: paymentFintechService,
	}
}

// Method
func (paymentFintechController PaymentFintechController) GetListPaymentFintech(ctx *gin.Context) {
	response, responseErr := paymentFintechController.paymentFintechService.GetListPaymentFintech(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

	// ctx.JSON(http.StatusOK, "Hello Gin Framework!")
}

func (paymentFintechController PaymentFintechController) GetPaymentFintechByName(ctx *gin.Context) {
	paymentFintechName := ctx.Query("name")

	response, responseErr := paymentFintechController.paymentFintechService.GetPaymentFintechByName(ctx, paymentFintechName)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (paymentFintechController PaymentFintechController) CreateNewPaymentFintech(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var paymentFintech dbContext.CreatePaymentFintechParams
	err = json.Unmarshal(body, &paymentFintech)
	if err != nil {
		log.Println("Error while unmarshaling create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := paymentFintechController.paymentFintechService.CreateNewPaymentFintech(ctx, &paymentFintech)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (paymentFintechController PaymentFintechController) UpdatePaymentFintechById(ctx *gin.Context) {
	fintEntityId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var paymentFintech dbContext.CreatePaymentFintechParams
	err = json.Unmarshal(body, &paymentFintech)
	if err != nil {
		log.Println("Error while unmarshaling update category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := paymentFintechController.paymentFintechService.UpdatePaymentFintechById(ctx, &paymentFintech, int64(fintEntityId))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (paymentFintechController PaymentFintechController) DeletePaymentFintechById(ctx *gin.Context) {

	paymentFintechId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseErr := paymentFintechController.paymentFintechService.DeletePaymentFintechById(ctx, int64(paymentFintechId))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}
