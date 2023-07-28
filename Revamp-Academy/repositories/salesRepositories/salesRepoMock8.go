package salesRepositories

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	dbcontext "codeid.revampacademy/repositories/salesRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type RepoMockup8 struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewMockupApplyRepo8(dbHandler *sql.DB) *RepoMockup8 {
	return &RepoMockup8{
		dbHandler: dbHandler,
	}
}

func (rm RepoMockup8) ListMock8Group(ctx *gin.Context) ([]*dbcontext.CreateMergeMock8, *models.ResponseError) {

	store := dbcontext.New(rm.dbHandler)
	paymentGrup, err := store.ListMock8Group(ctx)

	listMock8Grup := make([]*dbcontext.CreateMergeMock8, 0)

	for _, v := range paymentGrup {
		sales := &dbcontext.CreateMergeMock8{
			Createprogram_entityParam:              v.Createprogram_entityParam,
			CreateUsersParam:                       v.CreateUsersParam,
			CreatePaymentFintechParams:             v.CreatePaymentFintechParams,
			CreatePaymentUsers_accountParams:       v.CreatePaymentUsers_accountParams,
			CreatePaymentTransaction_paymentParams: v.CreatePaymentTransaction_paymentParams,
			CreateSales_order_detailParams:         v.CreateSales_order_detailParams,
		}
		listMock8Grup = append(listMock8Grup, sales)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listMock8Grup, nil
}
func (rm RepoMockup8) GetMock8Group(ctx *gin.Context, poNo string) (*models.MergeMock8, *models.ResponseError) {

	store := dbcontext.New(rm.dbHandler)
	mockup, err := store.GetMock8Group(ctx, poNo)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &mockup, nil
}
