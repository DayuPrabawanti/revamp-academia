package paymentRepositories

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

func (ptr PaymentTopupRepository) GetTopupDetail(ctx *gin.Context, sourceBankEntityID int32, targetFintechEntityID int32) (*dbContext.TopupDetail, *models.ResponseError) {

	store := dbContext.New(ptr.dbHandler)
	paymentTopup, err := store.GetTopupDetail(ctx, sourceBankEntityID, targetFintechEntityID)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return paymentTopup, nil
}
