package bootcampService

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/bootcampRepository"
	"github.com/gin-gonic/gin"
)

type ProgramApplyService struct {
	programApplyRepository *bootcampRepository.ProgramApplyRepository
}

func NewProgramApplyService(programApplyRepository *bootcampRepository.ProgramApplyRepository) *ProgramApplyService {
	return &ProgramApplyService{
		programApplyRepository: programApplyRepository,
	}
}

func (pas ProgramApplyService) GetListProgramApply(ctx *gin.Context) ([]*models.BootcampProgramApply, *models.ResponseError) {
	return pas.programApplyRepository.GetListProgramApply(ctx)
}
