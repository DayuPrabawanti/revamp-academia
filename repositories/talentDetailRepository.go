package repositories

// import (
// 	"database/sql"
// 	"net/http"

// 	"codeid.revampacademy/models"
// 	"codeid.revampacademy/repositories/dbContext"
// 	"github.com/gin-gonic/gin"
// )

// type TalentsDetailMockupRepository struct {
// 	dbHandler   *sql.DB
// 	transaction *sql.Tx
// }

// func NewTalentDetailMockupRepository(dbHandler *sql.DB) *TalentsDetailMockupRepository {
// 	return &TalentsDetailMockupRepository{
// 		dbHandler: dbHandler,
// 	}
// }

// func (tdmr TalentsDetailMockupRepository) GetListTalentDetailMockup(ctx *gin.Context) ([]*models.TalentsDetailMockup, *models.ResponseError) {

// 	store := dbContext.New(tdmr.dbHandler)
// 	talentDetail, err := store.ListTalentsDetail(ctx)

// 	listTalentDetail := make([]*models.TalentsDetailMockup, 0)

// 	for _, v := range talentDetail {
// 		talents := &models.TalentsDetailMockup{
// 			MasterJobRole:      v.MasterJobRole,
// 			HrEmployee:         v.HrEmployee,
// 			UsersUser:          v.UsersUser,
// 			BootcampBatch:      v.BootcampBatch,
// 			JobhireTalentApply: v.JobhireTalentApply,
// 		}
// 		listTalentDetail = append(listTalentDetail, talents)
// 	}

// 	if err != nil {
// 		return nil, &models.ResponseError{
// 			Message: err.Error(),
// 			Status:  http.StatusInternalServerError,
// 		}
// 	}

// 	return listTalentDetail, nil
// }
