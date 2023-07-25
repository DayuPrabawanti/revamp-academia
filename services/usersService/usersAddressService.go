package usersService

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/usersRepository"
	"codeid.revampacademy/repositories/usersRepository/dbContext"
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

func (cs UserAddressService) GetAddress(ctx *gin.Context, id int32) (*models.UsersUsersAddress, *models.ResponseError) {
	return cs.userAddressRepository.GetAddress(ctx, id)
}

func (cs UserAddressService) CreateAddrees(ctx *gin.Context, userAddressParams *dbContext.CreateAddreesParams) (*models.UsersUsersAddress, *models.ResponseError) {
	responseErr := validateUserAddress(userAddressParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return cs.userAddressRepository.CreateAddrees(ctx, userAddressParams)
}



func validateUserAddress(userAddressParams *dbContext.CreateAddreesParams) *models.ResponseError {
	if userAddressParams.EtadAddrID == 0 {
		return &models.ResponseError{
			Message: "Invalid User id",
			Status:  http.StatusBadRequest,
		}
	}
	return nil

}