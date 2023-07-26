package salesRepositories

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	dbcontext "codeid.revampacademy/repositories/salesRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type RepoMockup4 struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewMockupApplyRepo4(dbHandler *sql.DB) *RepoMockup4 {
	return &RepoMockup4{
		dbHandler: dbHandler,
	}
}

func (rm RepoMockup4) ListMock4Group(ctx *gin.Context) ([]*dbcontext.CreateMergeMock4, *models.ResponseError) {

	store := dbcontext.New(rm.dbHandler)
	bootcampGrup, err := store.ListMock4Group(ctx)

	listMock4Grup := make([]*dbcontext.CreateMergeMock4, 0)

	for _, v := range bootcampGrup {
		sales := &dbcontext.CreateMergeMock4{
			Createprogram_entityParams:       v.Createprogram_entityParams,
			CreateProgramApplyParam:          v.CreateProgramApplyParam,
			CreateProgramApplyProgressParams: v.CreateProgramApplyProgressParams,
		}
		listMock4Grup = append(listMock4Grup, sales)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listMock4Grup, nil
}

func (rm RepoMockup4) GetMock4Group(ctx *gin.Context, id int32) (*models.MergeMock4, *models.ResponseError) {

	store := dbcontext.New(rm.dbHandler)
	mockup, err := store.GetMock4Group(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &mockup, nil
}
