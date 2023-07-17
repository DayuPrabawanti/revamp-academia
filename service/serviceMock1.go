package service

import (
	"net/http"

	models "codeid.revampacademy/models"
	"codeid.revampacademy/repositories"
	dbcontext "codeid.revampacademy/repositories/dbContext"
	"github.com/gin-gonic/gin"
)

type ServiceMock1 struct {
	repositoryMock1 *repositories.RepositoryMock1
}

func NewServiceMock1(repositoryMock1 *repositories.RepositoryMock1) *ServiceMock1 {
	return &ServiceMock1{
		repositoryMock1: repositoryMock1,
	}
}

func (sm ServiceMock1) GetListMock1(ctx *gin.Context) ([]*models.CurriculumProgramEntity, *models.ResponseError) {
	return sm.repositoryMock1.GetListMock1(ctx)
}

func (sm ServiceMock1) GetMockup1(ctx *gin.Context, nama string) (*models.CurriculumProgramEntity, *models.ResponseError) {
	return sm.repositoryMock1.GetMockup1(ctx, nama)
}

func validateMockup(mockupParams *dbcontext.Createprogram_entityParams) *models.ResponseError {
	if mockupParams.ProgTitle == "" {
		return &models.ResponseError{
			Message: "Invalid Title id",
			Status:  http.StatusBadRequest,
		}
	}
	return nil
}
