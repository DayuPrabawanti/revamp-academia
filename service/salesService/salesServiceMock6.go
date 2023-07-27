package salesservice

import (
	"codeid.revampacademy/models"
	sapo "codeid.revampacademy/repositories/salesRepositories"
	"github.com/gin-gonic/gin"
)

type ServiceMock6 struct {
	repoMockup6 *sapo.RepoMockup6
}

func NewServiceShoppingCart1(repoMockup6 *sapo.RepoMockup6) *ServiceMock6 {
	return &ServiceMock6{
		repoMockup6: repoMockup6,
	}
}

func (sm ServiceMock6) GetUsersIdShoopingCart1Service(ctx *gin.Context, id int64) (*models.MergeShopMock7, *models.ResponseError) {
	return sm.repoMockup6.GetUsersIdShoopingCart1Repo(ctx, id)
}

func (sm ServiceMock6) ListSalesOrderService(ctx *gin.Context) ([]*models.SalesSalesOrderDetail, *models.ResponseError) {
	return sm.repoMockup6.ListSalesOrder(ctx)
}
func (sm ServiceMock6) GetAccountNumbersService(ctx *gin.Context, account string) (*models.MergePayment, *models.ResponseError) {
	return sm.repoMockup6.GetAccountNumbersRepo(ctx, account)
}
