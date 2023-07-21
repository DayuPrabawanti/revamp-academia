package hrService

import (
	"codeid.revampacademy/models"
	hr "codeid.revampacademy/repositories/hrRepository"
	"github.com/gin-gonic/gin"
)

type TalentsDetailMockupService struct {
	talentDetailRepository *hr.TalentsDetailMockupRepository
}

func NewTalentDetailMockupService(talentDetailRepository *hr.TalentsDetailMockupRepository) *TalentsDetailMockupService {
	return &TalentsDetailMockupService{
		// struct				parameter
		talentDetailRepository: talentDetailRepository,
	}
}

func (tdms TalentsDetailMockupService) GetListTalentDetailMockup(ctx *gin.Context) ([]*models.TalentsDetailMockup, *models.ResponseError) {
	return tdms.talentDetailRepository.GetListTalentDetailMockup(ctx)
}
