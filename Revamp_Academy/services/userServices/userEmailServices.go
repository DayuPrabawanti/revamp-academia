package userservices

import (
	"net/http"

	mod "codeid.revampacademy/models/usersModel"
	usersrepository "codeid.revampacademy/repositories/usersRepository"
	"codeid.revampacademy/repositories/usersRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type UserEmailService struct {
	userEmailRepository *usersrepository.UserEmailRepository
}

func NewUserEmailService(userEmailRepository *usersrepository.UserEmailRepository) *UserEmailService {
	return &UserEmailService{
		userEmailRepository: userEmailRepository,
	}
}

func (cs UserEmailService) GetListUsersEmail(ctx *gin.Context) ([]*mod.UsersUsersEmail, *mod.ResponseError) {
	return cs.userEmailRepository.GetListUserEmail(ctx)
}

func (cs UserEmailService) GetUserEmail(ctx *gin.Context, id int32) (*mod.UsersUsersEmail, *mod.ResponseError) {
	return cs.userEmailRepository.GetUserEmail(ctx, id)
}

func (cs UserEmailService) CreateUserEmail(ctx *gin.Context, emailParams *dbContext.CreateEmailParams) (*mod.UsersUsersEmail, *mod.ResponseError) {
	responseErr := validateEmail(emailParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return cs.userEmailRepository.CreateUserEmail(ctx, emailParams)
}

// Update Table
func (cs UserEmailService) UpdateEmail(ctx *gin.Context, emailParams *dbContext.CreateEmailParams, id int64) *mod.ResponseError {
	responseErr := validateEmail(emailParams)
	if responseErr != nil {
		return responseErr
	}

	return cs.userEmailRepository.UpdateEmail(ctx, emailParams)
}

// Delete Table
func (cs UserEmailService) DeleteEmail(ctx *gin.Context, id int32) *mod.ResponseError {
	return cs.userEmailRepository.DeleteEmail(ctx, id)
}

func validateEmail(emailParams *dbContext.CreateEmailParams) *mod.ResponseError {
	if emailParams.PmailEntityID == 0 {
		return &mod.ResponseError{
			Message: "Invalid Email Adddress",
			Status:  http.StatusBadRequest,
		}
	}

	if emailParams.PmailAddress == "" {
		return &mod.ResponseError{
			Message: "Required Email Address",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}
