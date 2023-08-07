package usersRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/usersRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type EditPasswordRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewEditPasswordRepository(dbHandler *sql.DB) *EditPasswordRepository {
	return &EditPasswordRepository{
		dbHandler: dbHandler,
	}
}

func (cr EditPasswordRepository) GetPassword(ctx *gin.Context, id int32) (*dbContext.GetUsersPasswordParams, *models.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	userPassword, err := store.GetPassword(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &userPassword, nil
}

func (cr EditPasswordRepository) EditPassword(ctx *gin.Context, id int32, userPasswordParams *dbContext.EditUsersPasswordParams) *models.ResponseError {

	store := dbContext.New(cr.dbHandler)
	err := store.UpdatePassword(ctx, int32(id), *userPasswordParams)

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