package usersRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/usersRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type EditPhoneRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewEditPhoneRepository(dbHandler *sql.DB) *EditPhoneRepository {
	return &EditPhoneRepository{
		dbHandler: dbHandler,
	}
}

func (cr EditPhoneRepository) AddPhone(ctx *gin.Context, userPhoneParams *dbContext.AddPhonesParams, id int32) *models.ResponseError {

	store := dbContext.New(cr.dbHandler)
	err := store.AddPhone(ctx, *userPhoneParams, id)

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