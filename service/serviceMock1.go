package service

import (
	"net/http"

	models "codeid.revampacademy/models"
	"codeid.revampacademy/repositories"
	dbcontext "codeid.revampacademy/repositories/dbContext"
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

func (sm ServiceMock) GetMockup(ctx *gin.Context, nama string) (*dbcontext.CreateprogramEntityParams, *models.ResponseError) {
	return sm.repositoryMock.GetMockup(ctx, nama)
}

func (sm ServiceMock) GetMockupId(ctx *gin.Context, id int64) (*models.CurriculumProgramEntity, *models.ResponseError) {
	return sm.repositoryMock.GetMockupId(ctx, id)
}

func validateMockup(mockupParams *dbcontext.CreateprogramEntityParams) *models.ResponseError {
	if mockupParams.ProgTitle == "" {
		return &models.ResponseError{
			Message: "Invalid Title id",
			Status:  http.StatusBadRequest,
		}
	}
	return nil
}
