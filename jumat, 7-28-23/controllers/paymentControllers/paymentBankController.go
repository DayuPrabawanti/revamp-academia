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

type PaymentBankController struct {
	paymentBankService *services.PaymentBankService
}

// fungsi deklarasi variabel controller
func NewPaymentBankController(paymentBankService *services.PaymentBankService) *PaymentBankController {
	return &PaymentBankController{
		paymentBankService: paymentBankService,
	}
}

// method utk ambil getlist payment
func (PaymentBankController PaymentBankController) GetListPaymentBank(ctx *gin.Context) {
	response, responseErr := PaymentBankController.paymentBankService.GetListPaymentBank(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// method utk ambil getbyname payment
func (paymentBankController PaymentBankController) GetPaymentBankByName(ctx *gin.Context) {
	bankName := ctx.Query("name")

	response, responseErr := paymentBankController.paymentBankService.GetPaymentBankByName(ctx, string(bankName))
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// method utk create paymentbank
func (paymentBankController PaymentBankController) CreateNewPaymentBank(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var paymentBank dbContext.CreatePaymentBankParams
	err = json.Unmarshal(body, &paymentBank)
	if err != nil {
		log.Println("Error while unmarshaling create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := paymentBankController.paymentBankService.CreateNewPaymentBank(ctx, &paymentBank)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (paymentBankController PaymentBankController) UpdatePaymentBank(ctx *gin.Context) {

	bankEntityID, err := strconv.Atoi(ctx.Param("id"))

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

	var paymentBank dbContext.CreatePaymentBankParams
	err = json.Unmarshal(body, &paymentBank)
	if err != nil {
		log.Println("Error while unmarshaling update category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := paymentBankController.paymentBankService.UpdatePaymentBank(ctx, &paymentBank, int64(bankEntityID))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

// method utk delet payment bank
func (paymentBankController PaymentBankController) DeletePaymentBank(ctx *gin.Context) {

	bankEntityID, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseErr := paymentBankController.paymentBankService.DeletePaymentBank(ctx, int64(bankEntityID))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}
