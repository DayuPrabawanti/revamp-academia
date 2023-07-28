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

type PaymentTransactionController struct {
	paymentTransactionService *services.PaymentTransactionService
}

func NewPaymentTransactionController(paymentTransactionService *services.PaymentTransactionService) *PaymentTransactionController {
	return &PaymentTransactionController{
		paymentTransactionService: paymentTransactionService,
	}
}

func (paymentTransactionController PaymentTransactionController) GetListPaymentTransaction(ctx *gin.Context) {
	response, responseErr := paymentTransactionController.paymentTransactionService.GetListPaymentTransaction(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (PaymentTransactionController PaymentTransactionController) GetPaymentTransactionById(ctx *gin.Context) {
	trpaID, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := PaymentTransactionController.paymentTransactionService.GetPaymentTransactionById(ctx, int64(trpaID))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (paymentTransactionController PaymentTransactionController) CreateNewPaymentTransaction(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var paymentTransaction dbContext.CreatePaymentTransaction_paymentParams
	err = json.Unmarshal(body, &paymentTransaction)
	if err != nil {
		log.Println("Error while unmarshaling create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := paymentTransactionController.paymentTransactionService.CreateNewPaymentTransaction(ctx, &paymentTransaction)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (paymentTransactionController PaymentTransactionController) UpdatePaymentTransaction(ctx *gin.Context) {

	trpaID, err := strconv.Atoi(ctx.Param("id"))

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

	var paymentTransaction dbContext.CreatePaymentTransaction_paymentParams
	err = json.Unmarshal(body, &paymentTransaction)
	if err != nil {
		log.Println("Error while unmarshaling update category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := paymentTransactionController.paymentTransactionService.UpdatePaymentTransaction(ctx, &paymentTransaction, int64(trpaID))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (paymentTransactionController PaymentTransactionController) DeletePaymentTransaction(ctx *gin.Context) {

	trpaID, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseErr := paymentTransactionController.paymentTransactionService.DeletePaymentTransaction(ctx, int64(trpaID))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}
