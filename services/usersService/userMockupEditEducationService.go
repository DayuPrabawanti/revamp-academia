package usersService

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/usersRepository"
	"codeid.revampacademy/repositories/usersRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type EditEducationService struct {
	editUserEducationRepository *usersRepository.EditEducationRepository
}

func NewEditUserEducationService(EditUserEducationRepository *usersRepository.EditEducationRepository) *EditEducationService {
	return &EditEducationService{
		editUserEducationRepository: EditUserEducationRepository,
	}
}

func (cs EditEducationService) AddEducation(ctx *gin.Context, educationParams *dbContext.AddEducationParams, id int32) *models.ResponseError {
	return cs.editUserEducationRepository.AddEducation(ctx, educationParams, id)
}