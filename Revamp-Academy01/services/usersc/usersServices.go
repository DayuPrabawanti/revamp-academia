package usersc

import (
	"net/http"

	"codeid.revamptwo/models/usersm"
	"codeid.revamptwo/repositories/users"
	"codeid.revamptwo/repositories/users/dbContext"
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

func (cs UserService) GetListUser(ctx *gin.Context) ([]*usersm.UsersUser, *usersm.ResponseError) {
	return cs.userRepository.GetListUsers(ctx)
}

func (cs UserService) GetUser(ctx *gin.Context, id int32) (*usersm.UsersUser, *usersm.ResponseError) {
	return cs.userRepository.GetUser(ctx, id)
}

func (cs UserService) CreateUser(ctx *gin.Context, userParams *dbContext.CreateUsersParams) (*usersm.UsersUser, *usersm.ResponseError) {
	responseErr := validateUser(userParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return cs.userRepository.CreateUser(ctx, userParams)
}

func (cs UserService) UpdateUser(ctx *gin.Context, userParams *dbContext.CreateUsersParams, id int64) *usersm.ResponseError {
	responseErr := validateUser(userParams)
	if responseErr != nil {
		return responseErr
	}

	return cs.userRepository.UpdateUser(ctx, userParams)
}

func (cs UserService) DeleteUser(ctx *gin.Context, id int32) *usersm.ResponseError {
	return cs.userRepository.DeleteCategory(ctx, id)
}

func validateUser(userParams *dbContext.CreateUsersParams) *usersm.ResponseError {
	if userParams.UserEntityID == 0 {
		return &usersm.ResponseError{
			Message: "Invalid category id",
			Status:  http.StatusBadRequest,
		}
	}

	if userParams.UserName == "" {
		return &usersm.ResponseError{
			Message: "Required Name Category",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}

