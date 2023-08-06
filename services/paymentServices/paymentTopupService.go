package services

import (
	"database/sql"
	"net/http"

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

func (ptts PaymentTopupService) TransferTopup(ctx *gin.Context, bankCode string, accountFrom string, fintCode string, accountTo string, amount float64) (*dbContext.RecordTransactionUser, *models.ResponseError) {

	err := repositories.BeginTransaction(ptts.repositoriesManager)
	if err != nil {
		return nil, &models.ResponseError{
			Message: "Failed to start transaction",
			Status:  http.StatusBadRequest,
		}
	}

	// First, get user account information
	userAccountBank, responseErr := ptts.GetAccountByBankCodeAndAccountNumber(ctx, bankCode, accountFrom)
	if responseErr != nil {
		repositories.RollbackTransaction(ptts.repositoriesManager)
		return nil, responseErr
	}

	userAccountFint, responseErr := ptts.GetAccountByFintCodeAndAccountNumber(ctx, fintCode, accountTo)
	if responseErr != nil {
		repositories.RollbackTransaction(ptts.repositoriesManager)
		return nil, responseErr
	}

	// Update user's account with the debit amount
	updateParams := &dbContext.UpdatePaymentUsers_accountParams{
		Amount: amount,
	}
	_, responseErr = ptts.repositoriesManager.PaymentAccountRepository.UpdatePaymentUsers_accountMinus(ctx, updateParams, accountFrom)
	if responseErr != nil {
		repositories.RollbackTransaction(ptts.repositoriesManager)
		return nil, &models.ResponseError{
			Message: "Failed to update account balance",
			Status:  http.StatusInternalServerError,
		}
	}
	_, responseErr = ptts.repositoriesManager.PaymentAccountRepository.UpdatePaymentUsers_accountPlus(ctx, updateParams, accountTo)
	if responseErr != nil {
		repositories.RollbackTransaction(ptts.repositoriesManager)
		return nil, &models.ResponseError{
			Message: "Failed to update account balance",
			Status:  http.StatusInternalServerError,
		}
	}

	// Create a transaction record
	transactionUserBank := &dbContext.RecordTransactionUserParams{
		TrpaDebit:        sql.NullFloat64{},
		TrpaCredit:       sql.NullFloat64{Float64: amount, Valid: true},
		TrpaType:         "TR",
		TrpaNote:         "Transfer",
		TrpaFromID:       accountFrom,
		TrpaToID:         accountTo,
		TrpaUserEntityID: userAccountBank.UserEntityID,
	}
	_, responseErr = ptts.repositoriesManager.PaymentTransactionRepository.RecordPaymentTransactionUser(ctx, transactionUserBank)
	if responseErr != nil {
		repositories.RollbackTransaction(ptts.repositoriesManager)
		return nil, responseErr
	}

	transactionUserFint := &dbContext.RecordTransactionUserParams{
		TrpaDebit:        sql.NullFloat64{Float64: amount, Valid: true},
		TrpaCredit:       sql.NullFloat64{},
		TrpaType:         "TP",
		TrpaNote:         "Topup",
		TrpaFromID:       accountFrom,
		TrpaToID:         accountTo,
		TrpaUserEntityID: userAccountFint.UserEntityID,
	}
	_, responseErr = ptts.repositoriesManager.PaymentTransactionRepository.RecordPaymentTransactionUser(ctx, transactionUserFint)
	if responseErr != nil {
		repositories.RollbackTransaction(ptts.repositoriesManager)
		return nil, responseErr
	}

	repositories.CommitTransaction(ptts.repositoriesManager)

	return nil, &models.ResponseError{
		Message: "Transfer & Topup Successfull",
		Status:  http.StatusOK,
	}
}
