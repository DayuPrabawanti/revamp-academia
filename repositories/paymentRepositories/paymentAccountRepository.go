package repositories

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/paymentRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type PaymentAccountRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewPaymentAccountRepository(dbHandler *sql.DB) *PaymentAccountRepository {
	return &PaymentAccountRepository{
		dbHandler: dbHandler,
	}
}

func (par PaymentAccountRepository) GetListPaymentUsers_accountByUserName(ctx *gin.Context, userName string) ([]*dbContext.UserAccount, *models.ResponseError) {

	store := dbContext.New(par.dbHandler)
	paymentAccounts, err := store.ListPaymentUsers_accountByUserName(ctx, userName)

	listPaymentAccounts := make([]*dbContext.UserAccount, 0)

	for _, v := range paymentAccounts {
		paymentAccount := &dbContext.UserAccount{
			UserName:      v.UserName,
			AccountNumber: v.AccountNumber,
			Description:   v.Description,
			Saldo:         v.Saldo,
			Type:          v.Type,
		}
		listPaymentAccounts = append(listPaymentAccounts, paymentAccount)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listPaymentAccounts, nil
}

func (par PaymentAccountRepository) GetPaymentAccountByAccountNumber(ctx *gin.Context, usacAccountNumber string) (*dbContext.UserAccount, *models.ResponseError) {

	store := dbContext.New(par.dbHandler)
	paymentAccount, err := store.GetPaymentUsers_account(ctx, usacAccountNumber)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &paymentAccount, nil
}

func (par PaymentAccountRepository) CreateNewPaymentAccount(ctx *gin.Context, paymentAccountParams *dbContext.CreatePaymentUsers_accountParams) (*dbContext.UserAccount, *models.ResponseError) {

	store := dbContext.New(par.dbHandler)
	paymentAccount, err := store.CreatePaymentUsers_account(ctx, *paymentAccountParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return paymentAccount, nil
}

func (par *PaymentAccountRepository) UpdatePaymentUsers_accountPlus(ctx *gin.Context, params *dbContext.UpdatePaymentUsers_accountParams, usacAccountNumber string) (*dbContext.UserAccount, *models.ResponseError) {
	store := dbContext.New(par.dbHandler)
	paymentAccount, err := store.UpdatePaymentUsers_accountPlus(ctx, *params, usacAccountNumber)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return paymentAccount, nil
}

func (par *PaymentAccountRepository) UpdatePaymentUsers_accountMinus(ctx *gin.Context, params *dbContext.UpdatePaymentUsers_accountParams, usacAccountNumber string) (*dbContext.UserAccount, *models.ResponseError) {
	store := dbContext.New(par.dbHandler)
	paymentAccount, err := store.UpdatePaymentUsers_accountMinus(ctx, *params, usacAccountNumber)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return paymentAccount, nil
}

func (par PaymentAccountRepository) DeletePaymentAccountByAccNum(ctx *gin.Context, usacAccountNumber string) *models.ResponseError {

	store := dbContext.New(par.dbHandler)
	err := store.DeletePaymentUsers_account(ctx, usacAccountNumber)

	if err != nil {
		return &models.ResponseError{
			Message: "error when delete",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been deleted",
		Status:  http.StatusOK,
	}
}
