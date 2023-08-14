package jobhireRepositories

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/jobhireRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type CategoryRepo struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
	dbQueries   dbContext.Queries
}

func NewCategoryRepo(dbHandler *sql.DB) *CategoryRepo {
	return &CategoryRepo{
		dbHandler: dbHandler,
		dbQueries: *dbContext.New(dbHandler),
	}
}

func (cr CategoryRepo) GetListCategoryJob(ctx *gin.Context) ([]*models.JobhireJobCategory, *models.ResponseError) {
	market := dbContext.New(cr.dbHandler)
	categories, err := market.ListCategories(ctx)

	listCategories := make([]*models.JobhireJobCategory, 0)

	for _, v := range categories {
		category := models.JobhireJobCategory{
			JocaID:           v.JocaID,
			JocaName:         v.JocaName,
			JocaModifiedDate: v.JocaModifiedDate,
		}
		listCategories = append(listCategories, &category)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listCategories, nil
}
