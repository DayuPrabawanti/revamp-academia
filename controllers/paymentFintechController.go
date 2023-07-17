package controllers

import (
	"net/http"

	"codeid.revampacademy/services"
	"github.com/gin-gonic/gin"
)

type PaymentFintechController struct {
	paymentFintechService *services.PaymentFintechService
}

// declare cunstructor
func NewPaymentFintechController(paymentFintechService *services.PaymentFintechService) *PaymentFintechController {
	return &PaymentFintechController{
		paymentFintechService: paymentFintechService,
	}
}

// method
func (paymentFintechController PaymentFintechController) GetListPaymentFintech(ctx *gin.Context) {
	response, responseErr := paymentFintechController.paymentFintechService.GetListPaymentFintech(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
