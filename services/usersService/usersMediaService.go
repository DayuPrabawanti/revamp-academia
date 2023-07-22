package usersService

import (
	"net/http"

	"codeid.revampacademy/models"
	userrepo "codeid.revampacademy/repositories/usersRepository"
	dbcontext "codeid.revampacademy/repositories/usersRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type UserMediaService struct {
	userMediaRepository *userrepo.UserMediaRepository
}

func NewUserMediaService(userMediaRepository *userrepo.UserMediaRepository) *UserMediaService {
	return &UserMediaService{
		userMediaRepository: userMediaRepository,
	}
}

// GetList User Media
func (cs UserMediaService) GetListUserMedia(ctx *gin.Context) ([]*models.UsersUsersMedia, *models.ResponseError) {
	return cs.userMediaRepository.GetListUserMedia(ctx)
}

// Get User Media
func (cs UserMediaService) GetUserMedia(ctx *gin.Context, id int32) (*models.UsersUsersMedia, *models.ResponseError) {
	return cs.userMediaRepository.GetUserMedia(ctx, id)
}

// Create User Media
func (cs UserMediaService) CreateUserMedia(ctx *gin.Context, mediaParams *dbcontext.CreateMediaParams) (*models.UsersUsersMedia, *models.ResponseError) {
	responseErr := validateMedia(mediaParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return cs.userMediaRepository.CreateUserMedia(ctx, mediaParams)
}

// Update Table
func (cs UserMediaService) UpdateMedia(ctx *gin.Context, mediaParams *dbcontext.CreateMediaParams, id int64) *models.ResponseError {
	responseErr := validateMedia(mediaParams)
	if responseErr != nil {
		return responseErr
	}

	return cs.userMediaRepository.UpdateMedia(ctx, mediaParams)
}

// Delete Table
func (cs UserMediaService) DeleteMedia(ctx *gin.Context, id int32) *models.ResponseError {
	return cs.userMediaRepository.DeleteMedia(ctx, id)
}

// Validate
func validateMedia(mediaParams *dbcontext.CreateMediaParams) *models.ResponseError {
	if mediaParams.UsmeEntityID == 0 {
		return &models.ResponseError{
			Message: "Invalid Media ",
			Status:  http.StatusBadRequest,
		}
	}

	if mediaParams.UsmeEntityID == 0 {
		return &models.ResponseError{
			Message: "Required Media ",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}
