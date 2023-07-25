package hrRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/hrRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type TalentsMockupRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewTalentMockupRepository(dbHandler *sql.DB) *TalentsMockupRepository {
	return &TalentsMockupRepository{
		dbHandler: dbHandler,
	}
}

func (tmr TalentsMockupRepository) GetListTalentMockup(ctx *gin.Context) ([]*models.TalentsMockup, *models.ResponseError) {

	store := dbContext.New(tmr.dbHandler)
	talent, err := store.ListTalents(ctx)

	listTalent := make([]*models.TalentsMockup, 0)

	for _, v := range talent {
		talents := &models.TalentsMockup{
			UsersUser:                      v.UsersUser,
			BootcampBatch:                  v.BootcampBatch,
			BootcampBatchTraineeEvaluation: v.BootcampBatchTraineeEvaluation,
			CurriculumProgramEntity:        v.CurriculumProgramEntity,
		}
		listTalent = append(listTalent, talents)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listTalent, nil
}

func (tl TalentsMockupRepository) SearchTalent(ctx *gin.Context, userName, skillName, batchStatus string) ([]models.TalentsMockup, *models.ResponseError) {
	// Perform validation, if needed, for batchName and status
	// If validation fails, return appropriate response error

	store := dbContext.New(tl.dbHandler)
	talents, err := store.SearchTalent(ctx, userName, skillName, batchStatus)
	if err != nil {
		return nil, &models.ResponseError{
			Message: "Failed to search talents",
			Status:  http.StatusInternalServerError,
		}
	}

	return talents, nil
}

func (tl TalentsMockupRepository) PagingTalent(ctx *gin.Context, offset, pageSize int) ([]models.TalentsMockup, *models.ResponseError) {

	store := dbContext.New(tl.dbHandler)
	talents, err := store.PagingTalent(ctx, offset, pageSize)
	if err != nil {
		return nil, &models.ResponseError{
			Message: "Failed to fetch talents",
			Status:  http.StatusInternalServerError,
		}
	}

	return talents, nil
}

func (br TalentsMockupRepository) GetBatch(ctx *gin.Context, id int64) (*models.BootcampBatch, *models.ResponseError) {

	store := dbContext.New(br.dbHandler)
	batch, err := store.GetBatch(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &batch, nil
}

func (br TalentsMockupRepository) UpdateBatch(ctx *gin.Context, batchParams *dbContext.UpdateBatchParams) *models.ResponseError {

	store := dbContext.New(br.dbHandler)
	err := store.UpdateBatch(ctx, *batchParams)

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
