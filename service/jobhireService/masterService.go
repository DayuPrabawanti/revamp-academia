package jobhireService

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/jobhireRepositories"
	"codeid.revampacademy/repositories/jobhireRepositories/dbContext"

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

func (ms MasterService) CreateMasterAddressService(ctx *gin.Context, addressParams *dbContext.CreateMasterAddressParams) (*models.MasterAddress, *models.ResponseError) {
	responseErr := validateParams(addressParams)
	if responseErr != nil {
		return nil, responseErr
	}
	return ms.masterRepo.CreateMasterAddressRepo(ctx, addressParams)
}

func validateParams(addressParams *dbContext.CreateMasterAddressParams) *models.ResponseError {
	if addressParams.AddrID == 0 {
		return &models.ResponseError{
			Message: "Invalid Address id",
			Status:  http.StatusInternalServerError,
		}
	}

	if addressParams.AddrLine1 == "" {
		return &models.ResponseError{
			Message: "Nothing expression in address line 1",
			Status:  http.StatusBadRequest,
		}
	}
	return nil
}
