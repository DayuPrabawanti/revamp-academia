package services

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories"
	"github.com/gin-gonic/gin"

)

type MasterIndustryService struct {
	MasterIndustryServiceRepo *repositories.MasterIndustryRepository
}

func NewMasterIndustryService(MasterIndustryServiceRepo *repositories.MasterIndustryRepository) *MasterIndustryService {
	return &MasterIndustryService{
		MasterIndustryServiceRepo: MasterIndustryServiceRepo,
	}
}

func (mastry MasterIndustryService) GetMasterIndustryService(ctx *gin.Context, InduCodeID int32) (*models.MasterIndustry, *models.ResponseError) {
	return mastry.MasterIndustryServiceRepo.GetMasterIndustryRepo(ctx, InduCodeID)
}

	func (mastry MasterIndustryService) ListMasterIndustryService(ctx *gin.Context) ([]*models.MasterIndustry, *models.ResponseError) {
					return mastry.MasterIndustryServiceRepo.ListMasterIndustryRepo(ctx)
	}

