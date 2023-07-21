package usersService

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/usersRepository"
	"codeid.revampacademy/repositories/usersRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type SignUpService struct {
	signupRepository *usersRepository.SignUpRepository
}

func NewSignUpService(signUpRepository *usersRepository.SignUpRepository) *SignUpService {
	return &SignUpService{
		signupRepository: signUpRepository,
	}
}

func (cs *SignUpService) SignUpUser(ctx *gin.Context, signupParams *dbContext.SignUpUserParams) (*models.SignUpUser, *models.ResponseError) {
	responseErr := validateSignUp(signupParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return cs.signupRepository.CreateSignUp(ctx, signupParams)
}

func validateSignUp(signupParams *dbContext.SignUpUserParams) *models.ResponseError {
	if signupParams.User.UserEntityID != 0 {
		return &models.ResponseError{
			Message: "ID Pengguna tidak valid",
			Status:  http.StatusBadRequest,
		}
	}

	// if signupParams.User.UserName == "" {
	// 	return &models.ResponseError{
	// 		Message: "Nama Pengguna harus diisi",
	// 		Status:  http.StatusBadRequest,
	// 	}
	// }

	return nil
}