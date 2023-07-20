package usersrepository

import (
	"database/sql"
	"net/http"

	model "codeid.revampacademy/models/usersModel"
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

// GetList User Email
func (cr UserEmailRepository) GetListUserEmail(ctx *gin.Context) ([]*model.UsersUsersEmail, *model.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	usersEmail, err := store.ListEmail(ctx)

	listUsersEmail := make([]*model.UsersUsersEmail, 0)

	for _, v := range usersEmail {
		userEmail := &model.UsersUsersEmail{
			PmailEntityID:     v.PmailEntityID,
			PmailID:           v.PmailID,
			PmailAddress:      v.PmailAddress,
			PmailModifiedDate: v.PmailModifiedDate,
		}
		listUsersEmail = append(listUsersEmail, userEmail)
	}

	if err != nil {
		return nil, &model.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listUsersEmail, nil
}

// Get User Email
func (cr UserEmailRepository) GetUserEmail(ctx *gin.Context, id int32) (*model.UsersUsersEmail, *model.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	userEmail, err := store.GetEmail(ctx, int32(id))

	if err != nil {
		return nil, &model.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &userEmail, nil
}

// Create User Email
func (cr UserEmailRepository) CreateUserEmail(ctx *gin.Context, emailParams *dbContext.CreateEmailParams) (*model.UsersUsersEmail, *model.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	email, err := store.CreateEmail(ctx, *emailParams)

	if err != nil {
		return nil, &model.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return email, nil
}

// Update Table
func (cr UserEmailRepository) UpdateEmail(ctx *gin.Context, userEmailParams *dbContext.CreateEmailParams) *model.ResponseError {

	store := dbContext.New(cr.dbHandler)
	err := store.UpdateEmail(ctx, *userEmailParams)

	if err != nil {
		return &model.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &model.ResponseError{
		Message: "data has been update",
		Status:  http.StatusOK,
	}
}

// Delete Table
func (cr UserEmailRepository) DeleteEmail(ctx *gin.Context, id int32) *model.ResponseError {

	store := dbContext.New(cr.dbHandler)
	err := store.DeleteEmail(ctx, int32(id))

	if err != nil {
		return &model.ResponseError{
			Message: "error when delete",
			Status:  http.StatusInternalServerError,
		}
	}
	return &model.ResponseError{
		Message: "data has been deleted",
		Status:  http.StatusOK,
	}
}
