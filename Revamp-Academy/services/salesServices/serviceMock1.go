package salesServices

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/dbContext/salesContext"
	"codeid.revampacademy/repositories/salesRepositories"
	"github.com/gin-gonic/gin"
)

type ServiceMock struct {
	repositoryMock *salesRepositories.RepositoryMock
}

func NewServiceMock(repositoryMock *salesRepositories.RepositoryMock) *ServiceMock {
	return &ServiceMock{
		repositoryMock: repositoryMock,
	}
}

func (sm ServiceMock) GetMockup(ctx *gin.Context, nama string) (*salesContext.CreateprogramEntityParams, *models.ResponseError) {
	return sm.repositoryMock.GetMockup(ctx, nama)
}
