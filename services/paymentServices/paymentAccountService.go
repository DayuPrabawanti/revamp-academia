package services

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	repositories "codeid.revampacademy/repositories/paymentRepositories"
	"codeid.revampacademy/repositories/paymentRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type PaymentAccountService struct {
	// paymentAccountRepository *repositories.PaymentAccountRepository

	repositoriesManager *repositories.RepositoriesManager
}

//	func NewPaymentAccountService(paymentAccountRepository *repositories.PaymentAccountRepository) *PaymentAccountService {
//		return &PaymentAccountService{
//			paymentAccountRepository: paymentAccountRepository,
//		}
//	}
func NewPaymentAccountService(repoMgr *repositories.RepositoriesManager) *PaymentAccountService {
	return &PaymentAccountService{
		repositoriesManager: repoMgr,
	}
}

func (pas PaymentAccountService) GetListPaymentUsers_accountByUserName(ctx *gin.Context, userName string) ([]*dbContext.UserAccount, *models.ResponseError) {
	// return pas.paymentAccountRepository.GetListPaymentUsers_accountByUserName(ctx, userName)

	return pas.repositoriesManager.PaymentAccountRepository.GetListPaymentUsers_accountByUserName(ctx, userName)
}

func (pas PaymentAccountService) GetPaymentAccountByAccountNumber(ctx *gin.Context, usacAccountNumber string) (*dbContext.UserAccount, *models.ResponseError) {
	return pas.repositoriesManager.PaymentAccountRepository.GetPaymentAccountByAccountNumber(ctx, usacAccountNumber)
}

