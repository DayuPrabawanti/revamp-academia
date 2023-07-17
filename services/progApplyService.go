package services

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories"
	"github.com/gin-gonic/gin"
)

type ProgramApplyService struct {
	programApplyRepository *repositories.ProgramApplyRepository
}

func NewProgramApplyService(programApplyRepository *repositories.ProgramApplyRepository) *ProgramApplyService {
	return &ProgramApplyService{
		programApplyRepository: programApplyRepository,
	}
}

func (pas ProgramApplyService) GetListProgramApply(ctx *gin.Context) ([]*models.BootcampProgramApply, *models.ResponseError) {
	return pas.programApplyRepository.GetListProgramApply(ctx)
}
