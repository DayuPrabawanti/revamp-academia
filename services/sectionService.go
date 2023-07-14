package services

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories"
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
