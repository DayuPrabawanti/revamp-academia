package dbContext

import (
	"context"
	"net/http"

	"codeid.revampacademy/models"
)

type BootcampBatchEvaluationMockup struct {
	UserEntityID        int32  `db:"user_entity_id" json:"userEntityId"`
	UserFullname        string `json:"fullname"`
	UserPhoto           string `db:"user_photo" json:"userPhoto"`
	ProgTitle           string `db:"prog_title" json:"progTitle"`
	BatchID             int32  `db:"batch_id" json:"batchId"`
	BatchEntityID       int32  `db:"batch_entity_id" json:"batchEntityId"`
	BatchName           string `db:"batch_name" json:"batchName"`
	BatrStatus          string `db:"batr_status" json:"batrStatus"`
	BtevSkor            int32  `db:"btev_skor" json:"btevSkor"`
	BtevTraineeEntityID int32  `db:"btev_trainee_entity_id" json:"btevTraineeEntityId"`
}

type BootcampBatchTraineeReview struct {
	UserEntityID int32  `db:"user_entity_id" json:"userEntityId"`
	UserFullname string `json:"fullname"`
	BatrID       int32  `db:"batr_id" json:"batrId"`
	BatrStatus   string `db:"batr_status" json:"batrStatus"`
	BatrReview   string `db:"batr_review" json:"batrReview"`
}

const getBootcampBatchEvaluations = `-- name: GetBootcampBatchEvaluations :many
SELECT
	user_entity_id,
	btev_batch_id,

    batch_name,
	prog_title, 
    
    user_photo, 
    CONCAT (user_first_name, user_last_name) AS fullname, 
    batr_status, 
    SUM(btev_skor) AS total_skor,

	batch_entity_id,
	btev_trainee_entity_id
FROM 
    bootcamp.batch_trainee_evaluation
JOIN 
    bootcamp.batch_trainee 
ON 
    btev_trainee_entity_id = batr_trainee_entity_id
JOIN 
    users.users 
ON 
    batr_trainee_entity_id = user_entity_id
JOIN 
    bootcamp.batch
ON 
    batch_entity_id = user_entity_id
JOIN 
    curriculum.program_entity
ON 
    batch_entity_id = prog_entity_id
WHERE 
    batch_id = $1
GROUP BY 
	user_entity_id,
	btev_batch_id, 

    batch_name,
    prog_title, 

    user_photo,
    batr_status,
	
	batch_entity_id,
	btev_trainee_entity_id
ORDER BY
    btev_batch_id;
`

func (q *Queries) GetBootcampBatchEvaluation(ctx context.Context, batchID int32) ([]BootcampBatchEvaluationMockup, error) {
	rows, err := q.db.QueryContext(ctx, getBootcampBatchEvaluations, batchID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bootcampEvaluations []BootcampBatchEvaluationMockup

	for rows.Next() {
		var i BootcampBatchEvaluationMockup

		err := rows.Scan(
			&i.UserEntityID,
			&i.BatchID,
			&i.BatchName,
			&i.ProgTitle,
			&i.UserPhoto,
			&i.UserFullname,
			&i.BatrStatus,
			&i.BtevSkor,
			&i.BatchEntityID,
			&i.BtevTraineeEntityID,
		)

		if err != nil {
			return nil, err
		}
		bootcampEvaluations = append(bootcampEvaluations, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return bootcampEvaluations, nil
}

const listProgramEntity = `-- name: ListProgramEntity :many
SELECT prog_entity_id, prog_title, prog_headline, prog_type, prog_learning_type, prog_rating, prog_total_trainee, prog_image, prog_best_seller, prog_price, prog_language, prog_modified_date, prog_duration, prog_duration_type, prog_tag_skill, prog_city_id, prog_cate_id, prog_created_by, prog_status 
	FROM curriculum.program_entity
	ORDER BY prog_title
`

func (q *Queries) GetBatchTraineeReview(ctx context.Context, userEntityID int32) ([]BootcampBatchTraineeReview, error) {

	rows, err := q.db.QueryContext(ctx, getBatchTraineeReviews, userEntityID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bootcampReviews []BootcampBatchTraineeReview

	for rows.Next() {
		var i BootcampBatchTraineeReview

		err := rows.Scan(
			&i.UserEntityID,
			&i.UserFullname,
			&i.BatrID,
			&i.BatrStatus,
			&i.BatrReview,
		)

		if err != nil {
			return nil, err
		}
		bootcampReviews = append(bootcampReviews, i)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return bootcampReviews, nil
}

const listUsers = `-- name: ListUsers :many
SELECT user_entity_id, user_name, user_password, user_first_name, user_last_name, user_birth_date, user_email_promotion, user_demographic, user_modified_date, user_photo, user_current_role FROM users.users
ORDER BY user_name
`

type CreateBatchTraineeReviewParams struct {
	BatrID              int32  `db:"batr_id" json:"batrId"`
	BatrBatchID         int32  `json:"batrBatchID"`
	BatrTraineeEntityID int32  `json:"batrTraineeEntityId"`
	BatrStatus          string `json:"batrStatus"`
	BatrReview          string `json:"batrReview"`
}

func (q *Queries) CreateBatchTraineeReview(ctx context.Context, arg CreateBatchTraineeReviewParams) (*BootcampBatchTraineeReview, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createBatchTraineeReview,
		arg.BatrBatchID,
		arg.BatrTraineeEntityID,
		arg.BatrStatus,
		arg.BatrReview,
	)

	i := BootcampBatchTraineeReview{}
	err := row.Scan(
		&i.BatrID,
		&i.UserFullname,
	)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &BootcampBatchTraineeReview{
		UserFullname: i.UserFullname,
		BatrID:       i.BatrID,
	}, nil
}

const updateBatchTraineeReview = `-- name: UpdateBatchTraineeReview :exec
UPDATE 
	bootcamp.batch_trainee AS bt
SET
	batr_status = $1, 
	batr_review = $2
FROM 
	users.users AS u
WHERE 
	bt.batr_trainee_entity_id = $3 
AND 
	u.user_entity_id = bt.batr_trainee_entity_id
RETURNING 
	bt.batr_id,
	CONCAT(u.user_first_name, ' ', u.user_last_name) AS user_fullname
`

type UpdateBatchTraineeReviewParams struct {
	BatrStatus string `json:"batrStatus"`
	BatrReview string `json:"batrReview"`
	BatrID     int32  `json:"batrID"`
}

func (q *Queries) UpdateBatchTraineeReview(ctx context.Context, arg UpdateBatchTraineeReviewParams) (*BootcampBatchTraineeReview, error) {
	row := q.db.QueryRowContext(ctx, updateBatchTraineeReview,
		arg.BatrStatus,
		arg.BatrReview,
		arg.BatrID)

	var i BootcampBatchTraineeReview
	err := row.Scan(
		&i.UserEntityID,
		&i.UserFullname,
	)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

const deleteBatchTraineeReview = `-- name: DeleteBatchTraineeReview :exec
DELETE FROM bootcamp.batch_trainee
WHERE batr_id = $1
`

func (q *Queries) DeleteBatchTraineeReview(ctx context.Context, batrID int32) error {
	_, err := q.db.ExecContext(ctx, deleteBatchTraineeReview, batrID)
	return err
}
