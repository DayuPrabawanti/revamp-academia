package bootcampService

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/bootcampRepository"
	"codeid.revampacademy/repositories/bootcampRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type BootcampBatchEvaluationService struct {
	// bootcampBatchEvaluationRepository *bootcampRepository.BootcampBatchEvaluationRepository
	repositoryManager *bootcampRepository.RepositoryManager
}

func NewBootcampBatchEvaluationService(repoMgr *bootcampRepository.RepositoryManager) *BootcampBatchEvaluationService {
	return &BootcampBatchEvaluationService{
		repositoryManager: repoMgr,
	}
}

func (bes BootcampBatchEvaluationService) GetBootcampBatchEvaluation(ctx *gin.Context, id int32) ([]*dbContext.BootcampBatchEvaluationMockup, *models.ResponseError) {
	return bes.repositoryManager.BootcampBatchEvaluationRepository.GetBootcampBatchEvaluation(ctx, id)
}

func (bes BootcampBatchEvaluationService) GetBatchTraineeReview(ctx *gin.Context, id int32) ([]*dbContext.BootcampBatchTraineeReview, *models.ResponseError) {
	return bes.repositoryManager.BootcampBatchEvaluationRepository.GetBatchTraineeReview(ctx, id)
}

func (bes BootcampBatchEvaluationService) CreateBatchTraineeReview(ctx *gin.Context, arg *dbContext.CreateBatchTraineeReviewParams) (*dbContext.BootcampBatchTraineeReview, *models.ResponseError) {
	responseErr := validateBootcampBatchReview(arg)
	if responseErr != nil {
		return nil, responseErr
	}

	return bes.repositoryManager.BootcampBatchEvaluationRepository.CreateBatchTraineeReview(ctx, arg)
}

func validateBootcampBatchReview(batchTraineeReviewParams *dbContext.CreateBatchTraineeReviewParams) *models.ResponseError {
	if batchTraineeReviewParams.BatrTraineeEntityID == 0 {
		return &models.ResponseError{
			Message: "Invalid batch trainee status",
			Status:  http.StatusBadRequest,
		}
	}

	if batchTraineeReviewParams.BatrStatus == "" {
		return &models.ResponseError{
			Message: "Invalid batch trainee status",
			Status:  http.StatusBadRequest,
		}
	}

	if batchTraineeReviewParams.BatrReview == "" {
		return &models.ResponseError{
			Message: "Invalid batch trainee review",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}

func (bes BootcampBatchEvaluationService) UpdateBatchTraineeReview(ctx *gin.Context, batchReviewParams *dbContext.UpdateBatchTraineeReviewParams, id int64) (*dbContext.BootcampBatchTraineeReview, *models.ResponseError) {
	responseErr := validateUpdateBatchReview(batchReviewParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return bes.repositoryManager.BootcampBatchEvaluationRepository.UpdateBatchTraineeReview(ctx, batchReviewParams)
}

func validateUpdateBatchReview(batchReviewParams *dbContext.UpdateBatchTraineeReviewParams) *models.ResponseError {
	if batchReviewParams.BatrStatus == "" {
		return &models.ResponseError{
			Message: "Invalid batch trainee review",
			Status:  http.StatusBadRequest,
		}
	}

	if batchReviewParams.BatrReview == "" {
		return &models.ResponseError{
			Message: "Invalid batch trainee review",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}

func (bes BootcampBatchEvaluationService) DeleteBatchTraineeReview(ctx *gin.Context, id int64) *models.ResponseError {
	return bes.repositoryManager.BootcampBatchEvaluationRepository.DeleteBatchTraineeReview(ctx, id)
}

func (bes BootcampBatchEvaluationService) BootcampBatchTraineeReview(ctx *gin.Context, id int32, BatrStatus, BatrReview string) (*dbContext.BootcampBatchTraineeReview, *models.ResponseError) {

	err := bootcampRepository.BeginTransaction(bes.repositoryManager)
	if err != nil {
		return nil, &models.ResponseError{
			Message: "Failed to start transaction",
			Status:  http.StatusBadRequest,
		}
	}

	// First, get the existing review using the provided ID
	_, responseErr := bes.GetBatchTraineeReview(ctx, id)
	if responseErr != nil {
		bootcampRepository.RollbackTransaction(bes.repositoryManager)
		return nil, responseErr
	}

	// Update the review
	updateParams := &dbContext.UpdateBatchTraineeReviewParams{
		BatrStatus: BatrStatus,
		BatrReview: BatrReview,
	}
	updatedReview, responseErr := bes.UpdateBatchTraineeReview(ctx, updateParams, int64(id))
	if responseErr != nil {
		bootcampRepository.RollbackTransaction(bes.repositoryManager)
		return nil, &models.ResponseError{
			Message: "Failed to update Status Review",
			Status:  http.StatusInternalServerError,
		}
	}

	// Commit the transaction
	bootcampRepository.CommitTransaction(bes.repositoryManager)

	return updatedReview, nil
}
