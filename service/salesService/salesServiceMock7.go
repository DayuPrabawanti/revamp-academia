package salesservice

import (
	"codeid.revampacademy/models"
	sapo "codeid.revampacademy/repositories/salesRepositories"
	dbcontext "codeid.revampacademy/repositories/salesRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type ServiceMock7 struct {
	repoMockup7 *sapo.RepoMock7
}

func NewServiceShoppingCart2(repoMockup7 *sapo.RepoMock7) *ServiceMock7 {
	return &ServiceMock7{
		repoMockup7: repoMockup7,
	}
}

func (sm ServiceMock7) CreateOrderDetailService(ctx *gin.Context, orderParams *dbcontext.CreateSales_order_detailParams) (*models.SalesSalesOrderDetail, *models.ResponseError) {
	return sm.repoMockup7.CreateOrderDetail(ctx, orderParams)
}

func (sm ServiceMock7) CancelOrderDetailService(ctx *gin.Context, id int64) *models.ResponseError {
	return sm.repoMockup7.CancelOrderDetail(ctx, id)
}

func (sm ServiceMock7) GetAccountNumberMock7Service(ctx *gin.Context, account string) (*dbcontext.PaymentParams, *models.ResponseError) {
	return sm.repoMockup7.GetAccountNumbersMock7Repo(ctx, account)
}

func (sm ServiceMock7) GetUsersIdShoopingCart2Service(ctx *gin.Context, id int64) (*dbcontext.UserEntityParamsMork7, *models.ResponseError) {
	return sm.repoMockup7.GetUsersIdShoopingCart2Repo(ctx, id)
}
