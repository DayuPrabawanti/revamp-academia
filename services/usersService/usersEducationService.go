package usersService

import (
	"net/http"

	"codeid.revampacademy/models"
	userrepo "codeid.revampacademy/repositories/usersRepository"
	dbcontext "codeid.revampacademy/repositories/usersRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type UserEducationService struct {
	UserEducationRepository *userrepo.UserEducationRepository
}

func NewUserEducationService(UserEducationRepository *userrepo.UserEducationRepository) *UserEducationService {
	return &UserEducationService{
		UserEducationRepository: UserEducationRepository,
	}
}

// GetList User Education
func (cs UserEducationService) GetListUsersEducation(ctx *gin.Context) ([]*models.UsersUsersEducation, *models.ResponseError) {
	return cs.UserEducationRepository.GetListUsersEducation(ctx)
}

// Get User Education
func (cs UserEducationService) GetUserEducation(ctx *gin.Context, id int32) (*models.UsersUsersEducation, *models.ResponseError) {
	return cs.UserEducationRepository.GetUserEducation(ctx, id)
}

// Create User Media
func (cs UserEducationService) CreateUserEducation(ctx *gin.Context, educationPrams *dbcontext.CreateEducationParams) (*models.UsersUsersEducation, *models.ResponseError) {
	responseErr := validateEducation(educationPrams)
	if responseErr != nil {
		return nil, responseErr
	}

	return cs.UserEducationRepository.CreateUserEducation(ctx, educationPrams)
}

// Update Table
func (cs UserEducationService) UpdateEducation(ctx *gin.Context, educationPrams *dbcontext.CreateEducationParams, id int64) *models.ResponseError {
	responseErr := validateEducation(educationPrams)
	if responseErr != nil {
		return responseErr
	}

	return cs.UserEducationRepository.UpdateEducation(ctx, educationPrams)
}

// Delete Table
func (cs UserEducationService) DeleteEducation(ctx *gin.Context, id int32) *models.ResponseError {
	return cs.UserEducationRepository.DeleteEducation(ctx, id)
}

// Validate
func validateEducation(educationPrams *dbcontext.CreateEducationParams) *models.ResponseError {
	if educationPrams.UsduID == 0 {
		return &models.ResponseError{
			Message: "Invalid Media ",
			Status:  http.StatusBadRequest,
		}
	}

	if educationPrams.UsduEntityID == 0 {
		return &models.ResponseError{
			Message: "Required Media ",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}
