package services

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories"
	"github.com/gin-gonic/gin"
)

type TalentsMockupService struct {
	talentRepository *repositories.TalentsMockupRepository
}

func NewTalentMockupService(talentRepository *repositories.TalentsMockupRepository) *TalentsMockupService {
	return &TalentsMockupService{
		// struct				parameter
		talentRepository: talentRepository,
	}
}

func (tms TalentsMockupService) GetListTalentMockup(ctx *gin.Context) ([]*models.TalentsMockup, *models.ResponseError) {
	return tms.talentRepository.GetListTalentMockup(ctx)
}
