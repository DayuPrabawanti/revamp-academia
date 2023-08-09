package salesRepositories

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	dbcontext "codeid.revampacademy/repositories/salesRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type RepoMockup6 struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewRepoShoppingCart1(dbHandler *sql.DB) *RepoMockup6 {
	return &RepoMockup6{
		dbHandler: dbHandler,
	}
}

func (rm RepoMockup6) GetUsersIdShoopingCart1Repo(ctx *gin.Context, id int64) (*models.MergeShopMock7, *models.ResponseError) {

	store := dbcontext.New(rm.dbHandler)
	mockup, err := store.GetUsersIdMock6(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &mockup, nil
}

func (rm RepoMockup6) ListSalesOrder(ctx *gin.Context) ([]*models.SalesSalesOrderDetail, *models.ResponseError) {

	store := dbcontext.New(rm.dbHandler)
	orderDetail, err := store.ListSalesOrderMock6(ctx)

	listOrder := make([]*models.SalesSalesOrderDetail, 0)

	for _, v := range orderDetail {
		sales := &models.SalesSalesOrderDetail{
			SodeID:           v.SodeID,
			SodeQty:          v.SodeQty,
			SodeUnitPrice:    v.SodeUnitPrice,
			SodeUnitDiscount: v.SodeUnitDiscount,
			SodeLineTotal:    v.SodeLineTotal,
			SodeModifiedDate: v.SodeModifiedDate,
			SodeSoheID:       v.SodeSoheID,
			SodeProgEntityID: v.SodeProgEntityID,
		}
		listOrder = append(listOrder, sales)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listOrder, nil
}

func (rm RepoMockup6) GetAccountNumbersRepo(ctx *gin.Context, account string) (*models.MergePayment, *models.ResponseError) {

	store := dbcontext.New(rm.dbHandler)
	mockup, err := store.GetAccountNumbers(ctx, string(account))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &mockup, nil
}
