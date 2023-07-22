package salesServices

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/salesRepositories"
	"codeid.revampacademy/repositories/salesRepositories/dbContext"
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

func (sm ServiceMock) GetMockup(ctx *gin.Context, nama string) (*dbContext.CreateprogramEntityParams, *models.ResponseError) {
	return sm.repositoryMock.GetMockup(ctx, nama)
}

func (sm ServiceMock) GetListProgram(ctx *gin.Context, nama string) ([]*dbContext.CreateprogramEntityParams, *models.ResponseError) {
	return sm.repositoryMock.GetListProgram(ctx, nama)
}

func validateMockup(mockupParams *models.CurriculumProgramEntity) *models.ResponseError {
	if mockupParams.ProgEntityID == 0 {
		return &models.ResponseError{
			Message: "Invalid category id",
			Status:  http.StatusBadRequest,
		}
	}

	if mockupParams.ProgTitle == "" {
		return &models.ResponseError{
			Message: "Invalid category name",
			Status:  http.StatusBadRequest,
		}
	}
	return nil
}
