package usersRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/usersRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type EditExperienceRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewEditExperienceRepository(dbHandler *sql.DB) *EditExperienceRepository {
	return &EditExperienceRepository{
		dbHandler: dbHandler,
	}
}

func (cr EditExperienceRepository) AddExperience(ctx *gin.Context, userExperienceParams *dbContext.AddExperienceParams, id int32) *models.ResponseError {

	store := dbContext.New(cr.dbHandler)

	err := store.AddExperience(ctx, *userExperienceParams, id)

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