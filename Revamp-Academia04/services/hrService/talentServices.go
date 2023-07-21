package hrService

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/hrRepository"
	"github.com/gin-gonic/gin"
)

type TalentsMockupService struct {
	talentRepository *hrRepository.TalentsMockupRepository
}

func NewTalentMockupService(talentRepository *hrRepository.TalentsMockupRepository) *TalentsMockupService {
	return &TalentsMockupService{
		// struct				parameter
		talentRepository: talentRepository,
	}
}

func (tms TalentsMockupService) GetListTalentMockup(ctx *gin.Context) ([]*models.TalentsMockup, *models.ResponseError) {
	return tms.talentRepository.GetListTalentMockup(ctx)
}
