package dbContext

import (
	"context"

	"codeid.revampacademy/models"
)

const listProgramApplyProgresses = `-- name: ListProgramApplyProgresses :many
SELECT parog_id, parog_user_entity_id, parog_prog_entity_id, parog_action_date, parog_modified_date, parog_comment, parog_progress_name, parog_emp_entity_id, parog_status FROM bootcamp.program_apply_progress
ORDER BY parog_id
`

func (q *Queries) ListProgramApplyProgresses(ctx context.Context) ([]models.BootcampProgramApplyProgress, error) {
	rows, err := q.db.QueryContext(ctx, listProgramApplyProgresses)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.BootcampProgramApplyProgress
	for rows.Next() {
		var i models.BootcampProgramApplyProgress
		if err := rows.Scan(
			&i.ParogID,
			&i.ParogUserEntityID,
			&i.ParogProgEntityID,
			&i.ParogActionDate,
			&i.ParogModifiedDate,
			&i.ParogComment,
			&i.ParogProgressName,
			&i.ParogEmpEntityID,
			&i.ParogStatus,
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
