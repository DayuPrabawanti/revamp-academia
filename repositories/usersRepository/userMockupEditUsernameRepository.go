package usersRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/usersRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type EditUsernameRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewEditUsernameRepository(dbHandler *sql.DB) *EditUsernameRepository {
	return &EditUsernameRepository{
		dbHandler: dbHandler,
	}
}

func (cr EditUsernameRepository) GetUsername(ctx *gin.Context, id int32) (*dbContext.GetUsernameParams, *models.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	username, err := store.GetUsername(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &username, nil
}

func (cr EditUsernameRepository) EditUsername(ctx *gin.Context, usernameParams *dbContext.EditUsernameParams) *models.ResponseError {

	store := dbContext.New(cr.dbHandler)
	err := store.EditUsername(ctx, *usernameParams)

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