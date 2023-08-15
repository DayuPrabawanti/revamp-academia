package dbContext

import (
	"context"

	"codeid.revampacademy/models"
)

const listProgramApplies = `-- name: ListProgramApplies :many
SELECT prap_user_entity_id, prap_prog_entity_id, prap_test_score, prap_gpa, prap_iq_test, prap_review, prap_modified_date, prap_status FROM bootcamp.program_apply
ORDER BY prap_user_entity_id
`

func (q *Queries) ListProgramApplies(ctx context.Context) ([]models.BootcampProgramApply, error) {
	rows, err := q.db.QueryContext(ctx, listProgramApplies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.BootcampProgramApply
	for rows.Next() {
		var i models.BootcampProgramApply
		if err := rows.Scan(
			&i.PrapUserEntityID,
			&i.PrapProgEntityID,
			&i.PrapTestScore,
			&i.PrapGpa,
			&i.PrapIqTest,
			&i.PrapReview,
			&i.PrapModifiedDate,
			&i.PrapStatus,
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
