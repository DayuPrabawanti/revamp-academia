package bootcampRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/bootcampRepository/dbContext"
	"github.com/gin-gonic/gin"
)

// PROGRAM APPLY
type ProgramApplyRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewProgramApplyRepository(dbHandler *sql.DB) *ProgramApplyRepository {
	return &ProgramApplyRepository{
		dbHandler: dbHandler,
	}
}

func (par ProgramApplyRepository) GetListProgramApply(ctx *gin.Context) ([]*models.BootcampProgramApply, *models.ResponseError) {

	store := dbContext.New(par.dbHandler)
	progApplies, err := store.ListProgramApplies(ctx)

	listProgramApplies := make([]*models.BootcampProgramApply, 0)

	for _, v := range progApplies {
		programApply := &models.BootcampProgramApply{
			PrapUserEntityID: v.PrapUserEntityID,
			PrapProgEntityID: v.PrapProgEntityID,
			PrapTestScore:    v.PrapTestScore,
			PrapGpa:          v.PrapGpa,
			PrapIqTest:       v.PrapIqTest,
			PrapReview:       v.PrapReview,
			PrapModifiedDate: v.PrapModifiedDate,
			PrapStatus:       v.PrapStatus,
		}
		listProgramApplies = append(listProgramApplies, programApply)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listProgramApplies, nil
}
