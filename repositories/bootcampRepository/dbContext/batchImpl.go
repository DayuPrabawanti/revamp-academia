package dbContext

import (
	"context"
	"net/http"
	"time"

	"codeid.revampacademy/models"
)

const listBatchs = `-- name: ListBatchs :many
SELECT batch_id, batch_entity_id, batch_name, batch_description, batch_start_date, batch_end_date, batch_reason, batch_type, batch_modified_date, batch_status, batch_pic_id FROM bootcamp.batch
ORDER BY batch_id
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

const listProgramEntity = `-- name: ListProgramEntity :many
SELECT prog_entity_id, prog_title, prog_headline, prog_type, prog_learning_type, prog_rating, prog_total_trainee, prog_image, prog_best_seller, prog_price, prog_language, prog_modified_date, prog_duration, prog_duration_type, prog_tag_skill, prog_city_id, prog_cate_id, prog_created_by, prog_status 
	FROM curriculum.program_entity
	ORDER BY prog_title
`

func (q *Queries) ListProgramEntity(ctx context.Context) ([]models.CurriculumProgramEntity, error) {
	rows, err := q.db.QueryContext(ctx, listProgramEntity)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.CurriculumProgramEntity
	for rows.Next() {
		var i models.CurriculumProgramEntity
		if err := rows.Scan(
			&i.ProgEntityID,
			&i.ProgTitle,
			&i.ProgHeadline,
			&i.ProgType,
			&i.ProgLearningType,
			&i.ProgRating,
			&i.ProgTotalTrainee,
			&i.ProgImage,
			&i.ProgBestSeller,
			&i.ProgPrice,
			&i.ProgLanguage,
			&i.ProgModifiedDate,
			&i.ProgDuration,
			&i.ProgDurationType,
			&i.ProgTagSkill,
			&i.ProgCityID,
			&i.ProgCateID,
			&i.ProgCreatedBy,
			&i.ProgStatus,
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

const listUsers = `-- name: ListUsers :many
SELECT user_entity_id, user_name, user_password, user_first_name, user_last_name, user_birth_date, user_email_promotion, user_demographic, user_modified_date, user_photo, user_current_role FROM users.users
ORDER BY user_name
`

func (q *Queries) ListUsers(ctx context.Context) ([]models.UsersUser, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.UsersUser
	for rows.Next() {
		var i models.UsersUser
		if err := rows.Scan(
			&i.UserEntityID,
			&i.UserName,
			&i.UserPassword,
			&i.UserFirstName,
			&i.UserLastName,
			&i.UserBirthDate,
			&i.UserEmailPromotion,
			&i.UserDemographic,
			&i.UserModifiedDate,
			&i.UserPhoto,
			&i.UserCurrentRole,
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
