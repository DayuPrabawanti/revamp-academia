package dbContext

import (
	"context"
	"net/http"
	"time"

	"codeid.revampacademy/models"
)

const listBatchTrinee = `-- name: ListBatchTrinee :many
SELECT batr_id, batr_status, batr_certificated, batre_certificate_link, batr_access_token, batr_access_grant, batr_review, batr_total_score, batr_modified_date, batr_trainee_entity_id, batr_batch_id FROM bootcamp.batch_trainee
ORDER BY batr_id
`

func (q *Queries) ListBatchTrinee(ctx context.Context) ([]models.BootcampBatchTrainee, error) {
	rows, err := q.db.QueryContext(ctx, listBatchTrinee)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.BootcampBatchTrainee
	for rows.Next() {
		var i models.BootcampBatchTrainee
		if err := rows.Scan(
			&i.BatrID,
			&i.BatrStatus,
			&i.BatrCertificated,
			&i.BatreCertificateLink,
			&i.BatrAccessToken,
			&i.BatrAccessGrant,
			&i.BatrReview,
			&i.BatrTotalScore,
			&i.BatrModifiedDate,
			&i.BatrTraineeEntityID,
			&i.BatrBatchID,
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

const getBatchTrainee = `-- name: GetBatchTrainee :one
SELECT batr_id, batr_status, batr_certificated, batre_certificate_link, batr_access_token, batr_access_grant, batr_review, batr_total_score, batr_modified_date, batr_trainee_entity_id, batr_batch_id FROM bootcamp.batch_trainee
WHERE batr_id = $1
`

func (q *Queries) GetBatchTrainee(ctx context.Context, batrID int32) (models.BootcampBatchTrainee, error) {
	row := q.db.QueryRowContext(ctx, getBatchTrainee, batrID)
	var i models.BootcampBatchTrainee
	err := row.Scan(
		&i.BatrID,
		&i.BatrStatus,
		&i.BatrCertificated,
		&i.BatreCertificateLink,
		&i.BatrAccessToken,
		&i.BatrAccessGrant,
		&i.BatrReview,
		&i.BatrTotalScore,
		&i.BatrModifiedDate,
		&i.BatrTraineeEntityID,
		&i.BatrBatchID,
	)
	return i, err
}

const createBatchTrainee = `-- name: CreateBatchTrainee :one
INSERT INTO bootcamp.batch_trainee
(batr_id, batr_status, batr_certificated, batre_certificate_link, batr_access_token, batr_access_grant, batr_review, batr_total_score, batr_modified_date, batr_trainee_entity_id, batr_batch_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING batr_id
`

type CreateBatchTraineeParams struct {
	BatrID               int32     `db:"batr_id" json:"batrId"`
	BatrStatus           string    `db:"batr_status" json:"batrStatus"`
	BatrCertificated     string    `db:"batr_certificated" json:"batrCertificated"`
	BatreCertificateLink string    `db:"batre_certificate_link" json:"batreCertificateLink"`
	BatrAccessToken      string    `db:"batr_access_token" json:"batrAccessToken"`
	BatrAccessGrant      string    `db:"batr_access_grant" json:"batrAccessGrant"`
	BatrReview           string    `db:"batr_review" json:"batrReview"`
	BatrTotalScore       int32     `db:"batr_total_score" json:"batrTotalScore"`
	BatrModifiedDate     time.Time `db:"batr_modified_date" json:"batrModifiedDate"`
	BatrTraineeEntityID  int32     `db:"batr_trainee_entity_id" json:"batrTraineeEntityId"`
	BatrBatchID          int32     `db:"batr_batch_id" json:"batrBatchId"`
}

func (q *Queries) CreateBatchTrainee(ctx context.Context, arg CreateBatchTraineeParams) (*models.BootcampBatchTrainee, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createBatchTrainee,
		arg.BatrID,
		arg.BatrStatus,
		arg.BatrCertificated,
		arg.BatreCertificateLink,
		arg.BatrAccessToken,
		arg.BatrAccessGrant,
		arg.BatrReview,
		arg.BatrTotalScore,
		arg.BatrModifiedDate,
		arg.BatrTraineeEntityID,
		arg.BatrBatchID,
	)
	i := models.BootcampBatchTrainee{}
	err := row.Scan(
		&i.BatrID,
		&i.BatrStatus,
		&i.BatrCertificated,
		&i.BatreCertificateLink,
		&i.BatrAccessToken,
		&i.BatrAccessGrant,
		&i.BatrReview,
		&i.BatrTotalScore,
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
		BatrID:               i.BatrID,
		BatrStatus:           i.BatrStatus,
		BatrCertificated:     i.BatrCertificated,
		BatreCertificateLink: i.BatreCertificateLink,
		BatrAccessToken:      i.BatrAccessToken,
		BatrAccessGrant:      i.BatrAccessGrant,
		BatrReview:           i.BatrReview,
		BatrTotalScore:       i.BatrTotalScore,
		BatrModifiedDate:     i.BatrModifiedDate,
		BatrTraineeEntityID:  i.BatrTraineeEntityID,
		BatrBatchID:          i.BatrBatchID,
	}, nil
}

const updateBatchTrainee = `-- name: UpdateBatchTrainee :exec
UPDATE bootcamp.batch_trainee
SET batr_status = $2,
    batr_review = $3
WHERE batr_id = $1
`

type UpdateBatchTraineeParams struct {
	BatrID     int32  `db:"batr_id" json:"batrId"`
	BatrStatus string `db:"batr_status" json:"batrStatus"`
	BatrReview string `db:"batr_review" json:"batrReview"`
}

func (q *Queries) UpdateBatchTrainee(ctx context.Context, arg UpdateBatchTraineeParams) error {
	_, err := q.db.ExecContext(ctx, updateBatchTrainee, arg.BatrID, arg.BatrStatus, arg.BatrReview)
	return err
}

const deleteBatchTrainee = `-- name: DeleteBatchTrainee :exec
DELETE FROM bootcamp.batch_trainee
WHERE batr_id = $1
`

func (q *Queries) DeleteBatchTrainee(ctx context.Context, batrID int32) error {
	_, err := q.db.ExecContext(ctx, deleteBatchTrainee, batrID)
	return err
}
