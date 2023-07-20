package services

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories"
	"codeid.revampacademy/repositories/dbcontext"
	"github.com/gin-gonic/gin"
)

type SectionService struct {
	sectionRepository *repositories.SectionRepository
}

func NewSectionService(sectionRepository *repositories.SectionRepository) *SectionService {
	return &SectionService{
		sectionRepository: sectionRepository,
	}
}

func (se SectionService) GetListSection(ctx *gin.Context) ([]*models.CurriculumSection, *models.ResponseError) {
	return se.sectionRepository.GetListSection(ctx)
}

func (se SectionService) GetSections(ctx *gin.Context, id int64) (*models.CurriculumSection, *models.ResponseError) {
	return se.sectionRepository.GetSections(ctx, id)
}

func (se SectionService) Createsections(ctx *gin.Context, sectionsParams *dbcontext.CreatesectionsParams) (*models.CurriculumSection, *models.ResponseError) {
	responseErr := validateSections(sectionsParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return se.sectionRepository.Createsections(ctx, sectionsParams)
}

func (se SectionService) UpdateSections(ctx *gin.Context, sectionParams *dbcontext.CreatesectionsParams, id int64) *models.ResponseError {
	responseErr := validateSections(sectionParams)
	if responseErr != nil {
		return responseErr
	}

	return se.sectionRepository.UpdateSections(ctx, sectionParams)
}

func (se SectionService) DeleteSections(ctx *gin.Context, id int64) *models.ResponseError {
	return se.sectionRepository.DeleteSections(ctx, id)
}

func validateSections(sectionsParams *dbcontext.CreatesectionsParams) *models.ResponseError {
	if sectionsParams.SectID == 0 {
		return &models.ResponseError{
			Message: "Invalid program sect id",
			Status:  http.StatusBadRequest,
		}
	}

	if sectionsParams.SectTitle == "" {
		return &models.ResponseError{
			Message: "Invalid program sect name",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}
