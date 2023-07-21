package dbContext

import (
	"context"

	"codeid.revampacademy/models"
)

const listCategories = `-- name: ListCategories :many
SELECT joca_id, joca_name, joca_modified_date FROM jobHire.job_category
ORDER BY joca_name
`

func (q *Queries) ListCategories(ctx context.Context) ([]models.JobhireJobCategory, error) {
	rows, err := q.db.QueryContext(ctx, listCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.JobhireJobCategory
	for rows.Next() {
		var i models.JobhireJobCategory
		if err := rows.Scan(&i.JocaID, &i.JocaName, &i.JocaModifiedDate); err != nil {
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
