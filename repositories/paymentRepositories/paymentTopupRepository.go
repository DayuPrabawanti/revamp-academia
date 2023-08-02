package repositories

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/paymentRepositories/dbContext"
	"golang.org/x/net/context"
)

type PaymentTopupRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
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

// func (ptr PaymentTopupRepository) GetListTopupDetail(ctx *gin.Context) ([]*dbContext.TopupDetail, *models.ResponseError) {

// 	store := dbContext.New(ptr.dbHandler)
// 	paymentTopups, err := store.ListTopupDetail(ctx)

// 	listPaymentTopups := make([]*dbContext.TopupDetail, 0)

// 	for _, v := range paymentTopups {
// 		paymentTopup := &dbContext.TopupDetail{
// 			SourceName:    v.SourceName,
// 			SourceAccount: v.SourceAccount,
// 			SourceSaldo:   v.SourceSaldo,
// 			TargetName:    v.TargetName,
// 			TargetAccount: v.TargetAccount,
// 			TargetSaldo:   v.TargetSaldo,
// 		}
// 		listPaymentTopups = append(listPaymentTopups, paymentTopup)
// 	}

// 	if err != nil {
// 		return nil, &models.ResponseError{
// 			Message: err.Error(),
// 			Status:  http.StatusInternalServerError,
// 		}
// 	}

// 	return listPaymentTopups, nil
// }

// func (ptr PaymentTopupRepository) GetTopupDetailById(ctx *gin.Context, id int32) ([]*dbContext.TopupDetail, *models.ResponseError) {

// 	store := dbContext.New(ptr.dbHandler)
// 	paymentTopups, err := store.GetTopupDetailById(ctx, id)

// 	listPaymentTopups := make([]*dbContext.TopupDetail, 0)

// 	for _, v := range paymentTopups {
// 		paymentTopup := &dbContext.TopupDetail{
// 			SourceName:    v.SourceName,
// 			SourceAccount: v.SourceAccount,
// 			SourceSaldo:   v.SourceSaldo,
// 			TargetName:    v.TargetName,
// 			TargetAccount: v.TargetAccount,
// 			TargetSaldo:   v.TargetSaldo,
// 		}
// 		listPaymentTopups = append(listPaymentTopups, paymentTopup)
// 	}

// 	if err != nil {
// 		return nil, &models.ResponseError{
// 			Message: err.Error(),
// 			Status:  http.StatusInternalServerError,
// 		}
// 	}

// 	return listPaymentTopups, nil
// }

// func (ptr *PaymentTopupRepository) PerformTransfer(ctx *gin.Context, fromAccount string, amount float64, toAccount string, fromUserID int, toUserID int) *models.ResponseError {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Printf("Recovered from panic: %v", r)
// 		}
// 	}()

// 	err := ptr.paymentService.TransferFunds(ctx, fromAccount, amount, toAccount, fromUserID, toUserID)
// 	if err != nil {
// 		return &models.ResponseError{
// 			Message: err.Error(),
// 			Status:  http.StatusInternalServerError,
// 		}
// 	}
// 	return nil
// }
