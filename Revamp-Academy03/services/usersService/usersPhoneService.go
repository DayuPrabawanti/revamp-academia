package usersService

import (
	"net/http"

	"codeid.revampacademy/models/usersModel"
	"codeid.revampacademy/repositories/users"
	"codeid.revampacademy/repositories/users/dbContext"
	"github.com/gin-gonic/gin"
)

type UserPhoneService struct {
	userPhoneRepository *users.UserPhoneRepository
}

func NewUserPhoneService(UserPhoneRepository *users.UserPhoneRepository) *UserPhoneService {
	return &UserPhoneService{
		userPhoneRepository: UserPhoneRepository,
	}
}

func (cs UserPhoneService) GetListUsersPhone(ctx *gin.Context) ([]*usersModel.UsersUsersPhone, *usersModel.ResponseError) {
	return cs.userPhoneRepository.GetListUsersPhone(ctx)
}

func (cs UserPhoneService) GetPhone(ctx *gin.Context, id int32) (*usersModel.UsersUsersPhone, *usersModel.ResponseError) {
	return cs.userPhoneRepository.GetPhone(ctx, id)
}

func (cs UserPhoneService) CreatePhones(ctx *gin.Context, phoneParams *dbContext.CreatePhonesParams) (*usersModel.UsersUsersPhone, *usersModel.ResponseError) {
	responseErr := validatePhone(phoneParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return cs.userPhoneRepository.CreatePhones(ctx, phoneParams)
}

func (cs UserPhoneService) UpdatePhone(ctx *gin.Context, phoneParams *dbContext.CreatePhonesParams, id int64) *usersModel.ResponseError {
	responseErr := validatePhone(phoneParams)
	if responseErr != nil {
		return responseErr
	}

	return cs.userPhoneRepository.UpdatePhone(ctx, phoneParams)
}

func (cs UserPhoneService) DeletePhones(ctx *gin.Context, id int32) *usersModel.ResponseError {
	return cs.userPhoneRepository.DeletePhones(ctx, id)
}

func validatePhone(phoneParams *dbContext.CreatePhonesParams) *usersModel.ResponseError {
	if phoneParams.UspoEntityID == 0 {
		return &usersModel.ResponseError{
			Message: "Invalid Email Adddress",
			Status:  http.StatusBadRequest,
		}
	}

	if phoneParams.UspoNumber == "" {
		return &usersModel.ResponseError{
			Message: "Required Phone Number",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}