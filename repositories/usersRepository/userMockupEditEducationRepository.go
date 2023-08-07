package usersRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/usersRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type EditEducationRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewEditEducationRepository(dbHandler *sql.DB) *EditEducationRepository {
	return &EditEducationRepository{
		dbHandler: dbHandler,
	}
}

func (cr EditEducationRepository) AddEducation(ctx *gin.Context, userEducationParams *dbContext.AddEducationParams, id int32) *models.ResponseError {

	store := dbContext.New(cr.dbHandler)

	err := store.AddEducation(ctx, *userEducationParams, id)

	if err != nil {
		return &models.ResponseError{
			Message: "error when Insert",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been Inserted",
		Status:  http.StatusOK,
	}
}