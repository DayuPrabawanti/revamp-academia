package usersService

import (
	"net/http"

	"codeid.revampacademy/models/usersModel"
	"codeid.revampacademy/repositories/users"
	"codeid.revampacademy/repositories/users/dbContext"
	"github.com/gin-gonic/gin"
)

type UserEmailService struct {
	userEmailRepository *users.UserEmailRepository
}

func NewUserEmailService(UserEmailRepository *users.UserEmailRepository) *UserEmailService {
	return &UserEmailService{
		userEmailRepository: UserEmailRepository,
	}
}

func (cs UserEmailService) GetListUsersEmail(ctx *gin.Context) ([]*usersModel.UsersUsersEmail, *usersModel.ResponseError) {
	return cs.userEmailRepository.GetListUsersEmail(ctx)
}

func (cs UserEmailService) GetEmail(ctx *gin.Context, id int32) (*usersModel.UsersUsersEmail, *usersModel.ResponseError) {
	return cs.userEmailRepository.GetEmail(ctx, id)
}

func (cs UserEmailService) CreateEmail(ctx *gin.Context, emailParams *dbContext.CreateEmailParams) (*usersModel.UsersUsersEmail, *usersModel.ResponseError) {
	responseErr := validateEmail(emailParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return cs.userEmailRepository.CreateEmail(ctx, emailParams)
}

func (cs UserEmailService) UpdateEmail(ctx *gin.Context, emailParams *dbContext.CreateEmailParams, id int64) *usersModel.ResponseError {
	responseErr := validateEmail(emailParams)
	if responseErr != nil {
		return responseErr
	}

	return cs.userEmailRepository.UpdateEmail(ctx, emailParams)
}

func (cs UserEmailService) DeleteEmail(ctx *gin.Context, id int32) *usersModel.ResponseError {
	return cs.userEmailRepository.DeleteEmail(ctx, id)
}

func validateEmail(emailParams *dbContext.CreateEmailParams) *usersModel.ResponseError {
	if emailParams.PmailEntityID == 0 {
		return &usersModel.ResponseError{
			Message: "Invalid Email Adddress",
			Status:  http.StatusBadRequest,
		}
	}

	if emailParams.PmailAddress == "" {
		return &usersModel.ResponseError{
			Message: "Required Email Address",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}