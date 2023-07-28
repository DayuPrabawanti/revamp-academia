package paymentServices

import (
	"net/http"

	"codeid.revampacademy/models"
	repositories "codeid.revampacademy/repositories/paymentRepositories"
	"codeid.revampacademy/repositories/paymentRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type PaymentTransactionService struct {
	paymentTransactionRepository *repositories.PaymentTransactionRepository
}

func NewPaymentTransactionService(paymentTransactionRepository *repositories.PaymentTransactionRepository) *PaymentTransactionService {
	return &PaymentTransactionService{
		paymentTransactionRepository: paymentTransactionRepository,
	}
}

func (ptr PaymentTransactionService) GetListPaymentTransaction(ctx *gin.Context) ([]*models.PaymentTransactionPayment, *models.ResponseError) {
	return ptr.paymentTransactionRepository.GetListPaymentTransaction(ctx)
}

func (ptr PaymentTransactionService) GetPaymentTransactionById(ctx *gin.Context, id int64) (*models.PaymentTransactionPayment, *models.ResponseError) {
	return ptr.paymentTransactionRepository.GetPaymentTransactionById(ctx, id)
}

func (ptr PaymentTransactionService) CreateNewPaymentTransaction(ctx *gin.Context, paymentTransactionParams *dbContext.CreatePaymentTransaction_paymentParams) (*models.PaymentTransactionPayment, *models.ResponseError) {
	responseErr := validatePaymentTransaction(paymentTransactionParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return ptr.paymentTransactionRepository.CreatePaymentTransaction(ctx, paymentTransactionParams)
}

func (ptr PaymentTransactionService) UpdatePaymentTransaction(ctx *gin.Context, paymentTransactionParams *dbContext.CreatePaymentTransaction_paymentParams, id int64) *models.ResponseError {
	responseErr := validatePaymentTransaction(paymentTransactionParams)
	if responseErr != nil {
		return responseErr
	}
	return ptr.paymentTransactionRepository.UpdatePaymentTransaction(ctx, paymentTransactionParams)

}

func (ptr PaymentTransactionService) DeletePaymentTransaction(ctx *gin.Context, id int64) *models.ResponseError {
	return ptr.paymentTransactionRepository.DeletePaymentTransaction(ctx, id)
}

func validatePaymentTransaction(paymentTransactionParams *dbContext.CreatePaymentTransaction_paymentParams) *models.ResponseError {
	if paymentTransactionParams.TrpaID == 0 {
		return &models.ResponseError{
			Message: "Invalid category id",
			Status:  http.StatusBadRequest,
		}
	}
	return nil

}
