package usersService

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/usersRepository"
	"codeid.revampacademy/repositories/usersRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	userRepository *usersRepository.UserRepository
}

func NewUserService(UserRepository *usersRepository.UserRepository) *UserService {
	return &UserService{
		userRepository: UserRepository,
	}
}

func (cs UserService) GetListUser(ctx *gin.Context) ([]*models.UsersUser, *models.ResponseError) {
	return cs.userRepository.GetListUsers(ctx)
}

func (cs UserService) GetUser(ctx *gin.Context, id int32) (*models.UsersUser, *models.ResponseError) {
	return cs.userRepository.GetUser(ctx, id)
}

func (cs UserService) CreateUser(ctx *gin.Context, userParams *dbContext.CreateUsersParams) (*models.UsersUser, *models.ResponseError) {
	responseErr := validateUser(userParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return cs.userRepository.CreateUser(ctx, userParams)
}

func (cs UserService) UpdateUser(ctx *gin.Context, userParams *dbContext.CreateUsersParams, id int64) *models.ResponseError {
	responseErr := validateUser(userParams)
	if responseErr != nil {
		return responseErr
	}

	return cs.userRepository.UpdateUser(ctx, userParams)
}

func (cs UserService) DeleteUser(ctx *gin.Context, id int32) *models.ResponseError {
	return cs.userRepository.DeleteUser(ctx, id)
}

// func (cs UserService) GetProfileUser(ctx *gin.Context, id int32) ([]*dbContext.UsernameParams, *models.ResponseError) {
// 	return cs.userRepository.GetProfileUser(ctx, id)
// }


func validateUser(userParams *dbContext.CreateUsersParams) *models.ResponseError {
	if userParams.UserEntityID == 0 {
		return &models.ResponseError{
			Message: "Invalid User id",
			Status:  http.StatusBadRequest,
		}
	}

	if userParams.UserName.Valid == false {
		return &models.ResponseError{
			Message: "Required Username",
			Status:  http.StatusBadRequest,
		}
	}

	if userParams.UserPassword.Valid == false {
		return &models.ResponseError{
			Message: "Required Password",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}

