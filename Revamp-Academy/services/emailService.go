package services

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories"
	"github.com/gin-gonic/gin"
)

type UserEmailService struct {
	userEmailRepository *repositories.UserEmailRepository
}

func NewUserEmailService(UserEmailRepository *repositories.UserEmailRepository) *UserEmailService {
	return &UserEmailService{
		userEmailRepository: UserEmailRepository,
	}
}

func (cs UserEmailService) GetListUsersEmail(ctx *gin.Context) ([]*models.UsersUsersEmail, *models.ResponseError) {
	return cs.userEmailRepository.GetListUsersEmail(ctx)
}
