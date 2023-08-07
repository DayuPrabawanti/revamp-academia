package usersService

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/usersRepository"
	"codeid.revampacademy/repositories/usersRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type EditExperienceService struct {
	editUserExperienceRepository *usersRepository.EditExperienceRepository
}

func NewEditUserExperienceService(EditUserExperienceRepository *usersRepository.EditExperienceRepository) *EditExperienceService {
	return &EditExperienceService{
		editUserExperienceRepository: EditUserExperienceRepository,
	}
}

func (cs EditExperienceService) AddExperience(ctx *gin.Context, experienceParams *dbContext.AddExperienceParams, id int32) *models.ResponseError {
	return cs.editUserExperienceRepository.AddExperience(ctx, experienceParams, id)
}