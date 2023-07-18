package services

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories"
	"codeid.revampacademy/repositories/dbContext"
	"github.com/gin-gonic/gin"
)

type ServiceMock struct {
	repositoryMock *repositories.RepositoryMock
}

func NewServiceMock(repositoryMock *repositories.RepositoryMock) *ServiceMock {
	return &ServiceMock{
		repositoryMock: repositoryMock,
	}
}

func (sm ServiceMock) GetMockup(ctx *gin.Context, nama string) (*dbContext.CreateprogramEntityParams, *models.ResponseError) {
	return sm.repositoryMock.GetMockup(ctx, nama)
}
