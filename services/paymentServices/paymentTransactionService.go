package services

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/models/features"
	repositories "codeid.revampacademy/repositories/paymentRepositories"
	"codeid.revampacademy/repositories/paymentRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type PaymentTransactionService struct {
	// paymentTransactionRepository *repositories.PaymentTransactionRepository
	repositoriesManager repositories.RepositoriesManager
}

// func NewPaymentTransactionService(paymentTransactionRepository *repositories.PaymentTransactionRepository) *PaymentTransactionService {
// 	return &PaymentTransactionService{
// 		paymentTransactionRepository: paymentTransactionRepository,
// 	}
// }

func NewPaymentTransactionService(repoMgr *repositories.RepositoriesManager) *PaymentTransactionService {
	return &PaymentTransactionService{
		repositoriesManager: *repoMgr,
	}
}

// func (ptr PaymentTransactionService) GetListPaymentTransaction(ctx *gin.Context) ([]*dbContext.TransactionUser, *models.ResponseError) {
// 	return ptr.paymentTransactionRepository.GetListPaymentTransaction(ctx)
// }

func (pts PaymentTransactionService) GetListPaymentTransaction(ctx *gin.Context) ([]*dbContext.RecordTransactionUser, *models.ResponseError) {
	// return pts.paymentTransactionRepository.GetListPaymentTransaction(ctx)

	return pts.repositoriesManager.PaymentTransactionRepository.GetListPaymentTransaction(ctx)
}

func (pts PaymentTransactionService) GetPaymentTransactionById(ctx *gin.Context, metadata *features.Metadata) ([]*dbContext.RecordTransactionUser, *models.ResponseError) {
	return pts.repositoriesManager.PaymentTransactionRepository.GetPaymentTransactionById(ctx, metadata)
}

func (pts PaymentTransactionService) RecordPaymentTransactionUser(ctx *gin.Context, params *dbContext.RecordTransactionUserParams) (*dbContext.RecordTransactionUser, *models.ResponseError) {
	responseErr := validatePaymentTransaction(params)
	if responseErr != nil {
		return nil, responseErr
	}

	return pts.repositoriesManager.PaymentTransactionRepository.RecordPaymentTransactionUser(ctx, params)
}

// func (ptr PaymentTransactionService) UpdatePaymentTransaction(ctx *gin.Context, paymentTransactionParams *dbContext.CreatePaymentTransaction_paymentParams, id int64) *models.ResponseError {
// 	responseErr := validatePaymentTransaction(paymentTransactionParams)
// 	if responseErr != nil {
// 		return responseErr
// 	}
// 	return ptr.paymentTransactionRepository.UpdatePaymentTransaction(ctx, paymentTransactionParams)

// }

func (ptr PaymentTransactionService) DeletePaymentTransaction(ctx *gin.Context, trpaCodeNumber string) *models.ResponseError {
	return ptr.repositoriesManager.PaymentTransactionRepository.DeletePaymentTransaction(ctx, trpaCodeNumber)
}

func validatePaymentTransaction(params *dbContext.RecordTransactionUserParams) *models.ResponseError {
	if params.TrpaUserEntityID == 0 {
		return &models.ResponseError{
			Message: "Invalid TrpaUserEntityID",
			Status:  http.StatusBadRequest,
		}
	}
	return nil

}
