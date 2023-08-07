package usersService

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/usersRepository"
	"codeid.revampacademy/repositories/usersRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type EditEmailService struct {
	editUserEmailRepository *usersRepository.EditEmailRepository
}

func NewEditUserEmailService(EditUserEmailRepository *usersRepository.EditEmailRepository) *EditEmailService {
	return &EditEmailService{
		editUserEmailRepository: EditUserEmailRepository,
	}
}

// func (cs EditEmailService) AddEmail(ctx *gin.Context, emailParams *dbContext.AddEmailParams, id int32) (*dbContext.AddEmailParams, *models.ResponseError) {
// 	// responseErr := validateEmail(emailParams)
// 	// if responseErr != nil {
// 	// 	return nil, responseErr
// 	// }

// 	return cs.editUserEmailRepository.AddEmail(ctx, emailParams, id)
// }

func (cs EditEmailService) AddEmail(ctx *gin.Context, emailParams *dbContext.AddEmailParams, id int32) *models.ResponseError {
	return cs.editUserEmailRepository.AddEmail(ctx, emailParams, id)
}