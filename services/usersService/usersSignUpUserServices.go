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

func (cs *SignUpService) SignUpUser(ctx *gin.Context, userParams *dbContext.CreateUsersParams, emailParams *dbContext.CreateEmailParams, phoneParams *dbContext.CreatePhonesParams) (*models.SignUpUser, *models.ResponseError) {
	return cs.signupRepository.CreateSignUpUser(ctx, userParams, emailParams, phoneParams)
}

func validateSignUp(signupParams *dbContext.SignUpUserParams) *models.ResponseError {
	if signupParams.User.UserName.Valid == false {
		return &models.ResponseError{
			Message: "Username tidak Boleh Kosong",
			Status:  http.StatusBadRequest,
		}
	}

	if signupParams.Email.PmailAddress.Valid== false {
		return &models.ResponseError{
			Message: "Email tidak Boleh Kosong",
			Status:  http.StatusBadRequest,
		}
	}

	// if signupParams.Phone.UspoNumber == "" {
	// 	return &models.ResponseError{
	// 		Message: "Mohon isi nomor telepon",
	// 		Status:  http.StatusBadRequest,
	// 	}
	// }

	// if signupParams.User.UserName == "" {
	// 	return &models.ResponseError{
	// 		Message: "Nama Pengguna harus diisi",
	// 		Status:  http.StatusBadRequest,
	// 	}
	// }

	return nil
}