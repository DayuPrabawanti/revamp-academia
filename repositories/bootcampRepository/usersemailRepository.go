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
