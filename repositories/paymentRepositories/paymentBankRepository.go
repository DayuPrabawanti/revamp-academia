package repositories

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/paymentRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type PaymentBankRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewPaymentBankRepository(dbHandler *sql.DB) *PaymentBankRepository {
	return &PaymentBankRepository{
		dbHandler: dbHandler,
	}
}

// 1a. fungsi utk ambil get list
func (pbr PaymentBankRepository) GetListPaymentBank(ctx *gin.Context) ([]*dbContext.Bank, *models.ResponseError) {
	store := dbContext.New(pbr.dbHandler)
	paymentBanks, err := store.ListPaymentBank(ctx)
	listPaymentBanks := make([]*dbContext.Bank, 0)

	for _, v := range paymentBanks {
		paymentBank := &dbContext.Bank{
			BankEntityID: v.BankEntityID,
			BankCode:     v.BankCode,
		}
		listPaymentBanks = append(listPaymentBanks, paymentBank)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return listPaymentBanks, nil
}

// 1a. fungsi utk ambil get by name
func (pbr PaymentBankRepository) GetPaymentBankByName(ctx *gin.Context, name string) (*dbContext.Bank, *models.ResponseError) {
	store := dbContext.New(pbr.dbHandler)
	paymentBank, err := store.GetPaymentBank(ctx, name)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &paymentBank, nil
}

// 1b. fungsi utk create data payment
func (pbr PaymentBankRepository) CreateNewPaymentBank(ctx *gin.Context, paymentBankParams *dbContext.CreatePaymentBankParams) (*dbContext.Bank, *models.ResponseError) {
	store := dbContext.New(pbr.dbHandler)
	paymentBank, err := store.CreatePaymentBank(ctx, *paymentBankParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return paymentBank, nil
}

func (cr PaymentBankRepository) UpdatePaymentBank(ctx *gin.Context, paymentBankParams *dbContext.CreatePaymentBankParams, bankEntityID int64) *models.ResponseError {

	store := dbContext.New(cr.dbHandler)
	err := store.UpdatePaymentBank(ctx, *paymentBankParams, int32(bankEntityID))

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

func (cr PaymentBankRepository) DeletePaymentBank(ctx *gin.Context, id int64) *models.ResponseError {

	store := dbContext.New(cr.dbHandler)
	err := store.DeletePaymentBank(ctx, int32(id))

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
