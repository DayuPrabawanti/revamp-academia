package usersService

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/usersRepository"
	"github.com/gin-gonic/gin"
)

type UserAddressService struct {
	userAddressRepository *usersRepository.UserAddressRepository
}

func NewUserAddressService(UserAddressRepository *usersRepository.UserAddressRepository) *UserAddressService {
	return &UserAddressService{
		userAddressRepository: UserAddressRepository,
	}
}

func (cs UserAddressService) GetListUserAddress(ctx *gin.Context) ([]*models.UsersUsersAddress, *models.ResponseError) {
	return cs.userAddressRepository.GetListUserAddress(ctx)
}
