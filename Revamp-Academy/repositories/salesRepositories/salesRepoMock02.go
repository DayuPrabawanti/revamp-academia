package salesRepositories

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	dbcontext "codeid.revampacademy/repositories/salesRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type RepoMockup2 struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewMockupApplyRepo2(dbHandler *sql.DB) *RepoMockup2 {
	return &RepoMockup2{
		dbHandler: dbHandler,
	}
}

func (rm RepoMockup2) ListBootcampGroup(ctx *gin.Context) ([]*dbcontext.CreateMergeMock2, *models.ResponseError) {

	store := dbcontext.New(rm.dbHandler)
	bootcampGrup, err := store.ListBootcampGroup(ctx)

	listBootcampGrup := make([]*dbcontext.CreateMergeMock2, 0)

	for _, v := range bootcampGrup {
		sales := &dbcontext.CreateMergeMock2{
			CreateProgramApplyParams:      v.CreateProgramApplyParams,
			CreateBatchParams:             v.CreateBatchParams,
			CreateInstructorProgramParams: v.CreateInstructorProgramParams,
		}
		listBootcampGrup = append(listBootcampGrup, sales)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listBootcampGrup, nil
}
