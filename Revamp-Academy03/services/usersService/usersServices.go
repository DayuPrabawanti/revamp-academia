package usersService

import (
	"net/http"

	"codeid.revampacademy/models/usersModel"
	"codeid.revampacademy/repositories/users"
	"codeid.revampacademy/repositories/users/dbContext"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	userRepository *users.UserRepository
}

func NewUserService(UserRepository *users.UserRepository) *UserService {
	return &UserService{
		userRepository: UserRepository,
	}
}

func (cs UserService) GetListUser(ctx *gin.Context) ([]*usersModel.UsersUser, *usersModel.ResponseError) {
	return cs.userRepository.GetListUsers(ctx)
}

func (cs UserService) GetUser(ctx *gin.Context, id int32) (*usersModel.UsersUser, *usersModel.ResponseError) {
	return cs.userRepository.GetUser(ctx, id)
}

func (cs UserService) CreateUser(ctx *gin.Context, userParams *dbContext.CreateUsersParams) (*usersModel.UsersUser, *usersModel.ResponseError) {
	responseErr := validateUser(userParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return cs.userRepository.CreateUser(ctx, userParams)
}

func (cs UserService) UpdateUser(ctx *gin.Context, userParams *dbContext.CreateUsersParams, id int64) *usersModel.ResponseError {
	responseErr := validateUser(userParams)
	if responseErr != nil {
		return responseErr
	}

	return cs.userRepository.UpdateUser(ctx, userParams)
}

func (cs UserService) DeleteUser(ctx *gin.Context, id int32) *usersModel.ResponseError {
	return cs.userRepository.DeleteCategory(ctx, id)
}

func validateUser(userParams *dbContext.CreateUsersParams) *usersModel.ResponseError {
	if userParams.UserEntityID == 0 {
		return &usersModel.ResponseError{
			Message: "Invalid category id",
			Status:  http.StatusBadRequest,
		}
	}

	if userParams.UserName == "" {
		return &usersModel.ResponseError{
			Message: "Required Name Category",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}

