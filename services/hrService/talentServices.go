package hrService

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/hrRepository"
	"codeid.revampacademy/repositories/hrRepository/dbContext"
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

func (tl TalentsMockupService) SearchTalent(ctx *gin.Context, userName, skillName, batchStatus string) ([]models.TalentsMockup, *models.ResponseError) {
	// Perform validation, if needed, for batchName and status
	// If validation fails, return appropriate response error

	return tl.talentRepository.SearchTalent(ctx, userName, skillName, batchStatus)
}

func (tl TalentsMockupService) PagingTalent(ctx *gin.Context, offset, pageSize int) ([]models.TalentsMockup, *models.ResponseError) {

	return tl.talentRepository.PagingTalent(ctx, offset, pageSize)
}

func (bs TalentsMockupService) GetBatch(ctx *gin.Context, id int64) (*models.BootcampBatch, *models.ResponseError) {
	return bs.talentRepository.GetBatch(ctx, id)
}

func (bs TalentsMockupService) UpdateBatch(ctx *gin.Context, batchParams *dbContext.UpdateBatchParams, id int64) *models.ResponseError {
	responseErr := validateBatch(batchParams)
	if responseErr != nil {
		return responseErr
	}

	return bs.talentRepository.UpdateBatch(ctx, batchParams)
}

func validateBatch(batchParams *dbContext.UpdateBatchParams) *models.ResponseError {
	if batchParams.BatchID == 0 {
		return &models.ResponseError{
			Message: "Invalid batch id",
			Status:  http.StatusBadRequest,
		}
	}

	if batchParams.BatchStatus == "" {
		return &models.ResponseError{
			Message: "Invalid batch name",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}
