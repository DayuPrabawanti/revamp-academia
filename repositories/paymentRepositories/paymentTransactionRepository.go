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

func (ptr PaymentTransactionRepository) GetListPaymentTransaction(ctx *gin.Context) ([]*dbContext.RecordTransactionUser, *models.ResponseError) {
	store := dbContext.New(ptr.dbHandler)
	paymentTransaction, err := store.ListPaymentTransaction_payment(ctx)
	listPaymentTransactions := make([]*dbContext.RecordTransactionUser, 0)

	for _, v := range paymentTransaction {
		paymentTransaction := &dbContext.RecordTransactionUser{
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

func (ptr PaymentTransactionRepository) GetPaymentTransactionById(ctx *gin.Context, metadata *features.Metadata) ([]*dbContext.RecordTransactionUser, *models.ResponseError) {
	store := dbContext.New(ptr.dbHandler)
	paymentTransaction, err := store.GetPaymentTransaction_payment(ctx, metadata)
	getPaymentTransactions := make([]*dbContext.RecordTransactionUser, 0)

	for _, v := range paymentTransaction {
		paymentTransaction := &dbContext.RecordTransactionUser{
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

func (ptr PaymentTransactionRepository) RecordPaymentTransactionUser(ctx *gin.Context, params *dbContext.RecordTransactionUserParams) (*dbContext.RecordTransactionUser, *models.ResponseError) {
	store := dbContext.New(ptr.dbHandler)
	transactionUser, err := store.RecordPaymentTransactionUser(ctx, *params)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return transactionUser, nil
}

// func (ptr PaymentTransactionRepository) UpdatePaymentTransaction(ctx *gin.Context, paymentTransactionParams *dbContext.CreatePaymentTransaction_paymentParams) *models.ResponseError {
// 	store := dbContext.New(ptr.dbHandler)
// 	err := store.UpdatePaymentTransaction_payment(ctx, *paymentTransactionParams)

// 	if err != nil {
// 		return &models.ResponseError{
// 			Message: "error when update",
// 			Status:  http.StatusInternalServerError,
// 		}
// 	}
// 	return &models.ResponseError{
// 		Message: "data telah terupdate",
// 		Status:  http.StatusOK,
// 	}
// }

func (ptr PaymentTransactionRepository) DeletePaymentTransaction_paymentById(ctx *gin.Context, trpaCodeNumber string) *models.ResponseError {
	store := dbContext.New(ptr.dbHandler)
	err := store.DeletePaymentTransaction_payment(ctx, trpaCodeNumber)

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
