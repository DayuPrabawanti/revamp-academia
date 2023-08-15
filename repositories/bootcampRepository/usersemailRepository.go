package bootcampRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/bootcampRepository/dbContext"
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
			PmailEntityID:     v.PmailEntityID,
			PmailID:           v.PmailID,
			PmailAddress:      v.PmailAddress,
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

func (cr UserEmailRepository) GetUSerEmail(ctx *gin.Context, id int64) (*models.UsersUsersEmail, *models.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	pmailEntityId, err := store.GetUSerEmail(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &pmailEntityId, nil
}
