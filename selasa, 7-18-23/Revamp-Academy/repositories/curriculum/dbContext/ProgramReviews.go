package dbContext

import (
	"context"
	"time"

	curi "codeid.revampacademy/models/curriculum"
)

const createprogram_reviews = `-- name: Createprogram_reviews :one

INSERT INTO curriculum.program_reviews (prow_user_entity_id, 
prow_prog_entity_id, 
prow_review, 
prow_rating, 
prow_modified_date) 

VALUES($1,$2,$3,$4,$5)
RETURNING prow_user_entity_id
`

type Createprogram_reviewsParams struct {
	ProwUserEntityID int32     `db:"prow_user_entity_id" json:"prowUserEntityId"`
	ProwProgEntityID int32     `db:"prow_prog_entity_id" json:"prowProgEntityId"`
	ProwReview       string    `db:"prow_review" json:"prowReview"`
	ProwRating       int32     `db:"prow_rating" json:"prowRating"`
	ProwModifiedDate time.Time `db:"prow_modified_date" json:"prowModifiedDate"`
}

func (q *Queries) Createprogram_reviews(ctx context.Context, arg Createprogram_reviewsParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createprogram_reviews,
		arg.ProwUserEntityID,
		arg.ProwProgEntityID,
		arg.ProwReview,
		arg.ProwRating,
		arg.ProwModifiedDate,
	)
	var prow_user_entity_id int32
	err := row.Scan(&prow_user_entity_id)
	return prow_user_entity_id, err
}

const deleteprogram_reviews = `-- name: Deleteprogram_reviews :exec
DELETE FROM curriculum.program_reviews
WHERE prow_user_entity_id = $1
`

func (q *Queries) Deleteprogram_reviews(ctx context.Context, prowUserEntityID int32) error {
	_, err := q.db.ExecContext(ctx, deleteprogram_reviews, prowUserEntityID)
	return err
}

const getprogram_reviews = `-- name: Getprogram_reviews :one

SELECT prow_user_entity_id, prow_prog_entity_id, prow_review, prow_rating, prow_modified_date FROM curriculum.program_reviews
WHERE prow_user_entity_id = $1
`

// curriculum.program_reviews
func (q *Queries) Getprogram_reviews(ctx context.Context, prowUserEntityID int32) (curi.CurriculumProgramReview, error) {
	row := q.db.QueryRowContext(ctx, getprogram_reviews, prowUserEntityID)
	var i curi.CurriculumProgramReview
	err := row.Scan(
		&i.ProwUserEntityID,
		&i.ProwProgEntityID,
		&i.ProwReview,
		&i.ProwRating,
		&i.ProwModifiedDate,
	)
	return i, err
}

const listprogram_reviews = `-- name: Listprogram_reviews :many
SELECT prow_user_entity_id, prow_prog_entity_id, prow_review, prow_rating, prow_modified_date FROM curriculum.program_reviews
ORDER BY prow_rating
`

func (q *Queries) Listprogram_reviews(ctx context.Context) ([]curi.CurriculumProgramReview, error) {
	rows, err := q.db.QueryContext(ctx, listprogram_reviews)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []curi.CurriculumProgramReview
	for rows.Next() {
		var i curi.CurriculumProgramReview
		if err := rows.Scan(
			&i.ProwUserEntityID,
			&i.ProwProgEntityID,
			&i.ProwReview,
			&i.ProwRating,
			&i.ProwModifiedDate,
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

const updateprogram_reviews = `-- name: Updateprogram_reviews :exec
UPDATE curriculum.program_reviews
  set prow_review = $2,
  prow_rating = $3
WHERE prow_user_entity_id = $1
`

type Updateprogram_reviewsParams struct {
	ProwUserEntityID int32  `db:"prow_user_entity_id" json:"prowUserEntityId"`
	ProwReview       string `db:"prow_review" json:"prowReview"`
	ProwRating       int32  `db:"prow_rating" json:"prowRating"`
}

func (q *Queries) Updateprogram_reviews(ctx context.Context, arg Updateprogram_reviewsParams) error {
	_, err := q.db.ExecContext(ctx, updateprogram_reviews, arg.ProwUserEntityID, arg.ProwReview, arg.ProwRating)
	return err
}
