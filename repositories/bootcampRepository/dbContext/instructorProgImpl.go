package dbContext

import (
	"context"

	"codeid.revampacademy/models"
)

const listInstructorPrograms = `-- name: ListInstructorPrograms :many
SELECT batch_id, inpro_entity_id, inpro_emp_entity_id, inpro_modified_date FROM bootcamp.instructor_programs
ORDER BY batch_id
`

func (q *Queries) ListInstructorPrograms(ctx context.Context) ([]models.BootcampInstructorProgram, error) {
	rows, err := q.db.QueryContext(ctx, listInstructorPrograms)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.BootcampInstructorProgram
	for rows.Next() {
		var i models.BootcampInstructorProgram
		if err := rows.Scan(
			&i.BatchID,
			&i.InproEntityID,
			&i.InproEmpEntityID,
			&i.InproModifiedDate,
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
