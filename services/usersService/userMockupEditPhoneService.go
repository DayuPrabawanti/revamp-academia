package usersService

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/usersRepository"
	"codeid.revampacademy/repositories/usersRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type EditPhoneService struct {
	editUserPhoneRepository *usersRepository.EditPhoneRepository
}

func NewEditUserPhoneService(EditUserPhoneRepository *usersRepository.EditPhoneRepository) *EditPhoneService {
	return &EditPhoneService{
		editUserPhoneRepository: EditUserPhoneRepository,
	}
}

func (cs EditPhoneService) AddPhone(ctx *gin.Context, phoneParams *dbContext.AddPhonesParams, id int32) *models.ResponseError {
	return cs.editUserPhoneRepository.AddPhone(ctx, phoneParams, id)
}