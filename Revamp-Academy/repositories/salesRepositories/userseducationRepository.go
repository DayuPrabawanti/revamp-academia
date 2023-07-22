package salesRepositories

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/salesRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type EducationRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewEducationRepository(dbHandler *sql.DB) *EducationRepository {
	return &EducationRepository{
		dbHandler: dbHandler,
	}
}

func (cr EducationRepository) CreateEducation(ctx *gin.Context, educationParams *dbContext.CreateEducationParams) (*models.UsersUsersEducation, *models.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	education, err := store.CreateEducation(ctx, *educationParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return education, nil
}

func (cr EducationRepository) GetEducation(ctx *gin.Context, id int32) (*models.UsersUsersEducation, *models.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	education, err := store.GetEducation(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &education, nil
}
