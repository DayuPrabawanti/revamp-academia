package usersService

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/usersRepository"
	"codeid.revampacademy/repositories/usersRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type UserListProfileService struct {
	userListProfileRepository *usersRepository.UserListProfileRepository
}

func NewUserListProfileService(UserListProfileRepository *usersRepository.UserListProfileRepository) *UserListProfileService {
	return &UserListProfileService{
		userListProfileRepository: UserListProfileRepository,
	}
}

func (cs UserListProfileService) GetProfileUsers(ctx *gin.Context, id int32) (*dbContext.UsernameParams, *models.ResponseError) {
	return cs.userListProfileRepository.GetProfileUser(ctx, id)
}

func (cs UserListProfileService) GetProfileEmail(ctx *gin.Context, id int32) ([]*dbContext.ProfileEmailParams, *models.ResponseError) {
	return cs.userListProfileRepository.GetProfileEmail(ctx, id)
}

func (cs UserListProfileService) GetProfile(ctx *gin.Context, id int32) ([]*dbContext.TampilProfile, *models.ResponseError) {
	return cs.userListProfileRepository.GetProfile(ctx, id)
}