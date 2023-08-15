package bootcampRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/bootcampRepository/dbContext"
	"github.com/gin-gonic/gin"
)

// PROGRAM APPLY PROGRESS
type ProgramApplyProgressRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewProgramApplyProgressRepository(dbHandler *sql.DB) *ProgramApplyProgressRepository {
	return &ProgramApplyProgressRepository{
		dbHandler: dbHandler,
	}
}

func (papr ProgramApplyProgressRepository) GetListProgramApplyProgress(ctx *gin.Context) ([]*models.BootcampProgramApplyProgress, *models.ResponseError) {

	store := dbContext.New(papr.dbHandler)
	progApplyProgresses, err := store.ListProgramApplyProgresses(ctx)

	listProgramApplyProgresses := make([]*models.BootcampProgramApplyProgress, 0)

	for _, v := range progApplyProgresses {
		programApplyProgress := &models.BootcampProgramApplyProgress{
			ParogID:           v.ParogID,
			ParogUserEntityID: v.ParogUserEntityID,
			ParogProgEntityID: v.ParogProgEntityID,
			ParogActionDate:   v.ParogActionDate,
			ParogModifiedDate: v.ParogModifiedDate,
			ParogComment:      v.ParogComment,
			ParogProgressName: v.ParogProgressName,
			ParogEmpEntityID:  v.ParogEmpEntityID,
			ParogStatus:       v.ParogStatus,
		}
		listProgramApplyProgresses = append(listProgramApplyProgresses, programApplyProgress)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listProgramApplyProgresses, nil
}
