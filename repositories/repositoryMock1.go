package repositories

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	dbcontext "codeid.revampacademy/repositories/dbContext"
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

func (rm RepositoryMock) GetMockup(ctx *gin.Context, nama string) (*dbcontext.CreateprogramEntityParams, *models.ResponseError) {

	store := dbcontext.New(rm.dbHandler)
	mockup, err := store.GetProgramEntity(ctx, nama)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &mockup, nil
}

func (rm RepositoryMock) GetMockupId(ctx *gin.Context, id int64) (*models.CurriculumProgramEntity, *models.ResponseError) {

	store := dbcontext.New(rm.dbHandler)
	mockup, err := store.GetProgramEntityId(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &mockup, nil
}
