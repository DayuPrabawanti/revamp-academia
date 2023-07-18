package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"codeid.revampacademy/repositories/dbContext"
	"codeid.revampacademy/services"
	"github.com/gin-gonic/gin"
)

type PaymentAccountController struct {
	paymentAccountService *services.PaymentAccountService
}

// Declare constructor
func NewPaymentAccountController(paymentAccountService *services.PaymentAccountService) *PaymentAccountController {
	return &PaymentAccountController{
		paymentAccountService: paymentAccountService,
	}
}

// Method
func (paymentAccountController PaymentAccountController) GetListPaymentAccount(ctx *gin.Context) {
	response, responseErr := paymentAccountController.paymentAccountService.GetListPaymentAccount(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

	// ctx.JSON(http.StatusOK, "Hello Gin Framework!")
}

func (paymentAccountController PaymentAccountController) GetPaymentAccountByName(ctx *gin.Context) {
	paymentAccountName := ctx.Query("name")

	response, responseErr := paymentAccountController.paymentAccountService.GetPaymentAccountByName(ctx, paymentAccountName)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (paymentAccountController PaymentAccountController) CreateNewPaymentAccount(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var paymentAccount dbContext.CreatePaymentUsers_accountParams
	err = json.Unmarshal(body, &paymentAccount)
	if err != nil {
		log.Println("Error while unmarshaling create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := paymentAccountController.paymentAccountService.CreateNewPaymentAccount(ctx, &paymentAccount)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (paymentAccountController PaymentAccountController) UpdatePaymentAccountById(ctx *gin.Context) {
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

	var paymentAccount dbContext.CreatePaymentUsers_accountParams
	err = json.Unmarshal(body, &paymentAccount)
	if err != nil {
		log.Println("Error while unmarshaling update category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := paymentAccountController.paymentAccountService.UpdatePaymentAccountById(ctx, &paymentAccount, int64(fintEntityId))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (paymentAccountController PaymentAccountController) DeletePaymentAccountById(ctx *gin.Context) {

	paymwntAccountId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseErr := paymentAccountController.paymentAccountService.DeletePaymentAccountById(ctx, int64(paymwntAccountId))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}
