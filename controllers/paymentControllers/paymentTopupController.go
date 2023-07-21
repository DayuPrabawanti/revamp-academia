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

// Method
func (paymentTopupController PaymentTopupController) GetListTopupDetail(ctx *gin.Context) {
	response, responseErr := paymentTopupController.paymentTopupService.GetListTopupDetail(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

	// ctx.JSON(http.StatusOK, "Hello Gin Framework!")
}
