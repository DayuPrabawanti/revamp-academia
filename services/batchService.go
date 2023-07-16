package services

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories"
	"codeid.revampacademy/repositories/bootcamp/dbContext"
	"github.com/gin-gonic/gin"
)

type BatchService struct {
	batchRepository *repositories.BatchRepository
}

func NewBatchService(batchRepository *repositories.BatchRepository) *BatchService {
	return &BatchService{
		batchRepository: batchRepository,
	}
}

func (bs BatchService) GetListBatch(ctx *gin.Context) ([]*models.BootcampBatch, *models.ResponseError) {
	return bs.batchRepository.GetListBatch(ctx)
}

func (bs BatchService) GetListBatchTraineeEvaluation(ctx *gin.Context) ([]*models.BootcampBatchTraineeEvaluation, *models.ResponseError) {
	return bs.batchRepository.GetListBatchTraineeEvaluation(ctx)
}

func (bs BatchService) GetBatch(ctx *gin.Context, id int64) (*models.BootcampBatch, *models.ResponseError) {
	return bs.batchRepository.GetBatch(ctx, id)
}

func (bs BatchService) GetBatchTraineeEvaluation(ctx *gin.Context, id int64) (*models.BootcampBatchTraineeEvaluation, *models.ResponseError) {
	return bs.batchRepository.GetBatchTraineeEvaluation(ctx, id)
}

func (bs BatchService) CreateBatch(ctx *gin.Context, batchParams *dbContext.CreateBatchParams) (*models.BootcampBatch, *models.ResponseError) {
	responseErr := validateBatch(batchParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return bs.batchRepository.CreateBatch(ctx, batchParams)
}

// func (bs BatchService) CreateBatchTraineeEvaluation(ctx *gin.Context, batchtevParams *dbContext.CreateBatchTraineeEvaluationParams) (*models.BootcampBatchTraineeEvaluation, *models.ResponseError) {
// 	responseErr := validateBatch(batchtevParams)
// 	if responseErr != nil {
// 		return nil, responseErr
// 	}

// 	return bs.batchRepository.CreateBatchTraineeEvaluation(ctx, batchtevParams)
// }

func (bs BatchService) UpdateBatch(ctx *gin.Context, batchParams *dbContext.CreateBatchParams, id int64) *models.ResponseError {
	responseErr := validateBatch(batchParams)
	if responseErr != nil {
		return responseErr
	}

	return bs.batchRepository.UpdateBatch(ctx, batchParams)
}

// func (bs BatchService) UpdateBatchTraineeEvaluation(ctx *gin.Context, batchtevParams *dbContext.CreateBatchParams, id int64) *models.ResponseError {
// 	responseErr := validateBatch(batchParams)
// 	if responseErr != nil {
// 		return responseErr
// 	}

// 	return bs.batchRepository.UpdateBatchTraineeEvaluation(ctx, batchtevParams)
// }

func (bs BatchService) DeleteBatch(ctx *gin.Context, id int64) *models.ResponseError {
	return bs.batchRepository.DeleteBatch(ctx, id)
}

func (bs BatchService) DeleteBatchTraineeEvaluation(ctx *gin.Context, id int64) *models.ResponseError {
	return bs.batchRepository.DeleteBatch(ctx, id)
}

func validateBatch(batchParams *dbContext.CreateBatchParams) *models.ResponseError {
	if batchParams.BatchID == 0 {
		return &models.ResponseError{
			Message: "Invalid batch id",
			Status:  http.StatusBadRequest,
		}
	}

	if batchParams.BatchName == "" {
		return &models.ResponseError{
			Message: "Invalid batch name",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}
