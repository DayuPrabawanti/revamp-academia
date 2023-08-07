package usersService

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/usersRepository"
	"codeid.revampacademy/repositories/usersRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type AddAddressService struct {
	addAddressRepository *usersRepository.EditAddressRepository
}

func NewAddAddressService(addAddressRepository *usersRepository.EditAddressRepository) *AddAddressService {
	return &AddAddressService{
		addAddressRepository: addAddressRepository,
	}
}

func (cs *AddAddressService) AddAddress(ctx *gin.Context, masterAddressParams *dbContext.AddAddressParams, userAddressParams *dbContext.AddressParams, id int32)  *models.ResponseError {
	return cs.addAddressRepository.AddAddress(ctx, masterAddressParams, userAddressParams, id)
}