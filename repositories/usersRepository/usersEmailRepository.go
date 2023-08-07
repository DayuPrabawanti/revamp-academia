package usersRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/usersRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type UserEmailRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewUserEmailRepository(dbHandler *sql.DB) *UserEmailRepository {
	return &UserEmailRepository{
		dbHandler: dbHandler,
	}
}


func (cr UserEmailRepository) GetListUsersEmail(ctx *gin.Context) ([]*models.UsersUsersEmail, *models.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	usersEmail, err := store.ListEmail(ctx)

	listUsersEmail := make([]*models.UsersUsersEmail, 0)

	for _, v := range usersEmail {
		userEmail := &models.UsersUsersEmail{
			PmailEntityID: v.PmailEntityID,
			PmailID: v.PmailID,
			PmailAddress: v.PmailAddress,
			PmailModifiedDate: v.PmailModifiedDate,
		}
		listUsersEmail = append(listUsersEmail, userEmail)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listUsersEmail, nil
}


func (cr UserEmailRepository) GetEmail(ctx *gin.Context, id int32) ([]*dbContext.Email, *models.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	userEmail, err := store.GetEmail(ctx, int32(id))

	listEmail := make([]*dbContext.Email, 0)

	for _, v := range userEmail {
		user := &dbContext.Email{
			PmailAddress: v.PmailAddress,
		}
		listEmail = append(listEmail, user)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listEmail, nil
}

func (cr UserEmailRepository) CreateEmail(ctx *gin.Context, emailParams *dbContext.CreateEmailParams) (*models.UsersUsersEmail, *models.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	email, err := store.CreateEmail(ctx, *emailParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return email, nil
}

func (cr UserEmailRepository) UpdateEmail(ctx *gin.Context, userEmailParams *dbContext.CreateEmailParams) *models.ResponseError {

	store := dbContext.New(cr.dbHandler)
	err := store.UpdateEmail(ctx, *userEmailParams)

	if err != nil {
		return &models.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been update",
		Status:  http.StatusOK,
	}
}

func (cr UserEmailRepository) DeleteEmail(ctx *gin.Context, id int32) *models.ResponseError {

	store := dbContext.New(cr.dbHandler)
	err := store.DeleteEmail(ctx, int32(id))

	if err != nil {
		return &models.ResponseError{
			Message: "error when delete",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been deleted",
		Status:  http.StatusOK,
	}
}
