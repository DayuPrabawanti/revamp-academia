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

func (tl TalentsMockupService) SearchTalent(ctx *gin.Context, userName, userSkill, batchName, status string) ([]models.TalentsMockup, *models.ResponseError) {
	// Perform validation, if needed, for batchName and status
	// If validation fails, return appropriate response error

	return tl.talentRepository.SearchTalent(ctx, userName, userSkill, batchName, status)
}

func (tl TalentsMockupService) PagingTalent(ctx *gin.Context, offset, pageSize int) ([]models.TalentsMockup, *models.ResponseError) {

	return tl.talentRepository.PagingTalent(ctx, offset, pageSize)
}
