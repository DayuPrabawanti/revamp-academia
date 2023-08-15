package bootcampService

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/bootcampRepository"
	"github.com/gin-gonic/gin"
)

type UserEmailService struct {
	userEmailRepository *bootcampRepository.UserEmailRepository
}

func NewUserEmailService(UserEmailRepository *bootcampRepository.UserEmailRepository) *UserEmailService {
	return &UserEmailService{
		userEmailRepository: UserEmailRepository,
	}
}

func (cs UserEmailService) GetListUsersEmail(ctx *gin.Context) ([]*models.UsersUsersEmail, *models.ResponseError) {
	return cs.userEmailRepository.GetListUsersEmail(ctx)
}
