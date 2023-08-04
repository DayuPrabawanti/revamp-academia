package services

import (
	"net/http"

	"codeid.revampacademy/models"
	repositories "codeid.revampacademy/repositories/paymentRepositories"
	"codeid.revampacademy/repositories/paymentRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type PaymentFintechService struct {
	repositoriesManager *repositories.RepositoriesManager
}

func NewPaymentFintechService(repoMgr *repositories.RepositoriesManager) *PaymentFintechService {
	return &PaymentFintechService{
		repositoriesManager: repoMgr,
	}
}

func (pfs PaymentFintechService) GetListPaymentFintech(ctx *gin.Context) ([]*dbContext.Fintech, *models.ResponseError) {
	return pfs.repositoriesManager.PaymentFintechRepository.GetListPaymentFintech(ctx)
}

func (pfs PaymentFintechService) GetPaymentFintechByName(ctx *gin.Context, name string) (*dbContext.Fintech, *models.ResponseError) {
	return pfs.repositoriesManager.PaymentFintechRepository.GetPaymentFintechByName(ctx, name)
}

func (pfs PaymentFintechService) CreateNewPaymentFintech(ctx *gin.Context, paymentFintechParams *dbContext.CreatePaymentFintechParams) (*dbContext.Fintech, *models.ResponseError) {
	responseErr := validatepaymentFintech(paymentFintechParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return pfs.repositoriesManager.PaymentFintechRepository.CreateNewPaymentFintech(ctx, paymentFintechParams)
}

func (pfs PaymentFintechService) UpdatePaymentFintechById(ctx *gin.Context, paymentFintechParams *dbContext.CreatePaymentFintechParams, fintEntityID int64) *models.ResponseError {
	responseErr := validatepaymentFintech(paymentFintechParams)
	if responseErr != nil {
		return responseErr
	}

	return pfs.repositoriesManager.PaymentFintechRepository.UpdatePaymentFintechById(ctx, paymentFintechParams, fintEntityID)
}

func (pfs PaymentFintechService) DeletePaymentFintechById(ctx *gin.Context, id int64) *models.ResponseError {
	return pfs.repositoriesManager.PaymentFintechRepository.DeletePaymentFintechById(ctx, id)
}

func validatepaymentFintech(paymentFintechParams *dbContext.CreatePaymentFintechParams) *models.ResponseError {

	if paymentFintechParams.FintCode == "" {
		return &models.ResponseError{
			Message: "Invalid paymentFintech name",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}
