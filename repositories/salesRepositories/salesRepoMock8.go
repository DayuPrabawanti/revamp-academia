package salesrepositories

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	dbcontext "codeid.revampacademy/repositories/salesRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type RepoMockup8 struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewRepoShoppingCart3(dbHandler *sql.DB) *RepoMockup8 {
	return &RepoMockup8{
		dbHandler: dbHandler,
	}
}

func (rm RepoMockup8) GetIdSummaryOrderMock8Repo(ctx *gin.Context, poNo string) (*dbcontext.GetSummaryOrderMock8, *models.ResponseError) {

	store := dbcontext.New(rm.dbHandler)
	mockup, err := store.GetIdSummaryOrderMock8(ctx, string(poNo))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &mockup, nil
}
