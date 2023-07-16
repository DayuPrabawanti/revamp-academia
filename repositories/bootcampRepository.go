package repositories

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/bootcamp/dbContext"
	"github.com/gin-gonic/gin"
)

type BatchRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewBatchRepository(dbHandler *sql.DB) *BatchRepository {
	return &BatchRepository{
		dbHandler: dbHandler,
	}
}

func (br BatchRepository) GetListBatch(ctx *gin.Context) ([]*models.BootcampBatch, *models.ResponseError) {

	store := dbContext.New(br.dbHandler)
	batchs, err := store.ListBatchs(ctx)

	listBatchs := make([]*models.BootcampBatch, 0)

	for _, v := range batchs {
		batch := &models.BootcampBatch{
			BatchID:           v.BatchID,
			BatchEntityID:     v.BatchEntityID,
			BatchName:         v.BatchName,
			BatchDescription:  v.BatchDescription,
			BatchStartDate:    v.BatchStartDate,
			BatchEndDate:      v.BatchEndDate,
			BatchReason:       v.BatchReason,
			BatchType:         v.BatchType,
			BatchModifiedDate: v.BatchModifiedDate,
			BatchStatus:       v.BatchStatus,
			BatchPicID:        v.BatchPicID,
		}
		listBatchs = append(listBatchs, batch)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listBatchs, nil
}

func (br BatchRepository) GetListBatchTraineeEvaluation(ctx *gin.Context) ([]*models.BootcampBatchTraineeEvaluation, *models.ResponseError) {

	store := dbContext.New(br.dbHandler)
	btev, err := store.ListBatchTraineeEvaluations(ctx)

	listBatchTraineeEvaluations := make([]*models.BootcampBatchTraineeEvaluation, 0)

	for _, v := range btev {
		btrev := &models.BootcampBatchTraineeEvaluation{
			BtevID:              v.BtevID,
			BtevType:            v.BtevType,
			BtevHeader:          v.BtevHeader,
			BtevSection:         v.BtevSection,
			BtevSkill:           v.BtevSkill,
			BtevWeek:            v.BtevWeek,
			BtevSkor:            v.BtevSkor,
			BtevNote:            v.BtevNote,
			BtevModifiedDate:    v.BtevModifiedDate,
			BtevBatchID:         v.BtevBatchID,
			BtevTraineeEntityID: v.BtevTraineeEntityID,
		}
		listBatchTraineeEvaluations = append(listBatchTraineeEvaluations, btrev)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listBatchTraineeEvaluations, nil
}

func (br BatchRepository) GetBatch(ctx *gin.Context, id int64) (*models.BootcampBatch, *models.ResponseError) {

	store := dbContext.New(br.dbHandler)
	batch, err := store.GetBatch(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &batch, nil
}

func (br BatchRepository) GetBatchTraineeEvaluation(ctx *gin.Context, id int64) (*models.BootcampBatchTraineeEvaluation, *models.ResponseError) {

	store := dbContext.New(br.dbHandler)
	btev, err := store.GetBatchTraineeEvaluation(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &btev, nil
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

func (br BatchRepository) CreateBatchTraineeEvaluation(ctx *gin.Context, batchtev *dbContext.CreateBatchTraineeEvaluationParams) (*models.BootcampBatchTraineeEvaluation, *models.ResponseError) {

	store := dbContext.New(br.dbHandler)
	btev, err := store.CreateBatchTraineeEvaluation(ctx, *batchtev)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return btev, nil
}

func (br BatchRepository) UpdateBatch(ctx *gin.Context, batchParams *dbContext.CreateBatchParams) *models.ResponseError {

	store := dbContext.New(br.dbHandler)
	err := store.UpdateBatch(ctx, *batchParams)

	if err != nil {
		return &models.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been update",
		Status:  http.StatusOK,
	}
}

func (br BatchRepository) UpdateBatchTraineeEvaluation(ctx *gin.Context, batchtevParams *dbContext.CreateBatchParams) *models.ResponseError {

	store := dbContext.New(br.dbHandler)
	err := store.UpdateBatch(ctx, *batchtevParams)

	if err != nil {
		return &models.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been update",
		Status:  http.StatusOK,
	}
}

func (br BatchRepository) DeleteBatch(ctx *gin.Context, id int64) *models.ResponseError {

	store := dbContext.New(br.dbHandler)
	err := store.DeleteBatch(ctx, int32(id))

	if err != nil {
		return &models.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been deleted",
		Status:  http.StatusOK,
	}
}

func (br BatchRepository) DeleteBatchTraineeEvaluation(ctx *gin.Context, id int64) *models.ResponseError {

	store := dbContext.New(br.dbHandler)
	err := store.DeleteBatchTraineeEvaluation(ctx, int32(id))

	if err != nil {
		return &models.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been deleted",
		Status:  http.StatusOK,
	}
}
