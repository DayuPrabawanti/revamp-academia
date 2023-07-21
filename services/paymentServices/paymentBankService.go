package paymentServices

import (
	"net/http"

	"codeid.revampacademy/models"
	repositories "codeid.revampacademy/repositories/paymentRepositories"
	"codeid.revampacademy/repositories/paymentRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type PaymentBankService struct {
	paymentBankRepository *repositories.PaymentBankRepository
}

func NewPaymentBankService(paymentBankRepository *repositories.PaymentBankRepository) *PaymentBankService {
	return &PaymentBankService{
		paymentBankRepository: paymentBankRepository,
	}
}

// 1a. ambil get list untuk payment bank
func (pbr PaymentBankService) GetListPaymentBank(ctx *gin.Context) ([]*models.PaymentBank, *models.ResponseError) {
	return pbr.paymentBankRepository.GetListPaymentBank(ctx)
}

// 1a. ambil get by name untuk payment bank
func (pbr PaymentBankService) GetPaymentBankByName(ctx *gin.Context, name string) (*models.PaymentBank, *models.ResponseError) {
	return pbr.paymentBankRepository.GetPaymentBankByName(ctx, name)
}

// 1b. buat create paymentbank
func (pbr PaymentBankService) CreateNewPaymentBank(ctx *gin.Context, paymentBankParams *dbContext.CreatePaymentBankParams) (*models.PaymentBank, *models.ResponseError) {
	responseErr := validatePaymentBank(paymentBankParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return pbr.paymentBankRepository.CreateNewPaymentBank(ctx, paymentBankParams)
}

// 1b. update payment data bank
func (pbr PaymentBankService) UpdatePaymentBank(ctx *gin.Context, paymentBankParams *dbContext.CreatePaymentBankParams, id int64) *models.ResponseError {
	responseErr := validatePaymentBank(paymentBankParams)
	if responseErr != nil {
		return responseErr
	}
	return pbr.paymentBankRepository.UpdatePaymentBank(ctx, paymentBankParams)

}

// 1b. delet payment bank
func (pbr PaymentBankService) DeletePaymentBank(ctx *gin.Context, id int64) *models.ResponseError {
	return pbr.paymentBankRepository.DeletePaymentBank(ctx, id)
}

func validatePaymentBank(paymentBankParams *dbContext.CreatePaymentBankParams) *models.ResponseError {
	if paymentBankParams.BankEntityID == 0 {
		return &models.ResponseError{
			Message: "Invalid category id",
			Status:  http.StatusBadRequest,
		}
	}

	if paymentBankParams.BankName == "" {
		return &models.ResponseError{
			Message: "Invalid category name",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}
