package bootcampService

import (
	"codeid.revampacademy/models"
	"github.com/gin-gonic/gin"
)

type ProgramEntityService struct {
	programEntityRepository *repositories.ProgramentityRepository
}

func NewProgramEntityService(programEntityRepository *repositories.ProgramentityRepository) *ProgramEntityService {
	return &ProgramEntityService{
		programEntityRepository: programEntityRepository,
	}
}

func (pe ProgramEntityService) GetListProgramEntity(ctx *gin.Context) ([]*models.CurriculumProgramEntity, *models.ResponseError) {
	return pe.programEntityRepository.GetListProgramEntity(ctx)
}
