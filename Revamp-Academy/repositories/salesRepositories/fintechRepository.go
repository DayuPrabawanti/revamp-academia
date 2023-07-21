package salesRepositories

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/dbContext/salesContext"
	"github.com/gin-gonic/gin"
)

type FintechRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewFintechRepository(dbHandler *sql.DB) *FintechRepository {
	return &FintechRepository{
		dbHandler: dbHandler,
	}
}

func (cr FintechRepository) GetPaymentFintech(ctx *gin.Context, id int32) (*models.PaymentFintech, *models.ResponseError) {

	store := salesContext.New(cr.dbHandler)
	specialoffer, err := store.GetPaymentFintech(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &specialoffer, nil
}
