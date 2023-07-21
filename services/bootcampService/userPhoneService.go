package bootcampService

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/bootcampRepository"
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
