package hrsSc

import (
	"codeid.revamptwo/models/hrsMdl"
	"codeid.revamptwo/repositories/hrs"
	"github.com/gin-gonic/gin"
)

type TalentsDetailMockupService struct {
	talentDetailRepository *hrs.TalentsDetailMockupRepository
}

func NewTalentDetailMockupService(talentDetailRepository *hrs.TalentsDetailMockupRepository) *TalentsDetailMockupService {
	return &TalentsDetailMockupService{
		// struct				parameter
		talentDetailRepository: talentDetailRepository,
	}
}

func (cs TalentsDetailMockupService) GetListTalentDetailMockup(ctx *gin.Context) ([]*hrsMdl.TalentsDetailMockup, *hrsMdl.ResponseError) {
	return cs.talentDetailRepository.GetListTalentDetailMockup(ctx)
}
