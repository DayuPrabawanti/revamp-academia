package services

import (
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

// func (pas PaymentAccountService) DebitSaldo(ctx *gin.Context, usacAccountNumber string, amount float64, note string) (*dbContext.TransactionUser, *models.ResponseError) {

// 	err := repositories.BeginTransaction(pas.repositoryManager)
// 	if err != nil {
// 		return nil, &models.ResponseError{
// 			Message: "Failed to start transaction",
// 			Status:  http.StatusBadRequest,
// 		}
// 	}

// 	// First, get user account information
// 	userAccount, err := ps.GetPaymentUsers_account(ctx, usacAccountNumber)
// 	if err != nil {
// 		repositories.RollbackTransaction(ps.repositoryManager)
// 		return nil, &models.ResponseError{
// 			Message: "Failed to retrieve account information",
// 			Status:  http.StatusNotFound,
// 		}
// 	}

// 	// Update user's account with the debit amount
// 	updateParams := UpdatePaymentUsers_accountParams{
// 		Amount: amount,
// 	}
// 	_, err = ps.UpdatePaymentUsers_accountPlus(ctx, updateParams, usacAccountNumber)
// 	if err != nil {
// 		repositories.RollbackTransaction(ps.repositoryManager)
// 		return nil, &models.ResponseError{
// 			Message: "Failed to update account balance",
// 			Status:  http.StatusInternalServerError,
// 		}
// 	}

// 	// Create a transaction record
// 	transactionParams := CreateTransactionUser{
// 		TrpaDebit:        sql.NullFloat64{Float64: amount, Valid: true},
// 		TrpaNote:         note,
// 		TrpaFromID:       usacAccountNumber,       // adjust according to your logic
// 		TrpaToID:         "",                      // adjust according to your logic
// 		TrpaUserEntityID: int32(userAccount.Type), // adjust according to your logic
// 	}
// 	transaction, err := ps.CreatePaymentTransaction_payment(ctx, transactionParams)
// 	if err != nil {
// 		repositories.RollbackTransaction(ps.repositoryManager)
// 		return nil, &models.ResponseError{
// 			Message: "Failed to record transaction",
// 			Status:  http.StatusInternalServerError,
// 		}
// 	}

// 	repositories.CommitTransaction(ps.repositoryManager)

// 	return transaction, nil
// }
