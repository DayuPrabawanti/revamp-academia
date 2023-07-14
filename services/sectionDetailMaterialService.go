package services

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories"
	"github.com/gin-gonic/gin"
)

// SECTION DETAIL MATERIAL

type SectionDetailMaterialService struct {
	sectionDetailMaterialRepository *repositories.SectionDetailMaterialRepository
}

func NewSectionDetailMaterialService(sectionDetailMaterialRepository *repositories.SectionDetailMaterialRepository) *SectionDetailMaterialService {
	return &SectionDetailMaterialService{
		sectionDetailMaterialRepository: sectionDetailMaterialRepository,
	}
}

func (sd SectionDetailMaterialService) GetListSectionDetailMaterial(ctx *gin.Context) ([]*models.CurriculumSectionDetailMaterial, *models.ResponseError) {
	return sd.sectionDetailMaterialRepository.GetListSectionDetailMaterial(ctx)
}
