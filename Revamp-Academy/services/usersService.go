package services

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(UserRepository *repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: UserRepository,
	}
}

func (cs UserService) GetListUser(ctx *gin.Context) ([]*models.UsersUser, *models.ResponseError) {
	return cs.userRepository.GetListUsers(ctx)
}
