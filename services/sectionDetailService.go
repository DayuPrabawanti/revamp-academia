package services

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories"
	"codeid.revampacademy/repositories/dbcontext"
	"github.com/gin-gonic/gin"
)

// SECTION DETAIL

type SectionDetailService struct {
	sectionDetailRepository *repositories.SectionDetailRepository
}

func NewSectionDetailService(sectionDetailRepository *repositories.SectionDetailRepository) *SectionDetailService {
	return &SectionDetailService{
		sectionDetailRepository: sectionDetailRepository,
	}
}

func (sd SectionDetailService) GetListSectionDetail(ctx *gin.Context) ([]*models.CurriculumSectionDetail, *models.ResponseError) {
	return sd.sectionDetailRepository.GetListSectionDetail(ctx)
}

func (sd SectionDetailService) GetSectionDetail(ctx *gin.Context, id int64) (*models.CurriculumSectionDetail, *models.ResponseError) {
	return sd.sectionDetailRepository.GetSectionDetail(ctx, id)
}

func (sd SectionDetailService) CreateSectionDetail(ctx *gin.Context, sectionDetailParams *dbcontext.CreatesectionDetailParams) (*models.CurriculumSectionDetail, *models.ResponseError) {
	responseErr := validateSectionDetail(sectionDetailParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return sd.sectionDetailRepository.CreateSectionDetail(ctx, sectionDetailParams)
}

func (sd SectionDetailService) UpdateSectionDetail(ctx *gin.Context, sectionDetailParams *dbcontext.CreatesectionDetailParams, id int64) *models.ResponseError {
	responseErr := validateSectionDetail(sectionDetailParams)
	if responseErr != nil {
		return responseErr
	}

	return sd.sectionDetailRepository.UpdateSectionDetail(ctx, sectionDetailParams)
}

func (sd SectionDetailService) DeleteSectionDetail(ctx *gin.Context, id int64) *models.ResponseError {
	return sd.sectionDetailRepository.DeleteSectionDetail(ctx, id)
}

func validateSectionDetail(sectionDetailParams *dbcontext.CreatesectionDetailParams) *models.ResponseError {
	if sectionDetailParams.SecdID == 0 {
		return &models.ResponseError{
			Message: "Invalid program secd id",
			Status:  http.StatusBadRequest,
		}
	}

	if sectionDetailParams.SecdTitle.String == "" {
		return &models.ResponseError{
			Message: "Invalid program secd name",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}
