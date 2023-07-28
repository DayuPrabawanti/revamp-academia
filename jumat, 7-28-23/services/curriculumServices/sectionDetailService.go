package curriculumServices

import (
	models "codeid.revampacademy/models"
	repositories "codeid.revampacademy/repositories/curriculumRepositories"
	"github.com/gin-gonic/gin"
)

type SectionDetailService struct {
	sectionDetailRepository *repositories.SectionDetailRepository
}

func NewSectionDetailService(sectionDetailRepository *repositories.SectionDetailRepository) *SectionDetailService {
	return &SectionDetailService{
		sectionDetailRepository: sectionDetailRepository,
	}
}

func (sdm SectionDetailService) GetSectionDetail(ctx *gin.Context, id int64) (*models.CurriculumSectionDetail, *models.ResponseError) {
	return sdm.sectionDetailRepository.GetSectionDetail(ctx, id)
}
