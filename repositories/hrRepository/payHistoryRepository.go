package hrRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/hrRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type PayHistoryRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewPayHistoryRepository(dbHandler *sql.DB) *PayHistoryRepository {
	return &PayHistoryRepository{
		dbHandler: dbHandler,
	}
}

func (phr PayHistoryRepository) GetListPayHistory(ctx *gin.Context) ([]*models.HrEmployeePayHistory, *models.ResponseError) {

	store := dbContext.New(phr.dbHandler)
	payHistories, err := store.ListPayHistory(ctx)

	listPayHistory := make([]*models.HrEmployeePayHistory, 0)

	for _, v := range payHistories {
		payHistory := &models.HrEmployeePayHistory{
			EphiEntityID:       v.EphiEntityID,
			EphiRateChangeDate: v.EphiRateChangeDate,
			EphiRateSalary:     v.EphiRateSalary,
			EphiPayFrequence:   v.EphiPayFrequence,
			EphiModifiedDate:   v.EphiModifiedDate,
		}
		listPayHistory = append(listPayHistory, payHistory)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listPayHistory, nil
}

func (phr PayHistoryRepository) GetPayHistory(ctx *gin.Context, id int64) (*models.HrEmployeePayHistory, *models.ResponseError) {

	store := dbContext.New(phr.dbHandler)
	payHistory, err := store.GetPayHistory(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &payHistory, nil
}

func (phr PayHistoryRepository) CreatePayHistory(ctx *gin.Context, payHistoryParams *dbContext.CreatePayHistoryParams) (*models.HrEmployeePayHistory, *models.ResponseError) {

	store := dbContext.New(phr.dbHandler)
	payHistory, err := store.CreatePayHistory(ctx, *payHistoryParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return payHistory, nil
}

func (phr PayHistoryRepository) UpdatePayHistory(ctx *gin.Context, payHistoryParams *dbContext.UpdatePayHistoryParams) *models.ResponseError {

	store := dbContext.New(phr.dbHandler)
	err := store.UpdatePayHistory(ctx, *payHistoryParams)

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

func (phr PayHistoryRepository) DeletePayHistory(ctx *gin.Context, id int64) *models.ResponseError {

	store := dbContext.New(phr.dbHandler)
	err := store.DeletePayHistory(ctx, int32(id))

	if err != nil {
		return &models.ResponseError{
			Message: "error when delete",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been deleted",
		Status:  http.StatusOK,
	}
}
