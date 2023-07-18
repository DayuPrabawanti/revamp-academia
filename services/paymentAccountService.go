package services

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories"
	"codeid.revampacademy/repositories/dbContext"
	"github.com/gin-gonic/gin"
)

type PaymentAccountService struct {
	paymentAccountRepository *repositories.PaymentAccountRepository
}

func NewPaymentAccountService(paymentAccountRepository *repositories.PaymentAccountRepository) *PaymentAccountService {
	return &PaymentAccountService{
		paymentAccountRepository: paymentAccountRepository,
	}
}

func (pas PaymentAccountService) GetListPaymentAccount(ctx *gin.Context) ([]*models.PaymentUsersAccount, *models.ResponseError) {
	return pas.paymentAccountRepository.GetListPaymentAccount(ctx)
}

func (pas PaymentAccountService) GetPaymentAccountByName(ctx *gin.Context, name string) (*models.PaymentUsersAccount, *models.ResponseError) {
	return pas.paymentAccountRepository.GetPaymentAccountByName(ctx, name)
}

func (pas PaymentAccountService) CreateNewPaymentAccount(ctx *gin.Context, paymentAccountParams *dbContext.CreatePaymentUsers_accountParams) (*models.PaymentUsersAccount, *models.ResponseError) {
	responseErr := validatepaymentAccount(paymentAccountParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return pas.paymentAccountRepository.CreateNewPaymentAccount(ctx, paymentAccountParams)
}

func (pas PaymentAccountService) UpdatePaymentAccountById(ctx *gin.Context, paymentAccountParams *dbContext.CreatePaymentUsers_accountParams, id int64) *models.ResponseError {
	responseErr := validatepaymentAccount(paymentAccountParams)
	if responseErr != nil {
		return responseErr
	}

	return pas.paymentAccountRepository.UpdatePaymentAccountById(ctx, paymentAccountParams)
}

func (pas PaymentAccountService) DeletePaymentAccountById(ctx *gin.Context, id int64) *models.ResponseError {
	return pas.paymentAccountRepository.DeletePaymentAccountById(ctx, id)
}

func validatepaymentAccount(paymentAccountParams *dbContext.CreatePaymentUsers_accountParams) *models.ResponseError {
	if paymentAccountParams.UsacBankEntityID == 0 {
		return &models.ResponseError{
			Message: "Invalid paymentAccount id",
			Status:  http.StatusBadRequest,
		}
	}

	if paymentAccountParams.UsacAccountNumber.String == "" {
		return &models.ResponseError{
			Message: "Invalid paymentAccount name",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}
