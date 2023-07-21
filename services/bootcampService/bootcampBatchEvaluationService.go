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

func (bes BootcampBatchEvaluationService) GetListBootcampBatchEvaluation(ctx *gin.Context, batchId int32) ([]*models.BootcampBatchEvaluationMockup, *models.ResponseError) {
	return bes.bootcampBatchEvaluationRepository.GetListBootcampBatchEvaluation(ctx, batchId)
}

// func validateBootcampBatchEvaluation(bootcampBatchEvaluationParams *dbContext.CreateBootcampBatchEvaluationParams) *models.ResponseError {
// 	if bootcampBatchEvaluationParams.BatchID == 0 {
// 		return &models.ResponseError{
// 			Message: "Invalid batch id",
// 			Status:  http.StatusBadRequest,
// 		}
// 	}

// 	if bootcampBatchEvaluationParams.BatchName == "" {
// 		return &models.ResponseError{
// 			Message: "Invalid batch name",
// 			Status:  http.StatusBadRequest,
// 		}
// 	}

// 	return nil

// }
