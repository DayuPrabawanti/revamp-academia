package repositories

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/dbContext"
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

func (br BatchRepository) SearchBatch(ctx *gin.Context, batchID int32, status string) ([]models.BootcampBatch, error) {
	// Prepare the SQL query to search batches based on batchID and status.
	// Note: Modify the query according to your database schema.

	const searchBatchSQL = `
		SELECT batch_id, batch_entity_id, batch_name, batch_description, batch_start_date,
			batch_end_date, batch_reason, batch_type, batch_modified_date, batch_status, batch_pic_id
		FROM bootcamp.batch
		WHERE batch_id = $1 AND batch_status = $2
	`

	rows, err := br.dbHandler.QueryContext(ctx, searchBatchSQL, batchID, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var batches []models.BootcampBatch
	for rows.Next() {
		var batch models.BootcampBatch
		if err := rows.Scan(
			&batch.BatchID,
			&batch.BatchEntityID,
			&batch.BatchName,
			&batch.BatchDescription,
			&batch.BatchStartDate,
			&batch.BatchEndDate,
			&batch.BatchReason,
			&batch.BatchType,
			&batch.BatchModifiedDate,
			&batch.BatchStatus,
			&batch.BatchPicID,
		); err != nil {
			return nil, err
		}
		batches = append(batches, batch)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return batches, nil
}
