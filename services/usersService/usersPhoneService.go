package usersService

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/usersRepository"
	"codeid.revampacademy/repositories/usersRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type UserPhoneService struct {
	userPhoneRepository *usersRepository.UserPhoneRepository
}

func NewUserPhoneService(UserPhoneRepository *usersRepository.UserPhoneRepository) *UserPhoneService {
	return &UserPhoneService{
		userPhoneRepository: UserPhoneRepository,
	}
}

func (cs UserPhoneService) GetListUsersPhone(ctx *gin.Context) ([]*models.UsersUsersPhone, *models.ResponseError) {
	return cs.userPhoneRepository.GetListUsersPhone(ctx)
}

func (cs UserPhoneService) GetPhone(ctx *gin.Context, id int32) (*models.UsersUsersPhone, *models.ResponseError) {
	return cs.userPhoneRepository.GetPhone(ctx, id)
}

func (cs UserPhoneService) CreatePhones(ctx *gin.Context, phoneParams *dbContext.CreatePhonesParams) (*models.UsersUsersPhone, *models.ResponseError) {
	responseErr := validatePhone(phoneParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return cs.userPhoneRepository.CreatePhones(ctx, phoneParams)
}

func (cs UserPhoneService) UpdatePhone(ctx *gin.Context, phoneParams *dbContext.CreatePhonesParams, id int64) *models.ResponseError {
	responseErr := validatePhone(phoneParams)
	if responseErr != nil {
		return responseErr
	}

	return cs.userPhoneRepository.UpdatePhone(ctx, phoneParams)
}

func (cs UserPhoneService) DeletePhones(ctx *gin.Context, id int32) *models.ResponseError {
	return cs.userPhoneRepository.DeletePhones(ctx, id)
}

func validatePhone(phoneParams *dbContext.CreatePhonesParams) *models.ResponseError {
	if phoneParams.UspoEntityID == 0 {
		return &models.ResponseError{
			Message: "Invalid Phone Id",
			Status:  http.StatusBadRequest,
		}
	}

	if phoneParams.UspoNumber == "" {
		return &models.ResponseError{
			Message: "Required Phone Number",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}