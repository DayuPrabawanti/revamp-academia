package dbContext

import (
	"context"
	"net/http"
	"time"

	"codeid.revampacademy/models"
)

const createBatch = `-- name: CreateBatch :one
INSERT INTO bootcamp.batch
(batch_id, batch_entity_id, batch_name, batch_description, batch_start_date, batch_end_date, batch_reason, batch_type, batch_modified_date, batch_status, batch_pic_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING *
`

type CreateBatchParams struct {
	BatchID           int32     `db:"batch_id" json:"batchId"`
	BatchEntityID     int32     `db:"batch_entity_id" json:"batchEntityId"`
	BatchName         string    `db:"batch_name" json:"batchName"`
	BatchDescription  string    `db:"batch_description" json:"batchDescription"`
	BatchStartDate    time.Time `db:"batch_start_date" json:"batchStartDate"`
	BatchEndDate      time.Time `db:"batch_end_date" json:"batchEndDate"`
	BatchReason       string    `db:"batch_reason" json:"batchReason"`
	BatchType         string    `db:"batch_type" json:"batchType"`
	BatchModifiedDate time.Time `db:"batch_modified_date" json:"batchModifiedDate"`
	BatchStatus       string    `db:"batch_status" json:"batchStatus"`
	BatchPicID        int32     `db:"batch_pic_id" json:"batchPicId"`
}

func (q *Queries) CreateBatch(ctx context.Context, arg CreateBatchParams) (*models.BootcampBatch, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createBatch,
		arg.BatchID,
		arg.BatchEntityID,
		arg.BatchName,
		arg.BatchDescription,
		arg.BatchStartDate,
		arg.BatchEndDate,
		arg.BatchReason,
		arg.BatchType,
		arg.BatchModifiedDate,
		arg.BatchStatus,
		arg.BatchPicID,
	)
	i := models.BootcampBatch{}
	err := row.Scan(
		&i.BatchID,
		&i.BatchEntityID,
		&i.BatchName,
		&i.BatchDescription,
		&i.BatchStartDate,
		&i.BatchEndDate,
		&i.BatchReason,
		&i.BatchType,
		&i.BatchModifiedDate,
		&i.BatchStatus,
		&i.BatchPicID,
	)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.BootcampBatch{
		BatchID:           i.BatchID,
		BatchEntityID:     i.BatchEntityID,
		BatchName:         i.BatchName,
		BatchDescription:  i.BatchDescription,
		BatchStartDate:    i.BatchStartDate,
		BatchEndDate:      i.BatchEndDate,
		BatchReason:       i.BatchReason,
		BatchType:         i.BatchType,
		BatchModifiedDate: i.BatchModifiedDate,
		BatchStatus:       i.BatchStatus,
		BatchPicID:        i.BatchPicID,
	}, nil
}

const listBatchs = `-- name: ListBatchs :many
SELECT batch_id, batch_entity_id, batch_name, batch_description, batch_start_date, batch_end_date, batch_reason, batch_type, batch_modified_date, batch_status, batch_pic_id FROM bootcamp.batch
ORDER BY batch_name
`

func (q *Queries) ListBatchs(ctx context.Context) ([]models.BootcampBatch, error) {
	rows, err := q.db.QueryContext(ctx, listBatchs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.BootcampBatch
	for rows.Next() {
		var i models.BootcampBatch
		if err := rows.Scan(
			&i.BatchID,
			&i.BatchEntityID,
			&i.BatchName,
			&i.BatchDescription,
			&i.BatchStartDate,
			&i.BatchEndDate,
			&i.BatchReason,
			&i.BatchType,
			&i.BatchModifiedDate,
			&i.BatchStatus,
			&i.BatchPicID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getBatch = `-- name: GetBatch :one
SELECT batch_id, batch_entity_id, batch_name, batch_description, batch_start_date, batch_end_date, batch_reason, batch_type, batch_modified_date, batch_status, batch_pic_id FROM bootcamp.batch
WHERE batch_id = $1
`

func (q *Queries) GetBatch(ctx context.Context, batchID int32) (models.BootcampBatch, error) {
	row := q.db.QueryRowContext(ctx, getBatch, batchID)
	var i models.BootcampBatch
	err := row.Scan(
		&i.BatchID,
		&i.BatchEntityID,
		&i.BatchName,
		&i.BatchDescription,
		&i.BatchStartDate,
		&i.BatchEndDate,
		&i.BatchReason,
		&i.BatchType,
		&i.BatchModifiedDate,
		&i.BatchStatus,
		&i.BatchPicID,
	)
	return i, err
}

const updateBatch = `-- name: UpdateBatch :exec
UPDATE bootcamp.batch
SET batch_name = $2,
    batch_description = $3
WHERE batch_id = $1
`

type UpdateBatchParams struct {
	BatchID          int32  `db:"batch_id" json:"batchId"`
	BatchName        string `db:"batch_name" json:"batchName"`
	BatchDescription string `db:"batch_description" json:"batchDescription"`
}

func (q *Queries) UpdateBatch(ctx context.Context, arg CreateBatchParams) error {
	_, err := q.db.ExecContext(ctx, updateBatch, arg.BatchID, arg.BatchName, arg.BatchDescription)
	return err
}

const deleteBatch = `-- name: DeleteBatch :exec
DELETE  FROM bootcamp.batch
WHERE batch_id = $1
`

func (q *Queries) DeleteBatch(ctx context.Context, batchID int32) error {
	_, err := q.db.ExecContext(ctx, deleteBatch, batchID)
	return err
}

const searchBatch = `-- name: SearchBatch :many
SELECT batch_id, batch_entity_id, batch_name, batch_description, batch_start_date, batch_end_date, batch_reason, batch_type, batch_modified_date, batch_status, batch_pic_id FROM bootcamp.batch
WHERE batch_name like '%' || $1 || '%' AND batch_status = $2
`

func (q *Queries) SearchBatch(ctx context.Context, batchName string, status string) ([]models.BootcampBatch, error) {
	rows, err := q.db.QueryContext(ctx, searchBatch, batchName, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var batches []models.BootcampBatch
	for rows.Next() {
		var b models.BootcampBatch
		if err := rows.Scan(
			&b.BatchID,
			&b.BatchEntityID,
			&b.BatchName,
			&b.BatchDescription,
			&b.BatchStartDate,
			&b.BatchEndDate,
			&b.BatchReason,
			&b.BatchType,
			&b.BatchModifiedDate,
			&b.BatchStatus,
			&b.BatchPicID,
		); err != nil {
			return nil, err
		}
		batches = append(batches, b)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return batches, nil
}

const pagingBatch = `-- name: PagingBatch :many
SELECT batch_id, batch_entity_id, batch_name, batch_description, batch_start_date, batch_end_date, batch_reason, batch_type, batch_modified_date, batch_status, batch_pic_id FROM bootcamp.batch
ORDER BY batch_name
LIMIT $1 OFFSET $2
`

func (q *Queries) PagingBatch(ctx context.Context, limit, offset int) ([]models.BootcampBatch, error) {
	rows, err := q.db.QueryContext(ctx, pagingBatch, limit, offset)
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
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return batches, nil
}
