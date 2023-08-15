package bootcampRepository

import (
	"context"
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/models/features"
	"codeid.revampacademy/repositories/bootcampRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type BatchRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
	dbQueries   dbContext.Queries
}

func NewBatchRepository(dbHandler *sql.DB) *BatchRepository {
	return &BatchRepository{
		dbHandler: dbHandler,
		//add new fields
		dbQueries: *dbContext.New(dbHandler),
	}
}

func (br *BatchRepository) GetListBatch(ctx context.Context, metadata *features.Metadata) ([]*models.BootcampBatchMockup, *models.ResponseError) {
	batchs, err := br.dbQueries.ListBatchs(ctx, metadata)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	listBatchs := make([]*models.BootcampBatchMockup, 0)
	for _, v := range batchs {
		batch := &models.BootcampBatchMockup{
			BatchID:        v.BatchID,
			BatchName:      v.BatchName,
			ProgTitle:      v.ProgTitle,
			BatchStartDate: v.BatchStartDate,
			BatchEndDate:   v.BatchEndDate,
			UserName:       v.UserName,
			BatchStatus:    v.BatchStatus,
			Members:        v.Members,
		}
		listBatchs = append(listBatchs, batch)
	}

	return listBatchs, nil
}

func (br BatchRepository) GetBatchWithMembers(ctx *gin.Context, id int64) (*models.BootcampBatchMockup, *models.ResponseError) {
	store := dbContext.New(br.dbHandler)
	batch, err := store.GetBatchWithMembers(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return batch, nil
}

func (br BatchRepository) CreateBatch(ctx *gin.Context, batchParams *dbContext.CreateBatchParams) (*models.BootcampBatch, *models.ResponseError) {

	store := dbContext.New(br.dbHandler)
	batch, err := store.CreateBatch(ctx, *batchParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return batch, nil
}

func (br BatchRepository) CreateInstructorPrograms(ctx *gin.Context, instructorProgramsParams *dbContext.CreateInstructorProgramsParams) (*models.BootcampInstructorProgram, *models.ResponseError) {

	store := dbContext.New(br.dbHandler)
	batch, err := store.CreateInstructorPrograms(ctx, *instructorProgramsParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return batch, nil
}

func (br BatchRepository) CreateBatchTrainee(ctx *gin.Context, batchTraineeParams *dbContext.CreateBatchTraineeParams) (*models.BootcampBatchTrainee, *models.ResponseError) {

	store := dbContext.New(br.dbHandler)
	batch, err := store.CreateBatchTrainee(ctx, *batchTraineeParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return batch, nil
}

func (br BatchRepository) UpdateBatch(ctx *gin.Context, batchParams *dbContext.CreateBatchParams) *models.ResponseError {

	store := dbContext.New(br.dbHandler)
	err := store.UpdateBatch(ctx, *batchParams)

	if err != nil {
		return &models.ResponseError{
			Message: "error when update batch",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data  batch has been update",
		Status:  http.StatusOK,
	}
}

func (br BatchRepository) UpdateInstructorPrograms(ctx *gin.Context, instructorProgramsParams *dbContext.CreateInstructorProgramsParams) *models.ResponseError {

	store := dbContext.New(br.dbHandler)
	err := store.UpdateInstructorPrograms(ctx, *instructorProgramsParams)

	if err != nil {
		return &models.ResponseError{
			Message: "error when update instructor programs",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data instructor programs has been update",
		Status:  http.StatusOK,
	}
}

func (br BatchRepository) DeleteBatch(ctx *gin.Context, id int64) *models.ResponseError {

	store := dbContext.New(br.dbHandler)
	err := store.DeleteBatch(ctx, int32(id))

	if err != nil {
		return &models.ResponseError{
			Message: "error when update batch",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been deleted",
		Status:  http.StatusOK,
	}
}

func (br BatchRepository) DeleteInstructorPrograms(ctx *gin.Context, id int64) *models.ResponseError {

	store := dbContext.New(br.dbHandler)
	err := store.DeleteInstructorPrograms(ctx, int32(id))

	if err != nil {
		return &models.ResponseError{
			Message: "error when update instructor programs",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been deleted",
		Status:  http.StatusOK,
	}
}

func (br BatchRepository) DeleteBatchTrainee(ctx *gin.Context, id int64) *models.ResponseError {

	store := dbContext.New(br.dbHandler)
	err := store.DeleteBatchTrainee(ctx, int32(id))

	if err != nil {
		return &models.ResponseError{
			Message: "error when update batch trainee",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been deleted",
		Status:  http.StatusOK,
	}
}

func (br BatchRepository) DeleteBatchTrainee2(ctx *gin.Context, id int64, batch int64) *models.ResponseError {
	store := dbContext.New(br.dbHandler)
	err := store.DeleteBatchTrainee2(ctx, int32(id), int32(batch))

	if err != nil {
		return &models.ResponseError{
			Message: "Failed to delete batch trainee data",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "Batch trainee data has been deleted",
		Status:  http.StatusOK,
	}
}

func (br BatchRepository) SearchBatch(ctx *gin.Context, batchName, status string) ([]models.BootcampBatch, *models.ResponseError) {

	store := dbContext.New(br.dbHandler)
	batches, err := store.SearchBatch(ctx, batchName, status)
	if err != nil {
		return nil, &models.ResponseError{
			Message: "Failed to search batches",
			Status:  http.StatusInternalServerError,
		}
	}

	return batches, nil
}

func (br BatchRepository) DeleteBatchTransaction(ctx context.Context, id int64) error {
	store := dbContext.New(br.dbHandler)

	// Delete instructor programs
	err := store.DeleteInstructorPrograms(ctx, int32(id))
	if err != nil {
		return err
	}

	// Delete batch trainees
	err = store.DeleteBatchTrainee(ctx, int32(id))
	if err != nil {
		return err
	}

	// Delete batch
	err = store.DeleteBatch(ctx, int32(id))
	if err != nil {
		return err
	}

	return nil
}
