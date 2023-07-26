package salesServices

import (
	"codeid.revampacademy/models"
	sapo "codeid.revampacademy/repositories/salesRepositories"
	dbcontext "codeid.revampacademy/repositories/salesRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type ServiceMock2 struct {
	repoMockup2 *sapo.RepoMockup2
}

func NewMockupApplyService2(repoMockup2 *sapo.RepoMockup2) *ServiceMock2 {
	return &ServiceMock2{
		repoMockup2: repoMockup2,
	}
}

func (sm ServiceMock2) ListBootcampGroup(ctx *gin.Context) ([]*dbcontext.CreateMergeMock2, *models.ResponseError) {
	return sm.repoMockup2.ListBootcampGroup(ctx)
}
