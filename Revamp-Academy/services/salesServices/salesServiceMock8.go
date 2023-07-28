package salesServices

import (
	"codeid.revampacademy/models"
	sapo "codeid.revampacademy/repositories/salesRepositories"
	dbcontext "codeid.revampacademy/repositories/salesRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type ServiceMock8 struct {
	repoMockup8 *sapo.RepoMockup8
}

func NewMockupApplyService8(repoMock8 *sapo.RepoMockup8) *ServiceMock8 {
	return &ServiceMock8{
		repoMockup8: repoMock8,
	}
}

func (sm ServiceMock8) ListMock8Group(ctx *gin.Context) ([]*dbcontext.CreateMergeMock8, *models.ResponseError) {
	return sm.repoMockup8.ListMock8Group(ctx)
}

func (sm ServiceMock8) GetMock8Group(ctx *gin.Context, poNo string) (*models.MergeMock8, *models.ResponseError) {
	return sm.repoMockup8.GetMock8Group(ctx, poNo)
}
