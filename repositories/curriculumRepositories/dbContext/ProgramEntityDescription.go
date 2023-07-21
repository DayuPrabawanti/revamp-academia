package dbContext

import (
	"context"
	"database/sql"

	curi "codeid.revampacademy/models"
)

const createprogram_entity_description = `-- name: Createprogram_entity_description :one

INSERT INTO curriculum.program_entity_description (pred_prog_entity_id, 
pred_item_learning, 
pred_item_include, 
pred_requirment, 
pred_description, 
pred_target_level) 
VALUES($1,$2,$3,$4,$5,$6)
RETURNING pred_prog_entity_id
`

type Createprogram_entity_descriptionParams struct {
	PredProgEntityID int32          `db:"pred_prog_entity_id" json:"predProgEntityId"`
	PredItemLearning sql.NullString `db:"pred_item_learning" json:"predItemLearning"`
	PredItemInclude  sql.NullString `db:"pred_item_include" json:"predItemInclude"`
	PredRequirment   sql.NullString `db:"pred_requirment" json:"predRequirment"`
	PredDescription  sql.NullString `db:"pred_description" json:"predDescription"`
	PredTargetLevel  sql.NullString `db:"pred_target_level" json:"predTargetLevel"`
}

func (q *Queries) Createprogram_entity_description(ctx context.Context, arg Createprogram_entity_descriptionParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createprogram_entity_description,
		arg.PredProgEntityID,
		arg.PredItemLearning,
		arg.PredItemInclude,
		arg.PredRequirment,
		arg.PredDescription,
		arg.PredTargetLevel,
	)
	var pred_prog_entity_id int32
	err := row.Scan(&pred_prog_entity_id)
	return pred_prog_entity_id, err
}

const deleteprogram_entity_description = `-- name: Deleteprogram_entity_description :exec
DELETE FROM curriculum.program_entity_description
WHERE pred_prog_entity_id = $1
`

func (q *Queries) Deleteprogram_entity_description(ctx context.Context, predProgEntityID int32) error {
	_, err := q.db.ExecContext(ctx, deleteprogram_entity_description, predProgEntityID)
	return err
}

const getprogram_entity_description = `-- name: Getprogram_entity_description :one

SELECT pred_prog_entity_id, pred_item_learning, pred_item_include, pred_requirment, pred_description, pred_target_level FROM curriculum.program_entity_description
WHERE pred_prog_entity_id = $1
`

// curriculum.program_entity_description
func (q *Queries) Getprogram_entity_description(ctx context.Context, predProgEntityID int32) (curi.CurriculumProgramEntityDescription, error) {
	row := q.db.QueryRowContext(ctx, getprogram_entity_description, predProgEntityID)
	var i curi.CurriculumProgramEntityDescription
	err := row.Scan(
		&i.PredProgEntityID,
		&i.PredItemLearning,
		&i.PredItemInclude,
		&i.PredRequirment,
		&i.PredDescription,
		&i.PredTargetLevel,
	)
	return i, err
}

const listprogram_entity_description = `-- name: Listprogram_entity_description :many
SELECT pred_prog_entity_id, pred_item_learning, pred_item_include, pred_requirment, pred_description, pred_target_level FROM curriculum.program_entity_description
`

func (q *Queries) Listprogram_entity_description(ctx context.Context) ([]curi.CurriculumProgramEntityDescription, error) {
	rows, err := q.db.QueryContext(ctx, listprogram_entity_description)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []curi.CurriculumProgramEntityDescription
	for rows.Next() {
		var i curi.CurriculumProgramEntityDescription
		if err := rows.Scan(
			&i.PredProgEntityID,
			&i.PredItemLearning,
			&i.PredItemInclude,
			&i.PredRequirment,
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

const updateprogram_entity_description = `-- name: Updateprogram_entity_description :exec
UPDATE curriculum.program_entity_description
  set pred_item_learning= $2,
  pred_item_include = $3
WHERE pred_prog_entity_id= $1
`

type Updateprogram_entity_descriptionParams struct {
	PredProgEntityID int32          `db:"pred_prog_entity_id" json:"predProgEntityId"`
	PredItemLearning sql.NullString `db:"pred_item_learning" json:"predItemLearning"`
	PredItemInclude  sql.NullString `db:"pred_item_include" json:"predItemInclude"`
}

func (q *Queries) Updateprogram_entity_description(ctx context.Context, arg Updateprogram_entity_descriptionParams) error {
	_, err := q.db.ExecContext(ctx, updateprogram_entity_description, arg.PredProgEntityID, arg.PredItemLearning, arg.PredItemInclude)
	return err
}
