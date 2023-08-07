package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"codeid.revampacademy/models/features"
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
	searchBy := ctx.DefaultQuery("accountid", "")
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("pageno", "0"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pagesize", "10"))

	metadata := features.Metadata{
		SearchBy: searchBy,
		PageNo:   pageNo,
		PageSize: pageSize,
	}

	response, responseErr := PaymentTransactionController.paymentTransactionService.GetPaymentTransactionById(ctx, &metadata)
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (paymentTransactionController PaymentTransactionController) RecordPaymentTransactionUser(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var paymentTransaction dbContext.RecordTransactionUserParams
	err = json.Unmarshal(body, &paymentTransaction)
	if err != nil {
		log.Println("Error while unmarshaling create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := paymentTransactionController.paymentTransactionService.RecordPaymentTransactionUser(ctx, &paymentTransaction)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

// func (paymentTransactionController PaymentTransactionController) UpdatePaymentTransaction(ctx *gin.Context) {

// 	trpaID, err := strconv.Atoi(ctx.Param("id"))

// 	if err != nil {
// 		log.Println("Error while reading paramater id", err)
// 		ctx.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}

// 	body, err := io.ReadAll(ctx.Request.Body)
// 	if err != nil {
// 		log.Println("Error while reading update category request body", err)
// 		ctx.AbortWithError(http.StatusInternalServerError, err)
// 		return
// 	}

// 	var paymentTransaction dbContext.CreatePaymentTransaction_paymentParams
// 	err = json.Unmarshal(body, &paymentTransaction)
// 	if err != nil {
// 		log.Println("Error while unmarshaling update category request body", err)
// 		ctx.AbortWithError(http.StatusInternalServerError, err)
// 		return
// 	}

// 	response := paymentTransactionController.paymentTransactionService.UpdatePaymentTransaction(ctx, &paymentTransaction, int64(trpaID))
// 	if response != nil {
// 		ctx.AbortWithStatusJSON(response.Status, response)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, response)

// }

func (paymentTransactionController PaymentTransactionController) DeletePaymentTransaction_paymentById(ctx *gin.Context) {

	trpaCodeNumber := ctx.Query("id")

	responseErr := paymentTransactionController.paymentTransactionService.DeletePaymentTransaction_paymentById(ctx, trpaCodeNumber)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}
