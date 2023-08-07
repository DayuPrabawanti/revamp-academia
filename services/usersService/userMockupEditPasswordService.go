package usersService

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/usersRepository"
	"codeid.revampacademy/repositories/usersRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type EditPasswordService struct {
	editUserpasswordRepository *usersRepository.EditPasswordRepository
}

func NewEditUserPasswordService(EditUserpasswordRepository *usersRepository.EditPasswordRepository) *EditPasswordService {
	return &EditPasswordService{
		editUserpasswordRepository: EditUserpasswordRepository,
	}
}

func (cs EditPasswordService) GetUserPassword(ctx *gin.Context, id int32) (*dbContext.GetUsersPasswordParams, *models.ResponseError) {
	return cs.editUserpasswordRepository.GetPassword(ctx, id)
}

func (cs EditPasswordService) EditUserPassword(ctx *gin.Context, id int32, userPasswordParams *dbContext.EditUsersPasswordParams) *models.ResponseError {
	responseErr := validateUserPassword(userPasswordParams)
	if responseErr != nil {
		return responseErr
	}

	return cs.editUserpasswordRepository.EditPassword(ctx, id, userPasswordParams)
}

func validateUserPassword(usernameParams *dbContext.EditUsersPasswordParams) *models.ResponseError {
	if usernameParams.UserPassword == "" {
		return &models.ResponseError{
			Message: "Required Password",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}