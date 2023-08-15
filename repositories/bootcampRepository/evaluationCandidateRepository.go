package bootcampRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/bootcampRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type EvaluationCandidateRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewEvaluationCandidateRepository(dbHandler *sql.DB) *EvaluationCandidateRepository {
	return &EvaluationCandidateRepository{
		dbHandler: dbHandler,
	}
}

func (ecr EvaluationCandidateRepository) GetListEvaluationCandidate(ctx *gin.Context) ([]*models.EvaluationCandidateMockup, *models.ResponseError) {

	store := dbContext.New(ecr.dbHandler)
	evaluationCandidates, err := store.ListEvaluationCandidate(ctx)

	listEvaluationCandidates := make([]*models.EvaluationCandidateMockup, 0)

	for _, v := range evaluationCandidates {
		evaluationCandidate := &models.EvaluationCandidateMockup{
			BootcampBatch:                  v.BootcampBatch,
			BootcampBatchTraineeEvaluation: v.BootcampBatchTraineeEvaluation,
			CurriculumProgramEntity:        v.CurriculumProgramEntity,
			UsersUser:                      v.UsersUser,
			UsersUsersEducation:            v.UsersUsersEducation,
		}
		listEvaluationCandidates = append(listEvaluationCandidates, evaluationCandidate)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listEvaluationCandidates, nil
}
