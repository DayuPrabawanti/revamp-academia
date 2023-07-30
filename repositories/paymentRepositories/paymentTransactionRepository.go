package repositories

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/models/features"
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

func (ptr PaymentTransactionRepository) GetListPaymentTransaction(ctx *gin.Context) ([]*dbContext.TransactionUser, *models.ResponseError) {
	store := dbContext.New(ptr.dbHandler)
	paymentTransaction, err := store.ListPaymentTransaction_payment(ctx)
	listPaymentTransactions := make([]*dbContext.TransactionUser, 0)

	for _, v := range paymentTransaction {
		paymentTransaction := &dbContext.TransactionUser{
			TrpaCodeNumber:   v.TrpaCodeNumber,
			TrpaModifiedDate: v.TrpaModifiedDate,
			TrpaDebit:        v.TrpaDebit,
			TrpaCredit:       v.TrpaCredit,
			TrpaNote:         v.TrpaNote,
			TrpaOrderNumber:  v.TrpaOrderNumber,
			TrpaFromID:       v.TrpaFromID,
			TrpaToID:         v.TrpaToID,
			TrpaType:         v.TrpaType,
			UserName:         v.UserName,
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

func (ptr PaymentTransactionRepository) GetPaymentTransactionById(ctx *gin.Context, metadata *features.Metadata) ([]*dbContext.TransactionUser, *models.ResponseError) {
	store := dbContext.New(ptr.dbHandler)
	paymentTransaction, err := store.GetPaymentTransaction_payment(ctx, metadata)
	getPaymentTransactions := make([]*dbContext.TransactionUser, 0)

	for _, v := range paymentTransaction {
		paymentTransaction := &dbContext.TransactionUser{
			TrpaCodeNumber:   v.TrpaCodeNumber,
			TrpaModifiedDate: v.TrpaModifiedDate,
			TrpaDebit:        v.TrpaDebit,
			TrpaCredit:       v.TrpaCredit,
			TrpaNote:         v.TrpaNote,
			TrpaOrderNumber:  v.TrpaOrderNumber,
			TrpaFromID:       v.TrpaFromID,
			TrpaToID:         v.TrpaToID,
			TrpaType:         v.TrpaType,
			UserName:         v.UserName,
		}
		getPaymentTransactions = append(getPaymentTransactions, paymentTransaction)
	}
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return getPaymentTransactions, nil
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
