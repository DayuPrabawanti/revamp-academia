package repositories

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/paymentRepositories/dbContext"
	"github.com/gin-gonic/gin"
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

func (ptr PaymentTopupRepository) GetListTopupDetail(ctx *gin.Context) ([]*dbContext.TopupDetail, *models.ResponseError) {

	store := dbContext.New(ptr.dbHandler)
	paymentTopups, err := store.ListTopupDetail(ctx)

	listPaymentTopups := make([]*dbContext.TopupDetail, 0)

	for _, v := range paymentTopups {
		paymentTopup := &dbContext.TopupDetail{
			SourceName:    v.SourceName,
			SourceAccount: v.SourceAccount,
			SourceSaldo:   v.SourceSaldo,
			TargetName:    v.TargetName,
			TargetAccount: v.TargetAccount,
			TargetSaldo:   v.TargetSaldo,
		}
		listPaymentTopups = append(listPaymentTopups, paymentTopup)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listPaymentTopups, nil
}
