package hrService

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/hrRepository"
	"codeid.revampacademy/repositories/hrRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type TalentsDetailMockupService struct {
	talentDetailRepository *hrRepository.TalentsDetailMockupRepository
}

func NewTalentDetailMockupService(talentDetailRepository *hrRepository.TalentsDetailMockupRepository) *TalentsDetailMockupService {
	return &TalentsDetailMockupService{
		// struct				parameter
		talentDetailRepository: talentDetailRepository,
	}
}

func (tdms TalentsDetailMockupService) GetListTalentDetailMockup(ctx *gin.Context) ([]*models.TalentsDetailMockup, *models.ResponseError) {
	return tdms.talentDetailRepository.GetListTalentDetailMockup(ctx)
}

func (tdms TalentsDetailMockupService) GetTalentDetail(ctx *gin.Context, id int64) (*models.TalentsDetailMockup, *models.ResponseError) {
	return tdms.talentDetailRepository.GetTalentDetail(ctx, id)
}

func (tdms TalentsDetailMockupService) SearchTalentDetail(ctx *gin.Context, clitName string) ([]models.TalentDetailSearchUpdate, *models.ResponseError) {
	return tdms.talentDetailRepository.SearchTalentDetail(ctx, clitName)
}

func (bs TalentsDetailMockupService) UpdateSwitch(ctx *gin.Context, switchParams *dbContext.UpdateSwitchParams, id int64) *models.ResponseError {

	return bs.talentDetailRepository.UpdateSwitch(ctx, switchParams)
}
