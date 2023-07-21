package bootcampService

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/bootcampRepository"
	"codeid.revampacademy/repositories/bootcampRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type BatchService struct {
	batchRepository *bootcampRepository.BatchRepository
}

func NewBatchService(batchRepository *bootcampRepository.BatchRepository) *BatchService {
	return &BatchService{
		batchRepository: batchRepository,
	}
}

func (bs BatchService) GetListBatch(ctx *gin.Context) ([]*models.BootcampBatch, *models.ResponseError) {
	return bs.batchRepository.GetListBatch(ctx)
}

func (bs BatchService) GetBatch(ctx *gin.Context, id int64) (*models.BootcampBatch, *models.ResponseError) {
	return bs.batchRepository.GetBatch(ctx, id)
}

func (bs BatchService) CreateBatch(ctx *gin.Context, batchParams *dbContext.CreateBatchParams) (*models.BootcampBatch, *models.ResponseError) {
	responseErr := validateBatch(batchParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return bs.batchRepository.CreateBatch(ctx, batchParams)
}

func (bs BatchService) UpdateBatch(ctx *gin.Context, batchParams *dbContext.CreateBatchParams, id int64) *models.ResponseError {
	responseErr := validateBatch(batchParams)
	if responseErr != nil {
		return responseErr
	}

	return bs.batchRepository.UpdateBatch(ctx, batchParams)
}

func (bs BatchService) DeleteBatch(ctx *gin.Context, id int64) *models.ResponseError {
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

func (bs BatchService) SearchBatch(ctx *gin.Context, batchName, status string) ([]models.BootcampBatch, *models.ResponseError) {
	// Perform validation, if needed, for batchName and status
	// If validation fails, return appropriate response error

	return bs.batchRepository.SearchBatch(ctx, batchName, status)
}

func (bs BatchService) PagingBatch(ctx *gin.Context, offset, pageSize int) ([]models.BootcampBatch, *models.ResponseError) {

	return bs.batchRepository.PagingBatch(ctx, offset, pageSize)
}
