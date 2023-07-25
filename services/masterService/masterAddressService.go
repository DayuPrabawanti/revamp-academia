package masterService

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/masterRepository"
	"github.com/gin-gonic/gin"
)

type MasterAddressService struct {
	masterAddressRepository *masterRepository.MasterAddressRepository
}

func NewMasterAddressService(MasterAddressRepository *masterRepository.MasterAddressRepository) *MasterAddressService {
	return &MasterAddressService{
		masterAddressRepository: MasterAddressRepository,
	}
}

func (cs MasterAddressService) GetListMasterAddress(ctx *gin.Context) ([]*models.MasterAddress, *models.ResponseError) {
	return cs.masterAddressRepository.GetListMasterAddress(ctx)
}
