package bootcampService

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/bootcampRepository"
	"codeid.revampacademy/repositories/bootcampRepository/dbContext"
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

func (pe ProgramEntityService) UpdateProgramEntity(ctx *gin.Context, programEntityParams *dbContext.CurriculumProgramEntityParams, id int64) *models.ResponseError {
	responseErr := validateProgramEntity(programEntityParams)
	if responseErr != nil {
		return responseErr
	}

	return pe.programEntityRepository.UpdateProgramEntity(ctx, programEntityParams)
}
func validateProgramEntity(programEntityParams *dbContext.CurriculumProgramEntityParams) *models.ResponseError {
	if programEntityParams.ProgEntityID == 0 {
		return &models.ResponseError{
			Message: "Invalid program entity id",
			Status:  http.StatusBadRequest,
		}
	}

	if programEntityParams.ProgTitle == "" {
		return &models.ResponseError{
			Message: "Invalid program entity name",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}

func (pe ProgramEntityService) GetProgEntity(ctx *gin.Context, id int64) (*models.CurriculumProgramEntity, *models.ResponseError) {
	return pe.programEntityRepository.GetProgEntity(ctx, id)
}
