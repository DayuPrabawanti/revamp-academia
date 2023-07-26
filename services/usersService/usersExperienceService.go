package usersService

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/usersRepository"
	"codeid.revampacademy/repositories/usersRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type UserExperienceService struct {
	userExperienceRepository *usersRepository.UserExperienceRepository
}

func NewUserExperienceService(UserExperienceRepository *usersRepository.UserExperienceRepository) *UserExperienceService {
	return &UserExperienceService{
		userExperienceRepository: UserExperienceRepository,
	}
}

func (cs UserExperienceService) GetListUserExperience(ctx *gin.Context) ([]*models.UsersUsersExperience, *models.ResponseError) {
	return cs.userExperienceRepository.GetListUserExperience(ctx)
}

func (cs UserExperienceService) GetExperience(ctx *gin.Context, id int32) (*models.UsersUsersExperience, *models.ResponseError) {
	return cs.userExperienceRepository.GetExperience(ctx, id)
}

func (cs UserExperienceService) CreateExperience(ctx *gin.Context, experienceParams *dbContext.CreateExperienceParams) (*models.UsersUsersExperience, *models.ResponseError) {
	responseErr := validateExperience(experienceParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return cs.userExperienceRepository.CreateExperience(ctx, experienceParams)
}

func (cs UserExperienceService) UpdateUserExperience(ctx *gin.Context, experienceParams *dbContext.CreateExperienceParams, id int64) *models.ResponseError {
	responseErr := validateExperience(experienceParams)
	if responseErr != nil {
		return responseErr
	}

	return cs.userExperienceRepository.UpdateUserExperience(ctx, experienceParams)
}

func (cs UserExperienceService) DeleteExperience(ctx *gin.Context, id int32) *models.ResponseError {
	return cs.userExperienceRepository.DeleteExperience(ctx, id)
}

func validateExperience(experienceParams *dbContext.CreateExperienceParams) *models.ResponseError {
	if experienceParams.UsexID == 0 {
		return &models.ResponseError{
			Message: "Invalid Experience Id",
			Status:  http.StatusBadRequest,
		}
	}

	if experienceParams.UsexEntityID == 0 {
		return &models.ResponseError{
			Message: "Entity id cannot 0 value",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}