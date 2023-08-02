package usersService

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/usersRepository"
	"codeid.revampacademy/repositories/usersRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type SignUpEmployeeService struct {
	signupEmployeeRepository *usersRepository.SignUpEmployeeRepository
}

func NewSignUpEmployeeService(signUpEmployeeRepository *usersRepository.SignUpEmployeeRepository) *SignUpEmployeeService {
	return &SignUpEmployeeService{
		signupEmployeeRepository: signUpEmployeeRepository,
	}
}

func (cs *SignUpEmployeeService) SignUpEmplloyee(ctx *gin.Context, userParams *dbContext.CreateUsersParams, emailParams *dbContext.CreateEmailParams, phoneParams *dbContext.CreatePhonesParams) (*models.SignUpUser, *models.ResponseError) {
	return cs.signupEmployeeRepository.CreateSignUpEmployee(ctx, userParams, emailParams, phoneParams)
}