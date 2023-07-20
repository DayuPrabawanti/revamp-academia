package services

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories"
	"codeid.revampacademy/repositories/dbcontext"
	"github.com/gin-gonic/gin"
)

type ProgramEntityService struct {
	programEntityRepository *repositories.ProgramEntityRepository
}

func NewProgramEntityService(programEntityRepository *repositories.ProgramEntityRepository) *ProgramEntityService {
	return &ProgramEntityService{
		programEntityRepository: programEntityRepository,
	}
}

func (pe ProgramEntityService) GetListProgramEntity(ctx *gin.Context) ([]*models.CurriculumProgramEntity, *models.ResponseError) {
	return pe.programEntityRepository.GetListProgramEntity(ctx)
}

func (pe ProgramEntityService) Group(ctx *gin.Context) ([]*models.Group, *models.ResponseError) {
	return pe.programEntityRepository.Group(ctx)
}

func (pe ProgramEntityService) GetListMasterCategory(ctx *gin.Context) ([]*models.MasterCategory, *models.ResponseError) {
	return pe.programEntityRepository.GetListMasterCategory(ctx)
}

func (ps ProgramEntityService) GetProgramEntity(ctx *gin.Context, id int64) (*models.CurriculumProgramEntity, *models.ResponseError) {
	return ps.programEntityRepository.GetProgramEntity(ctx, id)
}

func (pe ProgramEntityService) CreateProgramEntity(ctx *gin.Context, programEntityParams *dbcontext.CreateProgramEntityParams) (*models.CurriculumProgramEntity, *models.ResponseError) {
	responseErr := validateProgramEntity(programEntityParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return pe.programEntityRepository.CreateProgramEntity(ctx, programEntityParams)
}

func (pe ProgramEntityService) CreateGroup(ctx *gin.Context, groupParams *dbcontext.CreateGroup) (*models.Group, *models.ResponseError) {

	return pe.programEntityRepository.CreateGroup(ctx, groupParams)
}

func (pe ProgramEntityService) UpdateProgramEntity(ctx *gin.Context, programEntityParams *dbcontext.CreateProgramEntityParams, id int64) *models.ResponseError {
	responseErr := validateProgramEntity(programEntityParams)
	if responseErr != nil {
		return responseErr
	}

	return pe.programEntityRepository.UpdateProgramEntity(ctx, programEntityParams)
}

func (pe ProgramEntityService) DeleteProgramEntity(ctx *gin.Context, id int64) *models.ResponseError {
	return pe.programEntityRepository.DeleteProgramEntity(ctx, id)
}

func validateProgramEntity(programEntityParams *dbcontext.CreateProgramEntityParams) *models.ResponseError {
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
