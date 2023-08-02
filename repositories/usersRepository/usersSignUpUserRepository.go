package usersRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/usersRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type SignUpRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewSignUpRepository(dbHandler *sql.DB) *SignUpRepository {
	return &SignUpRepository{
		dbHandler: dbHandler,
	}
}


func (cr SignUpRepository) CreateSignUpUser(ctx *gin.Context, userParams *dbContext.CreateUsersParams, emailParams *dbContext.CreateEmailParams, phoneParams *dbContext.CreatePhonesParams) (*models.SignUpUser, *models.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	user, err := store.CreateUsers(ctx, *userParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}

	createdEmail, err := store.CreateEmail(ctx, *emailParams)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}

	createdPhone, err := store.CreatePhones(ctx, *phoneParams)
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