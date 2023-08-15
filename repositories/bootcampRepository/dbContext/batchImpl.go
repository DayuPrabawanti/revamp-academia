package dbContext

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"codeid.revampacademy/models"
	"codeid.revampacademy/models/features"
)

const createBatch = `-- name: CreateBatch :one
INSERT INTO bootcamp.batch
(batch_id,batch_entity_id, batch_name,batch_start_date, batch_end_date, batch_status, batch_modified_date)
VALUES ($1, $2, $3, $4, $5, $6, Now())
RETURNING batch_id, batch_entity_id, batch_name,batch_start_date, batch_end_date, batch_status, batch_modified_date
`

type CreateBatchParams struct {
	BatchID           int32     `db:"batch_id" json:"batchId"`
	BatchEntityID     int32     `db:"batch_entity_id" json:"batchEntityId"`
	BatchName         string    `db:"batch_name" json:"batchName"`
	BatchStartDate    time.Time `db:"batch_start_date" json:"batchStartDate"`
	BatchEndDate      time.Time `db:"batch_end_date" json:"batchEndDate"`
	BatchStatus       string    `db:"batch_status" json:"batchStatus"`
	BatchModifiedDate time.Time `db:"batch_modified_date" json:"batchModifiedDate"`
}

func (q *Queries) CreateBatch(ctx context.Context, arg CreateBatchParams) (*models.BootcampBatch, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createBatch,
		arg.BatchID,
		arg.BatchEntityID,
		arg.BatchName,
		arg.BatchStartDate,
		arg.BatchEndDate,
		arg.BatchStatus,
	)
	i := models.BootcampBatch{}
	err := row.Scan(
		&i.BatchID,
		&i.BatchEntityID,
		&i.BatchName,
		&i.BatchStartDate,
		&i.BatchEndDate,
		&i.BatchStatus,
		&i.BatchModifiedDate,
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
		BatchStartDate:    i.BatchStartDate,
		BatchEndDate:      i.BatchEndDate,
		BatchStatus:       i.BatchStatus,
		BatchModifiedDate: i.BatchModifiedDate,
	}, nil
}

const createInstructorPrograms = `-- name: CreateInstructorPrograms :one
INSERT INTO bootcamp.instructor_programs
(batch_id, inpro_entity_id,inpro_emp_entity_id, inpro_modified_date)
VALUES ($1, $2, $3, Now())
RETURNING *
`

type CreateInstructorProgramsParams struct {
	BatchID           int32     `db:"batch_id" json:"batchId"`
	InproEntityID     int32     `db:"inpro_entity_id" json:"inproEntityId"`
	InproEmpEntityID  int32     `db:"inpro_emp_entity_id" json:"inproEmpEntityId"`
	InproModifiedDate time.Time `db:"inpro_modified_date" json:"inproModifiedDate"`
}

func (q *Queries) CreateInstructorPrograms(ctx context.Context, arg CreateInstructorProgramsParams) (*models.BootcampInstructorProgram, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createInstructorPrograms,
		arg.BatchID,
		arg.InproEntityID,
		arg.InproEmpEntityID,
	)
	i := models.BootcampInstructorProgram{}
	err := row.Scan(
		&i.BatchID,
		&i.InproEntityID,
		&i.InproEmpEntityID,
		&i.InproModifiedDate,
	)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.BootcampInstructorProgram{
		BatchID:           i.BatchID,
		InproEntityID:     i.InproEntityID,
		InproEmpEntityID:  i.InproEmpEntityID,
		InproModifiedDate: i.InproModifiedDate,
	}, nil
}

const createBatchTrainee = `-- name: CreateBatchTrainee :one
INSERT INTO bootcamp.batch_trainee
(batr_modified_date, batr_trainee_entity_id, batr_batch_id)
VALUES (Now(), $1, $2)
RETURNING batr_modified_date, batr_trainee_entity_id, batr_batch_id
`

type CreateBatchTraineeParams struct {
	BatrModifiedDate    time.Time `db:"batr_modified_date" json:"batrModifiedDate"`
	BatrTraineeEntityID int32     `db:"batr_trainee_entity_id" json:"batrTraineeEntityId"`
	BatrBatchID         int32     `db:"batr_batch_id" json:"batrBatchId"`
}

