package boocampService

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/bootcampRepository"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	userRepository *bootcampRepository.UserRepository
}

func NewUserService(UserRepository *bootcampRepository.UserRepository) *UserService {
	return &UserService{
		userRepository: UserRepository,
	}
}

func (cs UserService) GetListUser(ctx *gin.Context) ([]*models.UsersUser, *models.ResponseError) {
	return cs.userRepository.GetListUsers(ctx)
}
