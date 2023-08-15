package bootcampService

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/bootcampRepository"
	"codeid.revampacademy/repositories/bootcampRepository/dbContext"
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

func (cr UserEmailService) GetListUsersEmail(ctx *gin.Context) ([]*models.UsersUsersEmail, *models.ResponseError) {
	return cr.userEmailRepository.GetListUsersEmail(ctx)
}

func (cr UserEmailService) UpdateEmail(ctx *gin.Context, emailParams *dbContext.CreateEmailParams, id int64) *models.ResponseError {
	responseErr := validateEmail(emailParams)
	if responseErr != nil {
		return responseErr
	}

	return cr.userEmailRepository.UpdateEmail(ctx, emailParams)
}

func validateEmail(emailParams *dbContext.CreateEmailParams) *models.ResponseError {
	if emailParams.PmailEntityID == 0 {
		return &models.ResponseError{
			Message: "Invalid Email Adddress",
			Status:  http.StatusBadRequest,
		}
	}

	if emailParams.PmailAddress == "" {
		return &models.ResponseError{
			Message: "Required Email Address",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}

func (cr UserEmailService) GetUSerEmail(ctx *gin.Context, id int64) (*models.UsersUsersEmail, *models.ResponseError) {
	return cr.userEmailRepository.GetUSerEmail(ctx, id)
}
