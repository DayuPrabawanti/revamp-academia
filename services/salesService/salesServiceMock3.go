package salesService

import (
	"codeid.revampacademy/models"
	sapo "codeid.revampacademy/repositories/salesRepositories"
	dbcontext "codeid.revampacademy/repositories/salesRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type ServiceMock3 struct {
	repoMockup3 *sapo.RepoMockup3
}

func NewMockupApplyService(repoMockup3 *sapo.RepoMockup3) *ServiceMock3 {
	return &ServiceMock3{
		repoMockup3: repoMockup3,
	}
}

func (sm ServiceMock3) CreateUsers(ctx *gin.Context, userParam *dbcontext.CreateUsersParams) (*dbcontext.CreateUsersParams, *models.ResponseError) {
	return sm.repoMockup3.CreateUser(ctx, userParam)
}

func (sm ServiceMock3) CreateEducations(ctx *gin.Context, educationParam *dbcontext.CreateEducationParams) (*dbcontext.CreateEducationParams, *models.ResponseError) {
	return sm.repoMockup3.CreateEducation(ctx, educationParam)
}

func (sm ServiceMock3) CreateMedian(ctx *gin.Context, mediaParam *dbcontext.CreateMediaParams) (*dbcontext.CreateMediaParams, *models.ResponseError) {
	return sm.repoMockup3.CreateMedia(ctx, mediaParam)
}

func (sm ServiceMock3) CreateMergeMocks(ctx *gin.Context, mergeParam *dbcontext.CreateMergeMock) (*dbcontext.CreateMergeMock, *models.ResponseError) {
	return sm.repoMockup3.CreateMergeMock(ctx, mergeParam)
}

func (sm ServiceMock3) GetUsers(ctx *gin.Context, id int64) (*models.UsersUser, *models.ResponseError) {
	return sm.repoMockup3.GetUsers(ctx, id)
}
