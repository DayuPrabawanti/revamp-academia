package usersService

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/usersRepository"
	"codeid.revampacademy/repositories/usersRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type UserEmailService struct {
	userEmailRepository *usersRepository.UserEmailRepository
}

func NewUserEmailService(UserEmailRepository *usersRepository.UserEmailRepository) *UserEmailService {
	return &UserEmailService{
		userEmailRepository: UserEmailRepository,
	}
}

func (cs UserEmailService) GetListUsersEmail(ctx *gin.Context) ([]*models.UsersUsersEmail, *models.ResponseError) {
	return cs.userEmailRepository.GetListUsersEmail(ctx)
}

func (cs UserEmailService) GetEmail(ctx *gin.Context, id int32) ([]*dbContext.Email, *models.ResponseError) {
	return cs.userEmailRepository.GetEmail(ctx, id)
}

func (cs UserEmailService) CreateEmail(ctx *gin.Context, emailParams *dbContext.CreateEmailParams) (*models.UsersUsersEmail, *models.ResponseError) {
	responseErr := validateEmail(emailParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return cs.userEmailRepository.CreateEmail(ctx, emailParams)
}

func (cs UserEmailService) UpdateEmail(ctx *gin.Context, emailParams *dbContext.CreateEmailParams, id int64) *models.ResponseError {
	responseErr := validateEmail(emailParams)
	if responseErr != nil {
		return responseErr
	}

	return cs.userEmailRepository.UpdateEmail(ctx, emailParams)
}

func (cs UserEmailService) DeleteEmail(ctx *gin.Context, id int32) *models.ResponseError {
	return cs.userEmailRepository.DeleteEmail(ctx, id)
}


// func (cs UserEmailService) EmailProfile(ctx *gin.Context, id int32) ([]*models.UsersUsersEmail, *models.ResponseError) {
// 	return cs.userEmailRepository.Email(ctx, id)
// }


func validateEmail(emailParams *dbContext.CreateEmailParams) *models.ResponseError {
	if emailParams.PmailEntityID == 0 {
		return &models.ResponseError{
			Message: "Invalid Email Id",
			Status:  http.StatusBadRequest,
		}
	}

	if emailParams.PmailAddress.Valid == false {
		return &models.ResponseError{
			Message: "Required Email Address",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}