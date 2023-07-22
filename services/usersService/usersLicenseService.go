package usersService

import (
	"net/http"

	"codeid.revampacademy/models"
	userrepo "codeid.revampacademy/repositories/usersRepository"
	dbcontext "codeid.revampacademy/repositories/usersRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type UserLicenseService struct {
	UserLicenseRepository *userrepo.UserLicenseRepository
}

func NewUserLicenseService(UserLicenseRepository *userrepo.UserLicenseRepository) *UserLicenseService {
	return &UserLicenseService{
		UserLicenseRepository: UserLicenseRepository,
	}
}

// GetList User License
func (cs UserLicenseService) GetListUserLicense(ctx *gin.Context) ([]*models.UsersUsersLicense, *models.ResponseError) {
	return cs.UserLicenseRepository.GetListUsersLicense(ctx)
}

// Get User License
func (cs UserLicenseService) GetuserLicense(ctx *gin.Context, id int32) (*models.UsersUsersLicense, *models.ResponseError) {
	return cs.UserLicenseRepository.GetUserLicense(ctx, id)
}

// Create User Media
func (cs UserLicenseService) CreateUserLicense(ctx *gin.Context, licenseParams *dbcontext.CreateLicenseParams) (*models.UsersUsersLicense, *models.ResponseError) {
	responseErr := validateLicense(licenseParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return cs.UserLicenseRepository.CreateUserLicense(ctx, licenseParams)
}

// Update Table
func (cs UserLicenseService) UpdateLicense(ctx *gin.Context, licenseParams *dbcontext.CreateLicenseParams, id int64) *models.ResponseError {
	responseErr := validateLicense(licenseParams)
	if responseErr != nil {
		return responseErr
	}

	return cs.UserLicenseRepository.UpdateLicense(ctx, licenseParams)
}

// Delete Table
func (cs UserLicenseService) DeleteLicense(ctx *gin.Context, id int32) *models.ResponseError {
	return cs.UserLicenseRepository.DeleteLicense(ctx, id)
}

// Validate
func validateLicense(licenseParams *dbcontext.CreateLicenseParams) *models.ResponseError {
	if licenseParams.UsliID == 0 {
		return &models.ResponseError{
			Message: "Invalid Media ",
			Status:  http.StatusBadRequest,
		}
	}

	if licenseParams.UsliEntityID == 0 {
		return &models.ResponseError{
			Message: "Required Media ",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}