func (pas PaymentAccountService) CreateNewPaymentAccount(ctx *gin.Context, paymentAccountParams *dbContext.CreatePaymentUsers_accountParams) (*dbContext.UserAccount, *models.ResponseError) {
	responseErr := validatepaymentAccount(paymentAccountParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return pas.repositoriesManager.PaymentAccountRepository.CreateNewPaymentAccount(ctx, paymentAccountParams)
}

func validatepaymentAccount(paymentAccountParams *dbContext.CreatePaymentUsers_accountParams) *models.ResponseError {
	if paymentAccountParams.UsacBankEntityID == 0 {
		return &models.ResponseError{
			Message: "Invalid usacBankEntityID",
			Status:  http.StatusBadRequest,
		}
	}

	if paymentAccountParams.UsacUserEntityID == 0 {
		return &models.ResponseError{
			Message: "Invalid usacUserEntityID",
			Status:  http.StatusBadRequest,
		}
	}

	if paymentAccountParams.UsacAccountNumber == "" {
		return &models.ResponseError{
			Message: "Invalid usacAccountNumber",
			Status:  http.StatusBadRequest,
		}
	}

	if paymentAccountParams.UsacSaldo == 0 {
		return &models.ResponseError{
			Message: "Invalid usacSaldo / usacSaldo must > 0",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}

func (pas *PaymentAccountService) UpdatePaymentUsers_accountPlus(ctx *gin.Context, params *dbContext.UpdatePaymentUsers_accountParams, usacAccountNumber string) (*dbContext.UserAccount, *models.ResponseError) {
	responseErr := validateUpdatePayment(params, usacAccountNumber)
	if responseErr != nil {
		return nil, responseErr
	}

	return pas.repositoriesManager.PaymentAccountRepository.UpdatePaymentUsers_accountPlus(ctx, params, usacAccountNumber)
}

func validateUpdatePayment(params *dbContext.UpdatePaymentUsers_accountParams, usacAccountNumber string) *models.ResponseError {
	if params.Amount <= 0 {
		return &models.ResponseError{
			Message: "Invalid amount / amount must be > 0",
			Status:  http.StatusBadRequest,
		}
	}

	if usacAccountNumber == "" {
		return &models.ResponseError{
			Message: "Invalid usacAccountNumber",
			Status:  http.StatusBadRequest,
		}
	}

	return nil
}

func (pas *PaymentAccountService) UpdatePaymentUsers_accountMinus(ctx *gin.Context, params *dbContext.UpdatePaymentUsers_accountParams, usacAccountNumber string) (*dbContext.UserAccount, *models.ResponseError) {
	responseErr := validateUpdatePayment(params, usacAccountNumber)
	if responseErr != nil {
		return nil, responseErr
	}

	return pas.repositoriesManager.PaymentAccountRepository.UpdatePaymentUsers_accountMinus(ctx, params, usacAccountNumber)
}

func (pas PaymentAccountService) DeletePaymentAccountByAccNum(ctx *gin.Context, usacAccountNumber string) *models.ResponseError {
	return pas.repositoriesManager.PaymentAccountRepository.DeletePaymentAccountByAccNum(ctx, usacAccountNumber)
}

func (pas PaymentAccountService) RecordTransaction(ctx *gin.Context, params *dbContext.CreateTransactionUserParam) (*dbContext.TransactionUserDebit, *models.ResponseError) {
	responseErr := validateRecordTransaction(params)
	if responseErr != nil {
		return nil, responseErr
	}

	return pas.repositoriesManager.PaymentAccountRepository.RecordPaymentTransaction(ctx, params)
}

func validateRecordTransaction(paymentTransactionParams *dbContext.CreateTransactionUserParam) *models.ResponseError {
	if paymentTransactionParams.TrpaUserEntityID == 0 {
		return &models.ResponseError{
			Message: "Invalid TrpaCodeNumber",
			Status:  http.StatusBadRequest,
		}
	}
	return nil

}

func (pas PaymentAccountService) DebitSaldo(ctx *gin.Context, usacAccountNumber string, amount float64) (*dbContext.TransactionUserDebit, *models.ResponseError) {

	err := repositories.BeginTransaction(pas.repositoriesManager)
	if err != nil {
		return nil, &models.ResponseError{
			Message: "Failed to start transaction",
			Status:  http.StatusBadRequest,
		}
	}

	// Pertama, dapatkan informasi akun pengguna
	userAccount, responseErr := pas.GetPaymentAccountByAccountNumber(ctx, usacAccountNumber)
	if responseErr != nil {
		repositories.RollbackTransaction(pas.repositoriesManager) // Gantikan dengan metode yang benar untuk rollback transaksi
		return nil, responseErr
	}

	// Update akun pengguna dengan jumlah yang ingin ditambahkan ke saldo
	updateParams := &dbContext.UpdatePaymentUsers_accountParams{
		Amount: amount, // Jumlah yang ingin ditambahkan ke saldo
	}
	_, responseErr = pas.UpdatePaymentUsers_accountPlus(ctx, updateParams, usacAccountNumber)
	if responseErr != nil {
		repositories.RollbackTransaction(pas.repositoriesManager)
		return nil, responseErr
	}

	// Buat catatan transaksi
	transactionParams := &dbContext.CreateTransactionUserParam{
		TrpaDebit:        sql.NullFloat64{Float64: amount, Valid: true},
		TrpaCredit:       sql.NullFloat64{},
		TrpaType:         "SD",
		TrpaNote:         "Isi SALDO",
		TrpaFromID:       usacAccountNumber,
		TrpaToID:         "",                       // Sesuaikan sesuai dengan logika Anda
		TrpaUserEntityID: userAccount.UserEntityID, // Mengambil dari informasi akun pengguna
	}
	transaction, responseErr := pas.RecordTransaction(ctx, transactionParams) // Pastikan memanggil metode yang benar
	if responseErr != nil {
		repositories.RollbackTransaction(pas.repositoriesManager)
		return nil, responseErr
	}

	repositories.CommitTransaction(pas.repositoriesManager) // Gantikan dengan metode yang benar untuk commit transaksi

	return transaction, nil
}
