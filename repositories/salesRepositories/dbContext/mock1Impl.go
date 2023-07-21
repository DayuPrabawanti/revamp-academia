package dbContext

import (
	"context"
)

type CreateprogramEntityParams struct {
	ProgTitle        string `db:"prog_title" json:"progTitle"`
	ProgHeadline     string `db:"prog_headline" json:"progHeadline"`
	ProgLearningType string `db:"prog_learning_type" json:"progLearningType"`
	ProgImage        string `db:"prog_image" json:"progImage"`
	ProgPrice        int32  `db:"prog_price" json:"progPrice"`
	ProgDuration     int32  `db:"prog_duration" json:"progDuration"`
}

const getProgramEntity = `-- name: getProgramEntity :one
SELECT prog_title, prog_headline, prog_learning_type, prog_image, prog_price, prog_duration FROM curriculum.program_entity
WHERE prog_learning_type = $1
`

func (q *Queries) GetProgramEntity(ctx context.Context, nama string) (CreateprogramEntityParams, error) {
	row := q.db.QueryRowContext(ctx, getProgramEntity, nama)
	var i CreateprogramEntityParams
	err := row.Scan(
		&i.ProgTitle,
		&i.ProgHeadline,
		&i.ProgLearningType,
		&i.ProgImage,
		&i.ProgPrice,
		&i.ProgDuration,
	)
	return i, err
}

const listprogram_entity = `-- name: Listprogram_entity :many
select prog_title,prog_headline,prog_learning_type,prog_image,prog_price,prog_duration from curriculum.program_entity 
where prog_learning_type =$1
`

func (q *Queries) Listprogram_entity(ctx context.Context, nama string) ([]CreateprogramEntityParams, error) {
	rows, err := q.db.QueryContext(ctx, listprogram_entity, nama)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CreateprogramEntityParams
	for rows.Next() {
		var i CreateprogramEntityParams
		if err := rows.Scan(
			&i.ProgTitle,
			&i.ProgHeadline,
			&i.ProgLearningType,
			&i.ProgImage,
			&i.ProgPrice,
			&i.ProgDuration,
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
