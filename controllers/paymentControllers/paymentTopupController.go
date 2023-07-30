package controllers

import (
	"log"
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

func (paymentTopupController PaymentTopupController) GetTopupDetailById(ctx *gin.Context) {
	userEntityID, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := paymentTopupController.paymentTopupService.GetTopupDetailById(ctx, int32(userEntityID))
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
