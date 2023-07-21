package dbcontext

import (
	"context"

	models "codeid.revampacademy/db-generator/gen"
)

const listProgEntityDesc = `-- name: ListProgEntityDesc :many
SELECT pred_prog_entity_id, pred_item_learning, pred_item_include, pred_requirement, pred_description, pred_target_level 
FROM curriculum.program_entity_description
ORDER BY pred_item_learning
`

func (q *Queries) ListProgEntityDesc(ctx context.Context) ([]models.CurriculumProgramEntityDescription, error) {
	rows, err := q.db.QueryContext(ctx, listProgEntityDesc)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.CurriculumProgramEntityDescription
	for rows.Next() {
		var i models.CurriculumProgramEntityDescription
		if err := rows.Scan(
			&i.PredProgEntityID,
			&i.PredItemLearning,
			&i.PredItemInclude,
			&i.PredRequirement,
			&i.PredDescription,
			&i.PredTargetLevel,
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

const getProgEntityDesc = `-- name: GetProgEntityDesc :one
SELECT pred_prog_entity_id, pred_item_learning, pred_item_include, pred_requirement, pred_description, pred_target_level 
FROM curriculum.program_entity_description
	WHERE pred_prog_entity_id = $1
`

func (q *Queries) GetProgEntityDesc(ctx context.Context, predProgEntityID int16) (models.CurriculumProgramEntityDescription, error) {
	row := q.db.QueryRowContext(ctx, getProgEntityDesc, predProgEntityID)
	var i models.CurriculumProgramEntityDescription
	err := row.Scan(
		&i.PredProgEntityID,
		&i.PredItemLearning,
		&i.PredItemInclude,
		&i.PredRequirement,
		&i.PredDescription,
	)
	return i, err
}
