package bootcampService

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/bootcampRepository"
	"codeid.revampacademy/repositories/bootcampRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type UsersUsersEducation struct {
	EducationRepository *bootcampRepository.EducationRepository
}

func NewUserEducation(EducationRepository *bootcampRepository.EducationRepository) *UsersUsersEducation {
	return &UsersUsersEducation{
		EducationRepository: EducationRepository,
	}
}

func (ud UsersUsersEducation) GetListEducation(ctx *gin.Context) ([]*models.UsersUsersEducation, *models.ResponseError) {
	return ud.EducationRepository.GetListEducation(ctx)
}

func (ed UsersUsersEducation) UpdateEducation(ctx *gin.Context, EducationParams *dbContext.UsersUsersEducation, id int64) *models.ResponseError {
	responseErr := ValidateEducation(EducationParams)
	if responseErr != nil {
		return responseErr
	}

	return ed.EducationRepository.UpdateEducation(ctx, EducationParams)
}

func ValidateEducation(EducationParams *dbContext.UsersUsersEducation) *models.ResponseError {
	if EducationParams.UsduEntityID == 0 {
		return &models.ResponseError{
			Message: "Invalid Usdu Entity ID",
			Status:  http.StatusBadRequest,
		}
	}

	if EducationParams.UsduDegree == "" {
		return &models.ResponseError{
			Message: "Invalid Usdu Degree",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}

func (ed UsersUsersEducation) GetEducation(ctx *gin.Context, id int64) (*models.UsersUsersEducation, *models.ResponseError) {
	return ed.EducationRepository.GetEducation(ctx, id)
}
