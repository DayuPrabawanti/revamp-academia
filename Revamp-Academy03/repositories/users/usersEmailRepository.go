package users

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models/usersModel"
	"codeid.revampacademy/repositories/users/dbContext"
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


func (cr UserEmailRepository) GetListUsersEmail(ctx *gin.Context) ([]*usersModel.UsersUsersEmail, *usersModel.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	usersEmail, err := store.ListEmail(ctx)

	listUsersEmail := make([]*usersModel.UsersUsersEmail, 0)

	for _, v := range usersEmail {
		userEmail := &usersModel.UsersUsersEmail{
			PmailEntityID: v.PmailEntityID,
			PmailID: v.PmailID,
			PmailAddress: v.PmailAddress,
			PmailModifiedDate: v.PmailModifiedDate,
		}
		listUsersEmail = append(listUsersEmail, userEmail)
	}

	if err != nil {
		return nil, &usersModel.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listUsersEmail, nil
}


func (cr UserEmailRepository) GetEmail(ctx *gin.Context, id int32) (*usersModel.UsersUsersEmail, *usersModel.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	userEmail, err := store.GetEmail(ctx, int32(id))

	if err != nil {
		return nil, &usersModel.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &userEmail, nil
}

func (cr UserEmailRepository) CreateEmail(ctx *gin.Context, emailParams *dbContext.CreateEmailParams) (*usersModel.UsersUsersEmail, *usersModel.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	email, err := store.CreateEmail(ctx, *emailParams)

	if err != nil {
		return nil, &usersModel.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return email, nil
}

func (cr UserEmailRepository) UpdateEmail(ctx *gin.Context, userEmailParams *dbContext.CreateEmailParams) *usersModel.ResponseError {

	store := dbContext.New(cr.dbHandler)
	err := store.UpdateEmail(ctx, *userEmailParams)

	if err != nil {
		return &usersModel.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &usersModel.ResponseError{
		Message: "data has been update",
		Status:  http.StatusOK,
	}
}

func (cr UserEmailRepository) DeleteEmail(ctx *gin.Context, id int32) *usersModel.ResponseError {

	store := dbContext.New(cr.dbHandler)
	err := store.DeleteEmail(ctx, int32(id))

	if err != nil {
		return &usersModel.ResponseError{
			Message: "error when delete",
			Status:  http.StatusInternalServerError,
		}
	}
	return &usersModel.ResponseError{
		Message: "data has been deleted",
		Status:  http.StatusOK,
	}
}