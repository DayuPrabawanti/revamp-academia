package services

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories"
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
