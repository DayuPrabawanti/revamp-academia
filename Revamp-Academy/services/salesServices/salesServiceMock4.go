package salesServices

import (
	"codeid.revampacademy/models"
	sapo "codeid.revampacademy/repositories/salesRepositories"
	dbcontext "codeid.revampacademy/repositories/salesRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type ServiceMock4 struct {
	repoMockup4 *sapo.RepoMockup4
}

func NewMockupApplyService4(repoMock4 *sapo.RepoMockup4) *ServiceMock4 {
	return &ServiceMock4{
		repoMockup4: repoMock4,
	}
}

func (sm ServiceMock4) ListMock4Group(ctx *gin.Context) ([]*dbcontext.CreateMergeMock4, *models.ResponseError) {
	return sm.repoMockup4.ListMock4Group(ctx)
}

func (sm ServiceMock4) GetMock4Group(ctx *gin.Context, id int32) (*models.MergeMock4, *models.ResponseError) {
	return sm.repoMockup4.GetMock4Group(ctx, id)
}
