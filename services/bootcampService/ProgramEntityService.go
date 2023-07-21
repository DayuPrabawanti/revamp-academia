package bootcampService

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/bootcampRepository"
	"github.com/gin-gonic/gin"
)

type ProgramEntityService struct {
	programEntityRepository *bootcampRepository.ProgramentityRepository
}

func NewProgramEntityService(programEntityRepository *bootcampRepository.ProgramentityRepository) *ProgramEntityService {
	return &ProgramEntityService{
		programEntityRepository: programEntityRepository,
	}
}

func (pe ProgramEntityService) GetListProgramEntity(ctx *gin.Context) ([]*models.CurriculumProgramEntity, *models.ResponseError) {
	return pe.programEntityRepository.GetListProgramEntity(ctx)
}
