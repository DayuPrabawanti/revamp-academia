package jobhireService

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/jobhireRepositories"

	// "codeid.revampacademy/jobhireRepositories"
	"github.com/gin-gonic/gin"
)

type MasterService struct {
	masterRepo *jobhireRepositories.MasterRepo
}

func NewMasterService(masterRepo *jobhireRepositories.MasterRepo) *MasterService {
	return &MasterService{
		masterRepo: masterRepo,
	}
}

func (ms MasterService) GetListMasterAddress(ctx *gin.Context) ([]*models.MasterAddress, *models.ResponseError) {
	return ms.masterRepo.GetListMasterAddress(ctx)
}

func (ms MasterService) GetListMasterCity(ctx *gin.Context) ([]*models.MasterCity, *models.ResponseError) {
	return ms.masterRepo.GetListMasterCity(ctx)
}
