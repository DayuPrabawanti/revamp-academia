package repositories

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/dbContext"
	"github.com/gin-gonic/gin"
)

type PaymentFintechRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewPaymentFintechRepository(dbHandler *sql.DB) *PaymentFintechRepository {
	return &PaymentFintechRepository{
		dbHandler: dbHandler,
	}
}

func (pbr PaymentFintechRepository) GetListPaymentFintech(ctx *gin.Context) ([]*models.PaymentFintech, *models.ResponseError) {

	store := dbContext.New(pbr.dbHandler)
	paymentFintechs, err := store.ListPaymentFintech(ctx)

	listPaymentFintechs := make([]*models.PaymentFintech, 0)

	for _, v := range paymentFintechs {
		paymentFintech := &models.PaymentFintech{
			FintEntityID:     v.FintEntityID,
			FintCode:         v.FintCode,
			FintName:         v.FintName,
			FintModifiedDate: v.FintModifiedDate,
		}
		listPaymentFintechs = append(listPaymentFintechs, paymentFintech)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listPaymentFintechs, nil
}

// func (pbr PaymentFintechRepository) GetPaymentFintechByName(ctx *gin.Context, name string) (*models.PaymentFintech, *models.ResponseError) {
// 	store := dbContext.New(pbr.dbHandler)
// 	paymentFintech, err := store.GetPaymentFintech(ctx, name)

// 	if err != nil {
// 		return nil, &models.ResponseError{
// 			Message: err.Error(),
// 			Status:  http.StatusInternalServerError,
// 		}
// 	}
// 	return &paymentFintech, nil
// }
