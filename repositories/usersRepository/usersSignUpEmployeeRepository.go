package usersRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/usersRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type SignUpEmployeeRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewSignUpEmployeeRepository(dbHandler *sql.DB) *SignUpEmployeeRepository {
	return &SignUpEmployeeRepository{
		dbHandler: dbHandler,
	}
}

func (cr SignUpEmployeeRepository) CreateSignUpEmployee(ctx *gin.Context, userParams *dbContext.CreateUsersParams, emailParams *dbContext.CreateEmailParams, phoneParams *dbContext.CreatePhonesParams) (*models.SignUpUser, *models.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	user, err := store.CreateEmployee(ctx, *userParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}

	createdEmail, err := store.CreateEmailEmployee(ctx, *emailParams)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}

	createdPhone, err := store.CreatePhonesEmployee(ctx, *phoneParams)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	signUpUser := &models.SignUpUser{
		User:  *user,
		Email: *createdEmail,
		Phone: *createdPhone,
	}

	return signUpUser, nil
}