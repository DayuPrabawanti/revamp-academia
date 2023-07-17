package services

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories"
	"github.com/gin-gonic/gin"
)

type BatchTraineeEvaluationService struct {
	batchTraineeEvaluationRepository *repositories.BatchTraineeEvaluationRepository
}

func NewBatchTraineeEvaluationService(batchTraineeEvaluationRepository *repositories.BatchTraineeEvaluationRepository) *BatchTraineeEvaluationService {
	return &BatchTraineeEvaluationService{
		batchTraineeEvaluationRepository: batchTraineeEvaluationRepository,
	}
}

func (btes BatchTraineeEvaluationService) GetListBatchTraineeEvaluation(ctx *gin.Context) ([]*models.BootcampBatchTraineeEvaluation, *models.ResponseError) {
	return btes.batchTraineeEvaluationRepository.GetListBatchTraineeEvaluation(ctx)
}
