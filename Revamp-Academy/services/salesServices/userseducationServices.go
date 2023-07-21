package salesServices

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/dbContext/salesContext"
	"codeid.revampacademy/repositories/salesRepositories"
	"github.com/gin-gonic/gin"
)

type EducationService struct {
	educationRepository *salesRepositories.EducationRepository
}

func NewEducationService(educationRepository *salesRepositories.EducationRepository) *EducationService {
	return &EducationService{
		educationRepository: educationRepository,
	}
}

func (cs EducationService) CreateEducation(ctx *gin.Context, educationParams *salesContext.CreateEducationParams) (*models.UsersUsersEducation, *models.ResponseError) {
	responseErr := validateEducation(educationParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return cs.educationRepository.CreateEducation(ctx, educationParams)
}

func (cs EducationService) GetEducation(ctx *gin.Context, id int32) (*models.UsersUsersEducation, *models.ResponseError) {
	return cs.educationRepository.GetEducation(ctx, id)
}

func validateEducation(educationParams *salesContext.CreateEducationParams) *models.ResponseError {
	if educationParams.UsduID == 0 {
		return &models.ResponseError{
			Message: "Invalid category id",
			Status:  http.StatusBadRequest,
		}
	}

	if educationParams.UsduEntityID == 0 {
		return &models.ResponseError{
			Message: "Invalid category name",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}
