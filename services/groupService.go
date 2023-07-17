package services

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories"
	"github.com/gin-gonic/gin"
)

type GroupService struct {
	programEntityRepository *repositories.ProgramEntityRepository
}

func NewGroupService(programEntityRepository *repositories.ProgramEntityRepository) *GroupService {
	return &GroupService{
		programEntityRepository: programEntityRepository,
	}
}

func (gp GroupService) Group(ctx *gin.Context) ([]*models.Group, *models.ResponseError) {
	return gp.programEntityRepository.Group(ctx)
}
