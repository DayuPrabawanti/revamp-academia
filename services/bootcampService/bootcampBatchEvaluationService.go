package bootcampService

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/bootcampRepository"
	"github.com/gin-gonic/gin"
)

type BootcampBatchEvaluationService struct {
	bootcampBatchEvaluationRepository *bootcampRepository.BootcampBatchEvaluationRepository
}

func NewBootcampBatchEvaluationService(bootcampBatchEvaluationRepository *bootcampRepository.BootcampBatchEvaluationRepository) *BootcampBatchEvaluationService {
	return &BootcampBatchEvaluationService{
		bootcampBatchEvaluationRepository: bootcampBatchEvaluationRepository,
	}
}

func (bes BootcampBatchEvaluationService) GetListBootcampBatchEvaluation(ctx *gin.Context) ([]*models.BootcampBatchEvaluationMockup, *models.ResponseError) {
	return bes.bootcampBatchEvaluationRepository.GetListBootcampBatchEvaluation(ctx)
}
