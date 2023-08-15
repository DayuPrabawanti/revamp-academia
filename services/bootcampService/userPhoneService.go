package bootcampService

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/bootcampRepository"
	"codeid.revampacademy/repositories/bootcampRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type UserPhoneService struct {
	userPhoneRepository *bootcampRepository.UserPhoneRepository
}

func NewUserPhoneService(UserPhoneRepository *bootcampRepository.UserPhoneRepository) *UserPhoneService {
	return &UserPhoneService{
		userPhoneRepository: UserPhoneRepository,
	}
}

func (cs UserPhoneService) GetListUsersPhone(ctx *gin.Context) ([]*models.UsersUsersPhone, *models.ResponseError) {
	return cs.userPhoneRepository.GetListUsersPhone(ctx)
}

func (cs UserPhoneService) UpdatePhone(ctx *gin.Context, phoneParams *dbContext.CreatePhonesParams, id int64) *models.ResponseError {
	responseErr := validatePhone(phoneParams)
	if responseErr != nil {
		return responseErr
	}

	return cs.userPhoneRepository.UpdatePhone(ctx, phoneParams)
}
func validatePhone(phoneParams *dbContext.CreatePhonesParams) *models.ResponseError {
	if phoneParams.UspoEntityID == 0 {
		return &models.ResponseError{
			Message: "Invalid Email Adddress",
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

func (cr UserPhoneService) GetUserPhone(ctx *gin.Context, id int64) (*models.UsersUsersPhone, *models.ResponseError) {
	return cr.userPhoneRepository.GetUserPhone(ctx, id)
}
