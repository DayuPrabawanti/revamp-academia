package usersService

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/usersRepository"
	"codeid.revampacademy/repositories/usersRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type EditUsernameService struct {
	editUsernameRepository *usersRepository.EditUsernameRepository
}

func NewEditUsernameService(EditUsernameRepository *usersRepository.EditUsernameRepository) *EditUsernameService {
	return &EditUsernameService{
		editUsernameRepository: EditUsernameRepository,
	}
}

func (cs EditUsernameService) GetUsername(ctx *gin.Context, id int32) (*dbContext.GetUsernameParams, *models.ResponseError) {
	return cs.editUsernameRepository.GetUsername(ctx, id)
}

func (cs EditUsernameService) EditUsername(ctx *gin.Context, usernameParams *dbContext.EditUsernameParams, id int64) *models.ResponseError {
	responseErr := validateUsername(usernameParams)
	if responseErr != nil {
		return responseErr
	}

	return cs.editUsernameRepository.EditUsername(ctx, usernameParams)
}

func validateUsername(usernameParams *dbContext.EditUsernameParams) *models.ResponseError {
	if usernameParams.UserName == "" {
		return &models.ResponseError{
			Message: "Required Username",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}