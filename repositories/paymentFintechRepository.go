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

func (pfr PaymentFintechRepository) GetListPaymentFintech(ctx *gin.Context) ([]*models.PaymentFintech, *models.ResponseError) {

	store := dbContext.New(pfr.dbHandler)
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

func (pfr PaymentFintechRepository) GetPaymentFintechByName(ctx *gin.Context, name string) (*models.PaymentFintech, *models.ResponseError) {

	store := dbContext.New(pfr.dbHandler)
	paymentFintech, err := store.GetPaymentFintech(ctx, string(name))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &paymentFintech, nil
}

func (pfr PaymentFintechRepository) CreateNewPaymentFintech(ctx *gin.Context, paymentFintechParams *dbContext.CreatePaymentFintechParams) (*models.PaymentFintech, *models.ResponseError) {

	store := dbContext.New(pfr.dbHandler)
	paymentFintech, err := store.CreatePaymentFintech(ctx, *paymentFintechParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return paymentFintech, nil
}

func (pfr PaymentFintechRepository) UpdatePaymentFintechById(ctx *gin.Context, paymentFintechParams *dbContext.CreatePaymentFintechParams) *models.ResponseError {

	store := dbContext.New(pfr.dbHandler)
	err := store.UpdatePaymentFintech(ctx, *paymentFintechParams)

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

func (pfr PaymentFintechRepository) DeletePaymentFintechById(ctx *gin.Context, id int64) *models.ResponseError {

	store := dbContext.New(pfr.dbHandler)
	err := store.DeletePaymentFintech(ctx, int32(id))

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
