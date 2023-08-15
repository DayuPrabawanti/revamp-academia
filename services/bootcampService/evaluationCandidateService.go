package bootcampService

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/bootcampRepository"
	"codeid.revampacademy/repositories/bootcampRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type EvaluationCandidateService struct {
	// evaluationCandidateRepository *bootcampRepository.EvaluationCandidateRepository
	repositoryManager *bootcampRepository.RepositoryManager
}

func NewEvaluationCandidateService(repoMgr *bootcampRepository.RepositoryManager) *EvaluationCandidateService {
	return &EvaluationCandidateService{
		repositoryManager: repoMgr,
	}
}

func (ecs EvaluationCandidateService) GetEvaluationCandidate(ctx *gin.Context, id int32) ([]*dbContext.BootcampEvaluationCandidate, *models.ResponseError) {
	return ecs.repositoryManager.EvaluationCandidateRepository.GetEvaluationCandidate(ctx, id)
}

func (ecs EvaluationCandidateService) CreateEvaluationCandidate(ctx *gin.Context, evaluationCandidate *dbContext.CreateEvaluationCandidateParams) (*dbContext.BootcampEvaluationCandidate, *models.ResponseError) {
	responseErr := validateEvaluationCandidate(evaluationCandidate)
	if responseErr != nil {
		return nil, responseErr
	}

	return ecs.repositoryManager.EvaluationCandidateRepository.CreateEvaluationCandidate(ctx, evaluationCandidate)
}

func validateEvaluationCandidate(evaluationCandidateParams *dbContext.CreateEvaluationCandidateParams) *models.ResponseError {
	if evaluationCandidateParams.BtevType == "" {
		return &models.ResponseError{
			Message: "Invalid batch trainee evaluation type",
			Status:  http.StatusBadRequest,
		}
	}

	// menambahkan validasi lain berdasarkan kebutuhan
	if evaluationCandidateParams.BtevSkor <= 0 || evaluationCandidateParams.BtevSkor > 100 { // misalnya, skor harus antara 1 dan 100
		return &models.ResponseError{
			Message: "Invalid score. It should be between 1 and 100.",
			Status:  http.StatusBadRequest,
		}
	}

	return nil
}

func (ecs EvaluationCandidateService) UpdateEvaluationCandidate(ctx *gin.Context, evaluationCandidate *dbContext.UpdateEvaluationCandidateParams, userEntityId int32) (*dbContext.BootcampEvaluationCandidate, *models.ResponseError) {
	responseErr := validateScoreEvaluation(evaluationCandidate)
	if responseErr != nil {
		return nil, responseErr
	}

	return ecs.repositoryManager.EvaluationCandidateRepository.UpdateEvaluationCandidate(ctx, evaluationCandidate, userEntityId)
}

func validateScoreEvaluation(evaluationCandidateParams *dbContext.UpdateEvaluationCandidateParams) *models.ResponseError {
	if evaluationCandidateParams.BtevSkor <= 0 || evaluationCandidateParams.BtevSkor > 100 {
		return &models.ResponseError{
			Message: "Invalid batch trainee evaluation score. It should be between 1 and 100.",
			Status:  http.StatusBadRequest,
		}
	}
	return nil
}
