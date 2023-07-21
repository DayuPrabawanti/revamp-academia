package bootcampService

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/bootcampRepository"
	"github.com/gin-gonic/gin"
)

type BatchTraineeEvaluationService struct {
	batchTraineeEvaluationRepository *bootcampRepository.BatchTraineeEvaluationRepository
}

func NewBatchTraineeEvaluationService(batchTraineeEvaluationRepository *bootcampRepository.BatchTraineeEvaluationRepository) *BatchTraineeEvaluationService {
	return &BatchTraineeEvaluationService{
		batchTraineeEvaluationRepository: batchTraineeEvaluationRepository,
	}
}

func (btes BatchTraineeEvaluationService) GetListBatchTraineeEvaluation(ctx *gin.Context) ([]*models.BootcampBatchTraineeEvaluation, *models.ResponseError) {
	return btes.batchTraineeEvaluationRepository.GetListBatchTraineeEvaluation(ctx)
}
