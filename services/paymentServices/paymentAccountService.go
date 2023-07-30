package services

import (
	"net/http"

	"codeid.revampacademy/models"
	repositories "codeid.revampacademy/repositories/paymentRepositories"
	"codeid.revampacademy/repositories/paymentRepositories/dbContext"
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

func (pas PaymentAccountService) GetListPaymentAccount(ctx *gin.Context) ([]*dbContext.UserAccount, *models.ResponseError) {
	return pas.paymentAccountRepository.GetListPaymentAccount(ctx)
}

func (pas PaymentAccountService) GetPaymentAccountByName(ctx *gin.Context, name string) (*dbContext.UserAccount, *models.ResponseError) {
	return pas.paymentAccountRepository.GetPaymentAccountByName(ctx, name)
}

func (pas PaymentAccountService) CreateNewPaymentAccount(ctx *gin.Context, paymentAccountParams *dbContext.CreatePaymentUsers_accountParams) (*models.PaymentUsersAccount, *models.ResponseError) {
	responseErr := validatepaymentAccount(paymentAccountParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return pas.paymentAccountRepository.CreateNewPaymentAccount(ctx, paymentAccountParams)
}

func (pas PaymentAccountService) UpdatePaymentAccountByAccNum(ctx *gin.Context, paymentAccountParams *dbContext.CreatePaymentUsers_accountParams, accNum string) *models.ResponseError {
	responseErr := validatepaymentAccount(paymentAccountParams)
	if responseErr != nil {
		return responseErr
	}

	return pas.paymentAccountRepository.UpdatePaymentAccountByAccNum(ctx, paymentAccountParams)
}

func (pas PaymentAccountService) DeletePaymentAccountByAccNum(ctx *gin.Context, accNum string) *models.ResponseError {
	return pas.paymentAccountRepository.DeletePaymentAccountByAccNum(ctx, accNum)
}

func validatepaymentAccount(paymentAccountParams *dbContext.CreatePaymentUsers_accountParams) *models.ResponseError {
	if paymentAccountParams.UsacAccountNumber == "" {
		return &models.ResponseError{
			Message: "Invalid paymentAccount id",
			Status:  http.StatusBadRequest,
		}
	}

	if paymentAccountParams.UsacSaldo == 0 {
		return &models.ResponseError{
			Message: "Invalid paymentAccount name",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}
