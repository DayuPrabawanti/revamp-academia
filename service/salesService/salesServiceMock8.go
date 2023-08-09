package salesservice

import (
	"codeid.revampacademy/models"
	sapo "codeid.revampacademy/repositories/salesRepositories"
	dbcontext "codeid.revampacademy/repositories/salesRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type ServiceMockup8 struct {
	repoMockup8 *sapo.RepoMockup8
}

func NewServiceShoppingCart3(repoMockup8 *sapo.RepoMockup8) *ServiceMockup8 {
	return &ServiceMockup8{
		repoMockup8: repoMockup8,
	}
}

func (sm ServiceMockup8) GetIdSummaryOrderMock8Service(ctx *gin.Context, poNo string) (*dbcontext.GetSummaryOrderMock8, *models.ResponseError) {
	return sm.repoMockup8.GetIdSummaryOrderMock8Repo(ctx, poNo)
}
