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

// Method
// func (paymentTopupController PaymentTopupController) GetListTopupDetail(ctx *gin.Context) {
// 	response, responseErr := paymentTopupController.paymentTopupService.GetListTopupDetail(ctx)

// 	if responseErr != nil {
// 		ctx.JSON(responseErr.Status, responseErr)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, response)

// 	// ctx.JSON(http.StatusOK, "Hello Gin Framework!")
// }

// func (paymentTopupController PaymentTopupController) GetTopupDetailById(ctx *gin.Context) {
// 	userEntityID, err := strconv.Atoi(ctx.Param("id"))

// 	if err != nil {
// 		log.Println("Error while reading paramater id", err)
// 		ctx.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}

// 	response, responseErr := paymentTopupController.paymentTopupService.GetTopupDetailById(ctx, int32(userEntityID))
// 	if responseErr != nil {
// 		ctx.JSON(responseErr.Status, responseErr)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, response)
// }

// func (ptc *PaymentTopupController) PerformTransfer(ctx *gin.Context) {
// 	// Use ShouldBindJSON to automatically bind the incoming JSON to struct.
// 	var transferReq struct {
// 		FromAccount string  `json:"from_account"`
// 		ToAccount   string  `json:"to_account"`
// 		FromUserID  int     `json:"from_user_id"`
// 		ToUserID    int     `json:"to_user_id"`
// 		Amount      float64 `json:"amount"`
// 	}
// 	if err := ctx.ShouldBindJSON(&transferReq); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if responseErr := ptc.paymentTopupService.PerformTransfer(
// 		ctx,
// 		transferReq.FromAccount,
// 		transferReq.Amount,
// 		transferReq.ToAccount,
// 		transferReq.FromUserID,
// 		transferReq.ToUserID,
// 	); responseErr != nil {
// 		ctx.JSON(responseErr.Status, responseErr.Message)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{"status": "transfer successful"})
// }
