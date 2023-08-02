package services

import (
	"codeid.revampacademy/models"
	repositories "codeid.revampacademy/repositories/paymentRepositories"
	"codeid.revampacademy/repositories/paymentRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type PaymentTopupService struct {
	repositoriesManager *repositories.RepositoriesManager
}

func NewPaymentTopupService(repoMgr *repositories.RepositoriesManager) *PaymentTopupService {
	return &PaymentTopupService{
		repositoriesManager: repoMgr,
	}
}

func (ptts *PaymentTopupService) GetAccountByBankCodeAndAccountNumber(ctx *gin.Context, bankCode string, usacAccountNumber string) (*dbContext.BankAccount, *models.ResponseError) {
	return ptts.repositoriesManager.PaymentTopupRepository.GetAccountByBankCodeAndAccountNumber(ctx, bankCode, usacAccountNumber)
}

func (ptts *PaymentTopupService) GetAccountByFintCodeAndAccountNumber(ctx *gin.Context, fintCode string, usacAccountNumber string) (*dbContext.FintechAccount, *models.ResponseError) {
	return ptts.repositoriesManager.PaymentTopupRepository.GetAccountByFintCodeAndAccountNumber(ctx, fintCode, usacAccountNumber)
}
