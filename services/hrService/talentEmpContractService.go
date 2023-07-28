package hrService

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/hrRepository"
	"codeid.revampacademy/repositories/hrRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type TalentClientContractService struct {
	talentClientContractRepository *hrRepository.TalentClientContractRepository
}

func NewTalentClientContractService(TalentClientContractRepository *hrRepository.TalentClientContractRepository) *TalentClientContractService {
	return &TalentClientContractService{
		talentClientContractRepository: TalentClientContractRepository,
	}
}

func (ccs TalentClientContractService) GetTalentClientContract(ctx *gin.Context, id int64) (*models.TalentClientContractGetUpdate, *models.ResponseError) {
	return ccs.talentClientContractRepository.GetTalentClientContract(ctx, id)
}

func (ccs TalentClientContractService) UpdateTalentClientContract(ctx *gin.Context, talentClientContractParams *dbContext.UpdateTalentClientContractParams, id int64) *models.ResponseError {
	responseErr := validateTalentClientContract(talentClientContractParams)
	if responseErr != nil {
		return responseErr
	}

	return ccs.talentClientContractRepository.UpdateTalentClientContract(ctx, talentClientContractParams)
}

func validateTalentClientContract(talentClientContractParams *dbContext.UpdateTalentClientContractParams) *models.ResponseError {
	if talentClientContractParams.EccoEntityID == 0 {
		return &models.ResponseError{
			Message: "Invalid ecco id",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}
