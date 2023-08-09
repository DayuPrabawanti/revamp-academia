package salesRepositories

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	dbcontext "codeid.revampacademy/repositories/salesRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type RepoMock7 struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewRepoShoppingCart2(dbHandler *sql.DB) *RepoMock7 {
	return &RepoMock7{
		dbHandler: dbHandler,
	}
}

func (rm RepoMock7) GetUsersIdShoopingCart2Repo(ctx *gin.Context, id int64) (*dbcontext.UserEntityParamsMork7, *models.ResponseError) {

	store := dbcontext.New(rm.dbHandler)
	mockup, err := store.GetUsersIdMock7(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &mockup, nil
}

func (rm RepoMock7) CreateOrderDetail(ctx *gin.Context, orderParams *dbcontext.CreateSales_order_detailParams) (*models.SalesSalesOrderDetail, *models.ResponseError) {
	store := dbcontext.New(rm.dbHandler)
	order, err := store.CreateSales_order_detail(ctx, *orderParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return order, nil
}
func (rm RepoMock7) CancelOrderDetail(ctx *gin.Context, id int64) *models.ResponseError {

	store := dbcontext.New(rm.dbHandler)
	err := store.DeleteSales_order_detail(ctx, int16(id))

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

func (rm RepoMock7) GetAccountNumbersMock7Repo(ctx *gin.Context, account string) (*dbcontext.PaymentParams, *models.ResponseError) {

	store := dbcontext.New(rm.dbHandler)
	mockup, err := store.GetAccountNumberMock7(ctx, string(account))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &mockup, nil
}
