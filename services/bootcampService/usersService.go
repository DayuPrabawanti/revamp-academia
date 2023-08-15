package bootcampService

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/bootcampRepository"
	"codeid.revampacademy/repositories/bootcampRepository/dbContext"
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

func (cs UserService) UpdateUser(ctx *gin.Context, userParams *dbContext.CreateUsersParams, id int64) *models.ResponseError {
	responseErr := validateUser(userParams)
	if responseErr != nil {
		return responseErr
	}

	return cs.userRepository.UpdateUser(ctx, userParams)
}

func validateUser(userParams *dbContext.CreateUsersParams) *models.ResponseError {
	if userParams.UserEntityID == 0 {
		return &models.ResponseError{
			Message: "Invalid User id",
			Status:  http.StatusBadRequest,
		}
	}

	if userParams.UserName == "" {
		return &models.ResponseError{
			Message: "Required Name User",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}

func (cr UserService) GetUser(ctx *gin.Context, id int64) (*models.UsersUser, *models.ResponseError) {
	return cr.userRepository.GetUser(ctx, id)
}
