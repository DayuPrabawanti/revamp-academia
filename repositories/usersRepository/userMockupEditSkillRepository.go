package usersRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/usersRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type EditSkillRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewEditSkillRepository(dbHandler *sql.DB) *EditSkillRepository {
	return &EditSkillRepository{
		dbHandler: dbHandler,
	}
}

func (cr EditSkillRepository) AddSkill(ctx *gin.Context, userSkillParams *dbContext.AddSkillParams, id int32) *models.ResponseError {

	store := dbContext.New(cr.dbHandler)
	err := store.AddSkill(ctx, *userSkillParams, id)

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