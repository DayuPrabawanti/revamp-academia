package bootcampRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/bootcampRepository/dbContext"
	"github.com/gin-gonic/gin"
)

// INSTRUCTOR PROGRAMS
type InstructorProgramRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewInstructorProgramRepository(dbHandler *sql.DB) *InstructorProgramRepository {
	return &InstructorProgramRepository{
		dbHandler: dbHandler,
	}
}

func (ipr InstructorProgramRepository) GetListInstructorPrograms(ctx *gin.Context) ([]*models.BootcampInstructorProgram, *models.ResponseError) {

	store := dbContext.New(ipr.dbHandler)
	instructorProgs, err := store.ListInstructorPrograms(ctx)

	listInstructorPrograms := make([]*models.BootcampInstructorProgram, 0)

	for _, v := range instructorProgs {
		instructorProgram := &models.BootcampInstructorProgram{
			BatchID:           v.BatchID,
			InproEntityID:     v.InproEntityID,
			InproEmpEntityID:  v.InproEmpEntityID,
			InproModifiedDate: v.InproModifiedDate,
		}
		listInstructorPrograms = append(listInstructorPrograms, instructorProgram)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listInstructorPrograms, nil
}
