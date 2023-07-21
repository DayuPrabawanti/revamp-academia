package salesRepositories

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/dbContext/salesContext"
	"github.com/gin-gonic/gin"
)

type RepositoryMock struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewRepositoryMock(dbHandler *sql.DB) *RepositoryMock {
	return &RepositoryMock{
		dbHandler: dbHandler,
	}
}

func (rm RepositoryMock) GetMockup(ctx *gin.Context, nama string) (*salesContext.CreateprogramEntityParams, *models.ResponseError) {

	store := salesContext.New(rm.dbHandler)
	mockup, err := store.GetProgramEntity(ctx, nama)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &mockup, nil
}
