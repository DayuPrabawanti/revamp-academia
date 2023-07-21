package bootcampRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/bootcampRepository/dbContext"
	"github.com/gin-gonic/gin"
)

// BATCH TRAINEE EVALUATION
type BatchTraineeEvaluationRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewBatchTraineeEvaluationRepository(dbHandler *sql.DB) *BatchTraineeEvaluationRepository {
	return &BatchTraineeEvaluationRepository{
		dbHandler: dbHandler,
	}
}

func (bter BatchTraineeEvaluationRepository) GetListBatchTraineeEvaluation(ctx *gin.Context) ([]*models.BootcampBatchTraineeEvaluation, *models.ResponseError) {

	store := dbContext.New(bter.dbHandler)
	batchTraineeEvs, err := store.ListBatchTraineeEvaluations(ctx)

	listBatchTraineeEvaluations := make([]*models.BootcampBatchTraineeEvaluation, 0)

	for _, v := range batchTraineeEvs {
		batchTraineeEv := &models.BootcampBatchTraineeEvaluation{
			BtevID:              v.BtevID,
			BtevType:            v.BtevType,
			BtevHeader:          v.BtevHeader,
			BtevSection:         v.BtevSection,
			BtevSkill:           v.BtevSkill,
			BtevWeek:            v.BtevWeek,
			BtevSkor:            v.BtevSkor,
			BtevNote:            v.BtevNote,
			BtevModifiedDate:    v.BtevModifiedDate,
			BtevBatchID:         v.BtevBatchID,
			BtevTraineeEntityID: v.BtevTraineeEntityID,
		}
		listBatchTraineeEvaluations = append(listBatchTraineeEvaluations, batchTraineeEv)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listBatchTraineeEvaluations, nil
}
