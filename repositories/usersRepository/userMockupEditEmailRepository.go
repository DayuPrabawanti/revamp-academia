package usersRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/usersRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type EditEmailRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewEditEmailRepository(dbHandler *sql.DB) *EditEmailRepository {
	return &EditEmailRepository{
		dbHandler: dbHandler,
	}
}

// func (cr EditEmailRepository) AddEmail(ctx *gin.Context, emailParams *dbContext.AddEmailParams, id int32) (*dbContext.AddEmailParams, *models.ResponseError) {

// 	store := dbContext.New(cr.dbHandler)
// 	email, err := store.AddEmail(ctx, *emailParams, )

// 	if err != nil {
// 		return nil, &models.ResponseError{
// 			Message: err.Message,
// 			Status:  http.StatusInternalServerError,
// 		}
// 	}
// 	return email, nil
// }

func (cr EditEmailRepository) AddEmail(ctx *gin.Context, userEmailParams *dbContext.AddEmailParams, id int32) *models.ResponseError {

	store := dbContext.New(cr.dbHandler)
	err := store.AddEmail(ctx, *userEmailParams, id)

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