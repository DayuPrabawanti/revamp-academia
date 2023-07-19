package repositories

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/dbContext"
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

func (par PaymentAccountRepository) GetListPaymentAccount(ctx *gin.Context) ([]*models.PaymentUsersAccount, *models.ResponseError) {

	store := dbContext.New(par.dbHandler)
	paymentAccounts, err := store.ListPaymentUsers_account(ctx)

	listPaymentAccounts := make([]*models.PaymentUsersAccount, 0)

	for _, v := range paymentAccounts {
		paymentAccount := &models.PaymentUsersAccount{
			UsacBankEntityID:  v.UsacBankEntityID,
			UsacUserEntityID:  v.UsacUserEntityID,
			UsacAccountNumber: v.UsacAccountNumber,
			UsacSaldo:         v.UsacSaldo,
			UsacType:          v.UsacType,
			UsacStartDate:     v.UsacStartDate,
			UsacEndDate:       v.UsacEndDate,
			UsacModifiedDate:  v.UsacModifiedDate,
			UsacStatus:        v.UsacStatus,
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

func (par PaymentAccountRepository) GetPaymentAccountByName(ctx *gin.Context, name string) (*models.PaymentUsersAccount, *models.ResponseError) {

	store := dbContext.New(par.dbHandler)
	paymentAccount, err := store.GetPaymentUsers_account(ctx, string(name))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &paymentAccount, nil
}

func (par PaymentAccountRepository) CreateNewPaymentAccount(ctx *gin.Context, paymentAccountParams *dbContext.CreatePaymentUsers_accountParams) (*models.PaymentUsersAccount, *models.ResponseError) {

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

func (par PaymentAccountRepository) UpdatePaymentAccountById(ctx *gin.Context, paymentAccountParams *dbContext.CreatePaymentUsers_accountParams) *models.ResponseError {

	store := dbContext.New(par.dbHandler)
	err := store.UpdatePaymentUsers_account(ctx, *paymentAccountParams)

	if err != nil {
		return &models.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been update",
		Status:  http.StatusOK,
	}
}

func (par PaymentAccountRepository) DeletePaymentAccountById(ctx *gin.Context, id int64) *models.ResponseError {

	store := dbContext.New(par.dbHandler)
	err := store.DeletePaymentUsers_account(ctx, int32(id))

	if err != nil {
		return &models.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been deleted",
		Status:  http.StatusOK,
	}
}
