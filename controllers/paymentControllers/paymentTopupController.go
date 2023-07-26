package paymentControllers

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

func (paymentTopupController PaymentTopupController) GetTopupDetail(ctx *gin.Context) {
	sourceBankEntityID, err := strconv.Atoi(ctx.Param("bankId"))
	targetFintechEntityID, err := strconv.Atoi(ctx.Param("fintId"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := paymentTopupController.paymentTopupService.GetTopupDetail(ctx, int32(sourceBankEntityID), int32(targetFintechEntityID))
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
