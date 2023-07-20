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

func (sm ServiceMock) GetListProgram(ctx *gin.Context) ([]*dbcontext.CreateprogramEntityParams, *models.ResponseError) {
	return sm.repositoryMock.GetListProgram(ctx)
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
