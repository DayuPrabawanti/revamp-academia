package hrRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/hrRepository/dbContext"
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

func (tdmr TalentsDetailMockupRepository) GetListTalentDetailMockup(ctx *gin.Context) ([]*models.TalentsDetailMockup, *models.ResponseError) {

	store := dbContext.New(tdmr.dbHandler)
	talentDetail, err := store.ListTalentsDetail(ctx)

	listTalentDetail := make([]*models.TalentsDetailMockup, 0)

	for _, v := range talentDetail {
		talents := &models.TalentsDetailMockup{
			UsersUser:                      v.UsersUser,
			BootcampBatch:                  v.BootcampBatch,
			BootcampBatchTrainee:           v.BootcampBatchTrainee,
			BootcampBatchTraineeEvaluation: v.BootcampBatchTraineeEvaluation,
			UsersUsersEmail:                v.UsersUsersEmail,
			UsersUsersPhone:                v.UsersUsersPhone,
			CurriculumProgramEntity:        v.CurriculumProgramEntity,
			JobhireJobPost:                 v.JobhireJobPost,
			JobhireClient:                  v.JobhireClient,
			HrEmployeeClientContract:       v.HrEmployeeClientContract,
		}
		listTalentDetail = append(listTalentDetail, talents)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listTalentDetail, nil
}

func (tdmr TalentsDetailMockupRepository) GetTalentDetail(ctx *gin.Context, id int64) (*models.TalentsDetailMockup, *models.ResponseError) {

	store := dbContext.New(tdmr.dbHandler)
	talentDetails, err := store.GetTalentDetail(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &talentDetails, nil
}

func (tdmr TalentsDetailMockupRepository) SearchTalentDetail(ctx *gin.Context, clitName string) ([]models.TalentDetailSearchUpdate, *models.ResponseError) {
	store := dbContext.New(tdmr.dbHandler)
	talents, err := store.SearchTalentDetail(ctx, clitName)
	if err != nil {
		return nil, &models.ResponseError{
			Message: "Failed to search talents",
			Status:  http.StatusInternalServerError,
		}
	}

	return talents, nil
}

func (br TalentsDetailMockupRepository) UpdateSwitch(ctx *gin.Context, switchParams *dbContext.UpdateSwitchParams) *models.ResponseError {

	store := dbContext.New(br.dbHandler)
	err := store.UpdateSwitch(ctx, *switchParams)

	if err != nil {
		return &models.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been update",
		Status:  http.StatusOK,
	}
}
