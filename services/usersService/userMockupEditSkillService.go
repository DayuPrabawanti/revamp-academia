package usersService

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/usersRepository"
	"codeid.revampacademy/repositories/usersRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type EditSkillService struct {
	editUserSkillRepository *usersRepository.EditSkillRepository
}

func NewEditUserSkillService(EditUserSkillRepository *usersRepository.EditSkillRepository) *EditSkillService {
	return &EditSkillService{
		editUserSkillRepository: EditUserSkillRepository,
	}
}

func (cs EditSkillService) AddSkill(ctx *gin.Context, skillParams *dbContext.AddSkillParams, id int32) *models.ResponseError {
	return cs.editUserSkillRepository.AddSkill(ctx, skillParams, id)
}