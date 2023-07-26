package paymentRepositories

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/paymentRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type PaymentTransactionRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewPaymentTransactionRepository(dbHandler *sql.DB) *PaymentTransactionRepository {
	return &PaymentTransactionRepository{
		dbHandler: dbHandler,
	}
}

func (ptr PaymentTransactionRepository) GetListPaymentTransaction(ctx *gin.Context) ([]*models.PaymentTransactionPayment, *models.ResponseError) {
	store := dbContext.New(ptr.dbHandler)
	paymentTransaction, err := store.ListPaymentTransaction_payment(ctx)
	listPaymentTransactions := make([]*models.PaymentTransactionPayment, 0)

	for _, v := range paymentTransaction {
		paymentTransaction := &models.PaymentTransactionPayment{
			TrpaID:           v.TrpaID,
			TrpaCodeNumber:   v.TrpaCodeNumber,
			TrpaOrderNumber:  v.TrpaOrderNumber,
			TrpaDebit:        v.TrpaDebit,
			TrpaCredit:       v.TrpaCredit,
			TrpaType:         v.TrpaType,
			TrpaNote:         v.TrpaNote,
			TrpaModifiedDate: v.TrpaModifiedDate,
			TrpaSourceID:     v.TrpaSourceID,
			TrpaTargetID:     v.TrpaTargetID,
			TrpaUserEntityID: v.TrpaUserEntityID,
		}
		listPaymentTransactions = append(listPaymentTransactions, paymentTransaction)
	}
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return listPaymentTransactions, nil
}

func (ptr PaymentTransactionRepository) GetPaymentTransactionById(ctx *gin.Context, id int64) (*models.PaymentTransactionPayment, *models.ResponseError) {
	store := dbContext.New(ptr.dbHandler)
	paymentTransaction, err := store.GetPaymentTransaction_payment(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &paymentTransaction, nil
}

func (ptr PaymentTransactionRepository) CreatePaymentTransaction(ctx *gin.Context, paymentTransactionParams *dbContext.CreatePaymentTransaction_paymentParams) (*models.PaymentTransactionPayment, *models.ResponseError) {
	store := dbContext.New(ptr.dbHandler)
	paymentTransaction, err := store.CreatePaymentTransaction_payment(ctx, *paymentTransactionParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return paymentTransaction, nil
}

func (ptr PaymentTransactionRepository) UpdatePaymentTransaction(ctx *gin.Context, paymentTransactionParams *dbContext.CreatePaymentTransaction_paymentParams) *models.ResponseError {
	store := dbContext.New(ptr.dbHandler)
	err := store.UpdatePaymentTransaction_payment(ctx, *paymentTransactionParams)

	if err != nil {
		return &models.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data telah terupdate",
		Status:  http.StatusOK,
	}
}

func (ptr PaymentTransactionRepository) DeletePaymentTransaction(ctx *gin.Context, id int64) *models.ResponseError {
	store := dbContext.New(ptr.dbHandler)
	err := store.DeletePaymentTransaction_payment(ctx, int32(id))

	if err != nil {
		return &models.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data telah terhapus",
		Status:  http.StatusOK,
	}
}
