package hrService

import (
	"codeid.revampacademy/models"
	hr "codeid.revampacademy/repositories/hrRepository"
	"github.com/gin-gonic/gin"
)

type TalentsMockupService struct {
	talentRepository *hr.TalentsMockupRepository
}

func NewTalentMockupService(talentRepository *hr.TalentsMockupRepository) *TalentsMockupService {
	return &TalentsMockupService{
		// struct				parameter
		talentRepository: talentRepository,
	}
}

func (tms TalentsMockupService) GetListTalentMockup(ctx *gin.Context) ([]*models.TalentsMockup, *models.ResponseError) {
	return tms.talentRepository.GetListTalentMockup(ctx)
}
