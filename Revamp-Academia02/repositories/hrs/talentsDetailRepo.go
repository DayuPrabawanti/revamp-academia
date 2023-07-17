package hrs

import (
	"database/sql"
	"net/http"

	"codeid.revamptwo/models/hrsMdl"
	"codeid.revamptwo/repositories/hrs/dbContext"
	"github.com/gin-gonic/gin"
)

type TalentsDetailMockupRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewTalentDetailMockupRepository(dbHandler *sql.DB) *TalentsDetailMockupRepository {
	return &TalentsDetailMockupRepository{
		dbHandler: dbHandler,
	}
}

func (cr TalentsDetailMockupRepository) GetListTalentDetailMockup(ctx *gin.Context) ([]*hrsMdl.TalentsDetailMockup, *hrsMdl.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	talentDetail, err := store.ListTalentsDetail(ctx)

	listTalentDetail := make([]*hrsMdl.TalentsDetailMockup, 0)

	for _, v := range talentDetail {
		talents := &hrsMdl.TalentsDetailMockup{
			MasterJobRole:      v.MasterJobRole,
			HrEmployee:         v.HrEmployee,
			UsersUser:          v.UsersUser,
			BootcampBatch:      v.BootcampBatch,
			JobhireTalentApply: v.JobhireTalentApply,
		}
		listTalentDetail = append(listTalentDetail, talents)
	}

	if err != nil {
		return nil, &hrsMdl.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listTalentDetail, nil
}
