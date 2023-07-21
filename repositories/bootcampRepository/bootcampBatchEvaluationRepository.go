package bootcampRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/bootcampRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type BootcampBatchEvaluationRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewBootcampBatchEvaluationRepository(dbHandler *sql.DB) *BootcampBatchEvaluationRepository {
	return &BootcampBatchEvaluationRepository{
		dbHandler: dbHandler,
	}
}

func (ber BootcampBatchEvaluationRepository) GetListBootcampBatchEvaluation(ctx *gin.Context) ([]*models.BootcampBatchEvaluationMockup, *models.ResponseError) {

	store := dbContext.New(ber.dbHandler)
	bootcampBatchEvs, err := store.ListBootcampBatchEvaluation(ctx)

	listBootcampBatchEvaluations := make([]*models.BootcampBatchEvaluationMockup, 0)

	for _, v := range bootcampBatchEvs {
		bootcampBatchEv := &models.BootcampBatchEvaluationMockup{
			BootcampBatch:                  v.BootcampBatch,
			BootcampBatchTraineeEvaluation: v.BootcampBatchTraineeEvaluation,
			CurriculumProgramEntity:        v.CurriculumProgramEntity,
			UsersUser:                      v.UsersUser,
		}
		listBootcampBatchEvaluations = append(listBootcampBatchEvaluations, bootcampBatchEv)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listBootcampBatchEvaluations, nil
}
