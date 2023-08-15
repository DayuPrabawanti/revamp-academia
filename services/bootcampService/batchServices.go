package bootcampService

import (
	"context"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/models/features"
	"codeid.revampacademy/repositories/bootcampRepository"
	"codeid.revampacademy/repositories/bootcampRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type BatchService struct {
	repositoryManager *bootcampRepository.RepositoryManager
}

func NewBatchService(repoMgr *bootcampRepository.RepositoryManager) *BatchService {
	return &BatchService{
		repositoryManager: repoMgr,
	}
}

func (bs BatchService) GetListBatch(ctx *gin.Context, metadata *features.Metadata) ([]*models.BootcampBatchMockup, *models.ResponseError) {
	return bs.repositoryManager.BatchRepository.GetListBatch(ctx, metadata)
}

func (bs *BatchService) GetBatchWithMembers(ctx *gin.Context, id int64) (*models.BootcampBatchMockup, *models.ResponseError) {
	return bs.repositoryManager.BatchRepository.GetBatchWithMembers(ctx, int64(id))
}

func (bs BatchService) CreateBatch(ctx *gin.Context, batchParams *dbContext.CreateBatchParams) (*models.BootcampBatch, *models.ResponseError) {
	// responseErr := validateBatch(batchParams)
	// if responseErr != nil {
	// 	return nil, responseErr
	// }

	return bs.repositoryManager.BatchRepository.CreateBatch(ctx, batchParams)
}

func (bs BatchService) CreateInstructorPrograms(ctx *gin.Context, instructorProgramsParams *dbContext.CreateInstructorProgramsParams) (*models.BootcampInstructorProgram, *models.ResponseError) {
	return bs.repositoryManager.BatchRepository.CreateInstructorPrograms(ctx, instructorProgramsParams)
}

func (bs BatchService) CreateBatchTrainee(ctx *gin.Context, batchTraineeParams *dbContext.CreateBatchTraineeParams) (*models.BootcampBatchTrainee, *models.ResponseError) {
	return bs.repositoryManager.BatchRepository.CreateBatchTrainee(ctx, batchTraineeParams)
}

func (bs BatchService) UpdateBatch(ctx *gin.Context, batchParams *dbContext.CreateBatchParams, id int32) *models.ResponseError {
	responseErr := validateBatch(batchParams)
	if responseErr != nil {
		return responseErr
	}

	return bs.repositoryManager.BatchRepository.UpdateBatch(ctx, batchParams)
}

func (bs BatchService) UpdateInstructorPrograms(ctx *gin.Context, instructorProgramsParams *dbContext.CreateInstructorProgramsParams, id int32) *models.ResponseError {
	responseErr := validateInstructorPrograms(instructorProgramsParams)
	if responseErr != nil {
		return responseErr
	}

	return bs.repositoryManager.BatchRepository.UpdateInstructorPrograms(ctx, instructorProgramsParams)
}

func (bs BatchService) DeleteBatch(ctx *gin.Context, id int64) *models.ResponseError {
	return bs.repositoryManager.BatchRepository.DeleteBatch(ctx, id)
}

func (bs BatchService) DeleteInstructorPrograms(ctx *gin.Context, id int64) *models.ResponseError {
	return bs.repositoryManager.BatchRepository.DeleteInstructorPrograms(ctx, id)
}

func (bs BatchService) DeleteBatchTrainee(ctx *gin.Context, id int64) *models.ResponseError {
	return bs.repositoryManager.BatchRepository.DeleteBatchTrainee(ctx, id)
}

func (bs BatchService) DeleteBatchTrainee2(ctx *gin.Context, id int64, batch int64) *models.ResponseError {
	return bs.repositoryManager.BatchRepository.DeleteBatchTrainee2(ctx, id, batch)
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
func validateInstructorPrograms(instructorProgramsParams *dbContext.CreateInstructorProgramsParams) *models.ResponseError {
	if instructorProgramsParams.BatchID == 0 {
		return &models.ResponseError{
			Message: "Invalid batch id",
			Status:  http.StatusBadRequest,
		}
	}

	if instructorProgramsParams.InproEntityID == 0 {
		return &models.ResponseError{
			Message: "Invalid instructor programs entity id",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}

func (bs BatchService) SearchBatch(ctx *gin.Context, batchName, status string) ([]models.BootcampBatch, *models.ResponseError) {

	return bs.repositoryManager.BatchRepository.SearchBatch(ctx, batchName, status)
}

func (bs BatchService) CreateBatchInstructorTraineeDto(ctx *gin.Context, batchInstructorTraineeDto *models.CreateBatchInstructorTraineeDto) (*models.BootcampBatch, *models.ResponseError) {
	err := bootcampRepository.BeginTransaction(bs.repositoryManager)
	if err != nil {
		return nil, &models.ResponseError{
			Message: "Failed to start transaction",
			Status:  http.StatusBadRequest,
		}
	}

	// Create batch using the provided DTO
	response, responseErr := bs.CreateBatch(ctx, (*dbContext.CreateBatchParams)(&batchInstructorTraineeDto.CreateBatchDto))
	if responseErr != nil {
		bootcampRepository.RollbackTransaction(bs.repositoryManager)
		return nil, responseErr
	}

	// Create instructor programs associated with the batch
	for _, instructor := range batchInstructorTraineeDto.Instructor {
		instructorProgramsParams := &dbContext.CreateInstructorProgramsParams{
			BatchID:          response.BatchID, // Use the ID of the created batch
			InproEntityID:    instructor.InproEntityID,
			InproEmpEntityID: instructor.InproEmpEntityID,
		}
		_, responseErr = bs.CreateInstructorPrograms(ctx, instructorProgramsParams)
		if responseErr != nil {
			bootcampRepository.RollbackTransaction(bs.repositoryManager)
			return nil, responseErr
		}
	}

	// Create trainees associated with the batch
	for _, trainee := range batchInstructorTraineeDto.Trainee {
		batchTraineeParams := &dbContext.CreateBatchTraineeParams{
			BatrTraineeEntityID: trainee.BatrTraineeEntityID,
			BatrBatchID:         response.BatchID, // Use the ID of the created batch
		}
		_, responseErr = bs.CreateBatchTrainee(ctx, batchTraineeParams)
		if responseErr != nil {
			bootcampRepository.RollbackTransaction(bs.repositoryManager)
			return nil, responseErr
		}
	}

	bootcampRepository.CommitTransaction(bs.repositoryManager)

	return nil, &models.ResponseError{
		Message: "Data has been created",
		Status:  http.StatusOK,
	}
}

func (bs BatchService) DeleteBatchTransaction(ctx *gin.Context, id int64) *models.ResponseError {
	err := bs.repositoryManager.BatchRepository.DeleteBatchTransaction(context.Background(), id)
	if err != nil {
		return &models.ResponseError{
			Message: "Failed to delete batch transaction",
			Status:  http.StatusInternalServerError,
		}
	}

	return &models.ResponseError{
		Message: "Instructor programs, batch trainee data, and batch data have been deleted",
		Status:  http.StatusOK,
	}
}
