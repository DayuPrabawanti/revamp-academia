package controllers

import (
	"net/http"
	"strconv"

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

func (PaymentTopupController PaymentTopupController) GetAccountByBankCodeAndAccountNumber(ctx *gin.Context) {
	bankCode := ctx.Query("bankCode")
	usacAccountNumber := ctx.Query("usacAccountNumber")

	response, responseErr := PaymentTopupController.paymentTopupService.GetAccountByBankCodeAndAccountNumber(ctx, bankCode, usacAccountNumber)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (PaymentTopupController PaymentTopupController) GetAccountByFintCodeAndAccountNumber(ctx *gin.Context) {
	fintCode := ctx.Query("fintCode")
	usacAccountNumber := ctx.Query("usacAccountNumber")

	response, responseErr := PaymentTopupController.paymentTopupService.GetAccountByFintCodeAndAccountNumber(ctx, fintCode, usacAccountNumber)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (paymentTopupController *PaymentTopupController) TransferTopup(ctx *gin.Context) {
	bankCode := ctx.PostForm("bankCode")
	accountFrom := ctx.PostForm("accountFrom")
	fintCode := ctx.PostForm("fintCode")
	accountTo := ctx.PostForm("accountTo")
	amountStr := ctx.PostForm("amountStr")

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount"})
		return
	}

	transaction, responseErr := paymentTopupController.paymentTopupService.TransferTopup(ctx, bankCode, accountFrom, fintCode, accountTo, amount)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, transaction)
}