func (q *Queries) CreateBatchTrainee(ctx context.Context, arg CreateBatchTraineeParams) (*models.BootcampBatchTrainee, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createBatchTrainee,
		arg.BatrTraineeEntityID,
		arg.BatrBatchID,
	)
	i := models.BootcampBatchTrainee{}
	err := row.Scan(
		&i.BatrModifiedDate,
		&i.BatrTraineeEntityID,
		&i.BatrBatchID,
	)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.BootcampBatchTrainee{
		BatrModifiedDate:    i.BatrModifiedDate,
		BatrTraineeEntityID: i.BatrTraineeEntityID,
		BatrBatchID:         i.BatrBatchID,
	}, nil
}

const listBatchs = `-- name: ListBatchs :many
SELECT
    b.batch_id,
    b.batch_name,
    pe.prog_title AS technology,
    b.batch_start_date,
    b.batch_end_date,
    u.user_name AS trainer,
    b.batch_status,
    json_agg(json_build_object('userPhoto', tu.user_photo)) AS members
FROM bootcamp.batch b
LEFT JOIN curriculum.program_entity pe ON b.batch_entity_id = pe.prog_entity_id
LEFT JOIN bootcamp.instructor_programs ip ON b.batch_id = ip.batch_id
LEFT JOIN users.users u ON ip.inpro_entity_id = u.user_entity_id
LEFT JOIN bootcamp.batch_trainee bt ON b.batch_id = bt.batr_batch_id
LEFT JOIN users.users tu ON bt.batr_trainee_entity_id = tu.user_entity_id
WHERE batch_name like '%' || $3 || '%' OR batch_status = $4
GROUP BY b.batch_id, b.batch_name, pe.prog_title, b.batch_start_date, b.batch_end_date, u.user_name, b.batch_status
limit $1 offset $2
`

