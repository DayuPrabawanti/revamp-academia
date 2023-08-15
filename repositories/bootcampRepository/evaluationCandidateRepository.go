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
	dbQueries   dbContext.Queries
}

func NewEvaluationCandidateRepository(dbHandler *sql.DB) *EvaluationCandidateRepository {
	return &EvaluationCandidateRepository{
		dbHandler: dbHandler,
		//add new fields
		dbQueries: *dbContext.New(dbHandler),
	}
}

func (ecr EvaluationCandidateRepository) GetEvaluationCandidate(ctx *gin.Context, id int32) ([]*dbContext.BootcampEvaluationCandidate, *models.ResponseError) {

	// store := dbContext.New(ecr.dbHandler)
	evaluationCandidates, err := ecr.dbQueries.GetEvaluationCandidate(ctx, int32(id))

	evaluationCandidate := make([]*dbContext.BootcampEvaluationCandidate, 0)

	for _, v := range evaluationCandidates {
		evaluation := &dbContext.BootcampEvaluationCandidate{
			UserEntityID:   v.UserEntityID,
			UserFullname:   v.UserFullname,
			UserPhoto:      v.UserPhoto,
			UsduSchool:     v.UsduSchool,
			UsduFieldStudy: v.UsduFieldStudy,
			UsduGrade:      v.UsduGrade,
			ProgTitle:      v.ProgTitle,
			BatchID:        v.BatchID,
			BatchName:      v.BatchName,
			BatchStartDate: v.BatchStartDate,
			BatchEndDate:   v.BatchEndDate,
			BatrStatus:     v.BatrStatus,
			BtevBatchID:    v.BtevBatchID,
			BtevSkor:       v.BtevSkor,
			BtevType:       v.BtevType,
			BtevHeader:     v.BtevHeader,
			BtevSkill:      v.BtevSkill,
		}
		evaluationCandidate = append(evaluationCandidate, evaluation)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return evaluationCandidate, nil
}

func (ecr EvaluationCandidateRepository) CreateEvaluationCandidate(ctx *gin.Context, evaluationCandidateParams *dbContext.CreateEvaluationCandidateParams) (*dbContext.BootcampEvaluationCandidate, *models.ResponseError) {

	// store := dbContext.New(btr.dbHandler)
	evaluationCandidate, err := ecr.dbQueries.CreateEvaluationCandidate(ctx, *evaluationCandidateParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return evaluationCandidate, nil
}

// UpdateBatchTrainee updates an existing batch trainee in the database.
func (ecr EvaluationCandidateRepository) UpdateEvaluationCandidate(ctx *gin.Context, evaluationCandidateParams *dbContext.UpdateEvaluationCandidateParams, userEntityId int32) (*dbContext.BootcampEvaluationCandidate, *models.ResponseError) {
	// store := dbContext.New(btr.dbHandler)
	evaluationCandidate, err := ecr.dbQueries.UpdateEvaluationCandidate(ctx, *evaluationCandidateParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return evaluationCandidate, nil
}

// func (ecr EvaluationCandidateRepository) UpdateEvaluationCandidate(ctx *gin.Context, score, userEntityId int32) (*dbContext.BootcampEvaluationCandidate, error) {
// 	evaluationCandidate, err := ecr.dbQueries.UpdateEvaluationCandidate(ctx, dbContext.UpdateEvaluationCandidateParams{BtevSkor: score}, userEntityId)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return evaluationCandidate, nil
// }
