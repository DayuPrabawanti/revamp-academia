package services

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories"
	"codeid.revampacademy/repositories/dbContext"
	"github.com/gin-gonic/gin"
)

type PaymentFintechService struct {
	paymentFintechRepository *repositories.PaymentFintechRepository
}

func NewPaymentFintechService(paymentFintechRepository *repositories.PaymentFintechRepository) *PaymentFintechService {
	return &PaymentFintechService{
		paymentFintechRepository: paymentFintechRepository,
	}
}

func (pfs PaymentFintechService) GetListPaymentFintech(ctx *gin.Context) ([]*models.PaymentFintech, *models.ResponseError) {
	return pfs.paymentFintechRepository.GetListPaymentFintech(ctx)
}

func (pfs PaymentFintechService) GetPaymentFintechByName(ctx *gin.Context, name string) (*models.PaymentFintech, *models.ResponseError) {
	return pfs.paymentFintechRepository.GetPaymentFintechByName(ctx, name)
}

func (pfs PaymentFintechService) CreateNewPaymentFintech(ctx *gin.Context, paymentFintechParams *dbContext.CreatePaymentFintechParams) (*models.PaymentFintech, *models.ResponseError) {
	responseErr := validatepaymentFintech(paymentFintechParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return pfs.paymentFintechRepository.CreateNewPaymentFintech(ctx, paymentFintechParams)
}

func (pfs PaymentFintechService) UpdatePaymentFintechById(ctx *gin.Context, paymentFintechParams *dbContext.CreatePaymentFintechParams, id int64) *models.ResponseError {
	responseErr := validatepaymentFintech(paymentFintechParams)
	if responseErr != nil {
		return responseErr
	}

	return pfs.paymentFintechRepository.UpdatePaymentFintechById(ctx, paymentFintechParams)
}

func (pfs PaymentFintechService) DeletePaymentFintechById(ctx *gin.Context, id int64) *models.ResponseError {
	return pfs.paymentFintechRepository.DeletePaymentFintechById(ctx, id)
}

func validatepaymentFintech(paymentFintechParams *dbContext.CreatePaymentFintechParams) *models.ResponseError {
	if paymentFintechParams.FintEntityID == 0 {
		return &models.ResponseError{
			Message: "Invalid paymentFintech id",
			Status:  http.StatusBadRequest,
		}
	}

	if paymentFintechParams.FintName.String == "" {
		return &models.ResponseError{
			Message: "Invalid paymentFintech name",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}
