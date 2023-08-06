package controllers

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
func (paymentAccountController PaymentAccountController) GetListPaymentUsers_accountByUserName(ctx *gin.Context) {
	userName := ctx.Query("userName")

	response, responseErr := paymentAccountController.paymentAccountService.GetListPaymentUsers_accountByUserName(ctx, userName)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (paymentAccountController PaymentAccountController) GetPaymentAccountByAccountNumber(ctx *gin.Context) {
	usacAccountNumber := ctx.Query("usacAccountNumber")

	response, responseErr := paymentAccountController.paymentAccountService.GetPaymentAccountByAccountNumber(ctx, usacAccountNumber)
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

func (paymentAccountController PaymentAccountController) UpdatePaymentUsers_accountPlus(ctx *gin.Context) {
	usacAccountNumber := ctx.Query("usacAccountNumber")

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var paymentAccount dbContext.UpdatePaymentUsers_accountParams
	err = json.Unmarshal(body, &paymentAccount)
	if err != nil {
		log.Println("Error while unmarshaling create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := paymentAccountController.paymentAccountService.UpdatePaymentUsers_accountPlus(ctx, &paymentAccount, usacAccountNumber)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (paymentAccountController PaymentAccountController) UpdatePaymentUsers_accountMinus(ctx *gin.Context) {
	usacAccountNumber := ctx.Query("usacAccountNumber")

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var paymentAccount dbContext.UpdatePaymentUsers_accountParams
	err = json.Unmarshal(body, &paymentAccount)
	if err != nil {
		log.Println("Error while unmarshaling create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := paymentAccountController.paymentAccountService.UpdatePaymentUsers_accountMinus(ctx, &paymentAccount, usacAccountNumber)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (paymentAccountController PaymentAccountController) DeletePaymentAccountByAccNum(ctx *gin.Context) {
	usacAccountNumber := ctx.Query("usacAccountNumber")

	responseErr := paymentAccountController.paymentAccountService.DeletePaymentAccountByAccNum(ctx, usacAccountNumber)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (paymentAccountController *PaymentAccountController) DebitSaldo(ctx *gin.Context) {
	usacAccountNumber := ctx.PostForm("usacAccountNumber")
	amountStr := ctx.PostForm("amountStr")

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount"})
		return
	}

	transaction, responseErr := paymentAccountController.paymentAccountService.DebitSaldo(ctx, usacAccountNumber, amount)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, transaction)
}
