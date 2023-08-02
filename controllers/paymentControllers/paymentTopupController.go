package controllers

import (
	"net/http"

	services "codeid.revampacademy/services/paymentServices"
	"github.com/gin-gonic/gin"
)

type PaymentTopupController struct {
	paymentTopupService *services.PaymentTopupService
}

// Declare constructor
func NewPaymentTopupController(paymentTopupService *services.PaymentTopupService) *PaymentTopupController {
	return &PaymentTopupController{
		paymentTopupService: paymentTopupService,
	}
}

func (paymentTopupController PaymentTopupController) GetAccountByBankCodeAndAccountNumber(ctx *gin.Context) {
	bankCode := ctx.Query("bankCode")
	usacAccountNumber := ctx.Query("usacAccountNumber")

	response, responseErr := paymentTopupController.paymentTopupService.GetAccountByBankCodeAndAccountNumber(ctx, bankCode, usacAccountNumber)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (paymentTopupController PaymentTopupController) GetAccountByFintCodeAndAccountNumber(ctx *gin.Context) {
	fintCode := ctx.Query("fintCode")
	usacAccountNumber := ctx.Query("usacAccountNumber")

	response, responseErr := paymentTopupController.paymentTopupService.GetAccountByFintCodeAndAccountNumber(ctx, fintCode, usacAccountNumber)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
