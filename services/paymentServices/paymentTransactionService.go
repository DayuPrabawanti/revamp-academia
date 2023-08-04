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

//	func NewPaymentTransactionService(paymentTransactionRepository *repositories.PaymentTransactionRepository) *PaymentTransactionService {
//		return &PaymentTransactionService{
//			paymentTransactionRepository: paymentTransactionRepository,
//		}
//	}
func NewPaymentTransactionService(repoMgr *repositories.RepositoriesManager) *PaymentTransactionService {
	return &PaymentTransactionService{
		repositoriesManager: *repoMgr,
	}
}

func (pts PaymentTransactionService) GetListPaymentTransaction(ctx *gin.Context) ([]*dbContext.TransactionUser, *models.ResponseError) {
	// return pts.paymentTransactionRepository.GetListPaymentTransaction(ctx)

	return pts.repositoriesManager.PaymentTransactionRepository.GetListPaymentTransaction(ctx)
}

func (pts PaymentTransactionService) GetPaymentTransactionById(ctx *gin.Context, metadata *features.Metadata) ([]*dbContext.TransactionUser, *models.ResponseError) {
	return pts.repositoriesManager.PaymentTransactionRepository.GetPaymentTransactionById(ctx, metadata)
}

func (pts PaymentTransactionService) CreateNewPaymentTransaction(ctx *gin.Context, paymentTransactionParams *dbContext.CreateTransactionUser) (*dbContext.TransactionUser, *models.ResponseError) {
	responseErr := validatePaymentTransaction(paymentTransactionParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return pts.repositoriesManager.PaymentTransactionRepository.CreatePaymentTransaction(ctx, paymentTransactionParams)
}

// func (ptr PaymentTransactionService) UpdatePaymentTransaction(ctx *gin.Context, paymentTransactionParams *dbContext.CreatePaymentTransaction_paymentParams, id int64) *models.ResponseError {
// 	responseErr := validatePaymentTransaction(paymentTransactionParams)
// 	if responseErr != nil {
// 		return responseErr
// 	}
// 	return ptr.paymentTransactionRepository.UpdatePaymentTransaction(ctx, paymentTransactionParams)

// }

// func (ptr PaymentTransactionService) DeletePaymentTransaction(ctx *gin.Context, id int64) *models.ResponseError {
// 	return ptr.paymentTransactionRepository.DeletePaymentTransaction(ctx, id)
// }

func validatePaymentTransaction(paymentTransactionParams *dbContext.CreateTransactionUser) *models.ResponseError {
	if paymentTransactionParams.TrpaUserEntityID == 0 {
		return &models.ResponseError{
			Message: "Invalid TrpaCodeNumber",
			Status:  http.StatusBadRequest,
		}
	}
	return nil

}