func (q *Queries) ListBatchs(ctx context.Context, metadata *features.Metadata) ([]*models.BootcampBatchMockup, error) {
	rows, err := q.db.QueryContext(ctx, listBatchs, metadata.PageSize, metadata.Page, metadata.Batch, metadata.Status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var batches []*models.BootcampBatchMockup
	for rows.Next() {
		var batch models.BootcampBatchMockup
		var membersJSON sql.NullString

		err := rows.Scan(
			&batch.BatchID,
			&batch.BatchName,
			&batch.ProgTitle,
			&batch.BatchStartDate,
			&batch.BatchEndDate,
			&batch.UserName,
			&batch.BatchStatus,
			&membersJSON,
		)

		if err != nil {
			return nil, err
		}

		if membersJSON.Valid {
			var members []models.BatchMember
			err = json.Unmarshal([]byte(membersJSON.String), &members)
			if err != nil {
				return nil, err
			}

			batch.Members = members
		}

		batches = append(batches, &batch)
	}

	return batches, nil
}

const getBatchWithMembers = `
SELECT
    b.batch_id,
    b.batch_name,
    pe.prog_title AS technology,
    b.batch_start_date,
    b.batch_end_date,
    u.user_name AS trainer,
    b.batch_status,
    json_agg(json_build_object('userPhoto', tu.user_photo)) AS members
FROM bootcamp.batch b
JOIN curriculum.program_entity pe ON b.batch_entity_id = pe.prog_entity_id
JOIN bootcamp.instructor_programs ip ON b.batch_id = ip.batch_id
JOIN users.users u ON ip.inpro_entity_id = u.user_entity_id
LEFT JOIN bootcamp.batch_trainee bt ON b.batch_id = bt.batr_batch_id
LEFT JOIN users.users tu ON bt.batr_trainee_entity_id = tu.user_entity_id
WHERE b.batch_id = $1
GROUP BY b.batch_id, b.batch_name, pe.prog_title, b.batch_start_date, b.batch_end_date, u.user_name, b.batch_status
`

func (q *Queries) GetBatchWithMembers(ctx context.Context, batchID int32) (*models.BootcampBatchMockup, error) {
	row := q.db.QueryRowContext(ctx, getBatchWithMembers, batchID)

	batch := models.BootcampBatchMockup{}

	memberPhotosJSON := sql.NullString{}

	err := row.Scan(
		&batch.BatchID,
		&batch.BatchName,
		&batch.ProgTitle,
		&batch.BatchStartDate,
		&batch.BatchEndDate,
		&batch.UserName,
		&batch.BatchStatus,
		&memberPhotosJSON,
	)

	if err != nil {
		return nil, err
	}

	if memberPhotosJSON.Valid {
		var members []models.BatchMember
		err = json.Unmarshal([]byte(memberPhotosJSON.String), &members)
		if err != nil {
			return nil, err
		}

		batch.Members = members
	}

	return &batch, nil
}

const updateBatch = `-- name: UpdateBatch :exec
UPDATE bootcamp.batch
SET batch_entity_id = $2,
	batch_name = $3,
	batch_start_date = $4,
	batch_end_date = $5,
	batch_status = $6,
	batch_modified_date = Now()
WHERE batch_id = $1
`

// type UpdateBatchParams struct {
// 	// BatchID           int32     `db:"batch_id" json:"batchId"`
// 	BatchEntityID     int32     `db:"batch_entity_id" json:"batchEntityId"`
// 	BatchName         string    `db:"batch_name" json:"batchName"`
// 	BatchStartDate    time.Time `db:"batch_start_date" json:"batchStartDate"`
// 	BatchEndDate      time.Time `db:"batch_end_date" json:"batchEndDate"`
// 	BatchStatus       string    `db:"batch_status" json:"batchStatus"`
// 	BatchModifiedDate time.Time `db:"batch_modified_date" json:"batchModifiedDate"`
// }

func (q *Queries) UpdateBatch(ctx context.Context, arg CreateBatchParams) error {
	_, err := q.db.ExecContext(ctx, updateBatch,
		arg.BatchID,
		arg.BatchEntityID,
		arg.BatchName,
		arg.BatchStartDate,
		arg.BatchEndDate,
		arg.BatchStatus)
	return err
}

const updateInstructorPrograms = `-- name: UpdateInstructorPrograms :exec
UPDATE bootcamp.instructor_programs
SET inpro_entity_id = $2,
	inpro_emp_entity_id = $3,
	inpro_modified_date = Now()
WHERE batch_id = $1
`

func (q *Queries) UpdateInstructorPrograms(ctx context.Context, arg CreateInstructorProgramsParams) error {
	_, err := q.db.ExecContext(ctx, updateInstructorPrograms,
		arg.BatchID,
		arg.InproEntityID,
		arg.InproEmpEntityID)
	return err
}

const deleteBatch = `-- name: DeleteBatch :exec
DELETE FROM bootcamp.batch
WHERE batch_id = $1
`

func (q *Queries) DeleteBatch(ctx context.Context, batchID int32) error {
	_, err := q.db.ExecContext(ctx, deleteBatch, batchID)
	return err
}

const deleteInstructorPrograms = `-- name: DeleteInstructorPrograms :exec
DELETE FROM bootcamp.instructor_programs
WHERE batch_id = $1
`

func (q *Queries) DeleteInstructorPrograms(ctx context.Context, batchID int32) error {
	_, err := q.db.ExecContext(ctx, deleteInstructorPrograms, batchID)
	return err
}

const deleteBatchTrainee = `-- name: DeleteBatchTrainee :exec
DELETE FROM bootcamp.batch_trainee
WHERE batr_batch_id = $1
`

func (q *Queries) DeleteBatchTrainee(ctx context.Context, batchID int32) error {
	_, err := q.db.ExecContext(ctx, deleteBatchTrainee, batchID)
	return err
}

const deleteBatchTrainee2 = `-- name: DeleteBatchTrainee2 :exec
DELETE FROM bootcamp.batch_trainee
WHERE batr_trainee_entity_id = $1 AND batr_batch_id = $2
`

func (q *Queries) DeleteBatchTrainee2(ctx context.Context, batrTraineeEntityID int32, batrBatchID int32) error {
	_, err := q.db.ExecContext(ctx, deleteBatchTrainee2, batrTraineeEntityID, batrBatchID)
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
