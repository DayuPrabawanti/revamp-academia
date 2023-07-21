package bootcampService

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/bootcampRepository"
	"github.com/gin-gonic/gin"
)

type ProgramApplyProgressService struct {
	programApplyProgressRepository *bootcampRepository.ProgramApplyProgressRepository
}

func NewProgramApplyProgressService(programApplyProgressRepository *bootcampRepository.ProgramApplyProgressRepository) *ProgramApplyProgressService {
	return &ProgramApplyProgressService{
		programApplyProgressRepository: programApplyProgressRepository,
	}
}

func (paps ProgramApplyProgressService) GetListProgramApplyProgress(ctx *gin.Context) ([]*models.BootcampProgramApplyProgress, *models.ResponseError) {
	return paps.programApplyProgressRepository.GetListProgramApplyProgress(ctx)
}
