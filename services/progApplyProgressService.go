package services

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories"
	"github.com/gin-gonic/gin"
)

type ProgramApplyProgressService struct {
	programApplyProgressRepository *repositories.ProgramApplyProgressRepository
}

func NewProgramApplyProgressService(programApplyProgressRepository *repositories.ProgramApplyProgressRepository) *ProgramApplyProgressService {
	return &ProgramApplyProgressService{
		programApplyProgressRepository: programApplyProgressRepository,
	}
}

func (paps ProgramApplyProgressService) GetListProgramApplyProgress(ctx *gin.Context) ([]*models.BootcampProgramApplyProgress, *models.ResponseError) {
	return paps.programApplyProgressRepository.GetListProgramApplyProgress(ctx)
}
