package repositories

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/paymentRepositories/dbContext"
	"golang.org/x/net/context"
)

type PaymentTopupRepository struct {
	dbHandler *sql.DB
	// transaction *sql.Tx
}

func NewPaymentTopupRepository(dbHandler *sql.DB) *PaymentTopupRepository {
	return &PaymentTopupRepository{
		dbHandler: dbHandler,
	}
}

func (pttr *PaymentTopupRepository) GetAccountByBankCodeAndAccountNumber(ctx context.Context, bankCode string, usacAccountNumber string) (*dbContext.BankAccount, *models.ResponseError) {
	store := dbContext.New(pttr.dbHandler)
	bankAccount, err := store.GetBankAccount(ctx, bankCode, usacAccountNumber)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &bankAccount, nil
}

func (pttr *PaymentTopupRepository) GetAccountByFintCodeAndAccountNumber(ctx context.Context, fintCode string, usacAccountNumber string) (*dbContext.FintechAccount, *models.ResponseError) {
	store := dbContext.New(pttr.dbHandler)
	fintAccount, err := store.GetFintechAccount(ctx, fintCode, usacAccountNumber)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &fintAccount, nil
}
