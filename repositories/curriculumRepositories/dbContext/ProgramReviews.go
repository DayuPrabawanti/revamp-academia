package dbContext

import (
	"context"

	models "codeid.revampacademy/models"
)

const listProgReviews = `-- name: ListProgReviews :many
SELECT prow_user_entity_id, prow_prog_entity_id, prow_review, prow_rating, prow_modified_date 
FROM curriculum.program_reviews
ORDER BY prow_user_entity_id
`

func (q *Queries) ListProgReviews(ctx context.Context) ([]models.CurriculumProgramReview, error) {
	rows, err := q.db.QueryContext(ctx, listProgReviews)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.CurriculumProgramReview
	for rows.Next() {
		var i models.CurriculumProgramReview
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

const getProgramReviews = `-- name: GetProgramReviews :one
SELECT prow_user_entity_id, prow_prog_entity_id, prow_review, prow_rating, prow_modified_date 
FROM curriculum.program_reviews
	WHERE prow_user_entity_id = $1
`

func (q *Queries) GetProgramReviews(ctx context.Context, prowUserEntityId int16) (models.CurriculumProgramReview, error) {
	row := q.db.QueryRowContext(ctx, getProgramReviews, prowUserEntityId)
	var i models.CurriculumProgramReview
	err := row.Scan(
		&i.ProwUserEntityID,
		&i.ProwProgEntityID,
		&i.ProwReview,
		&i.ProwRating,
		&i.ProwModifiedDate,
	)
	return i, err
}
