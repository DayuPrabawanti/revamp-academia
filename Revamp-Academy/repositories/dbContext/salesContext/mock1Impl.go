package salesContext

import (
	"context"
	"database/sql"
)

type CreateprogramEntityParams struct {
	ProgTitle        string         `db:"prog_title" json:"progTitle"`
	ProgHeadline     string         `db:"prog_headline" json:"progHeadline"`
	ProgLearningType string         `db:"prog_learning_type" json:"progLearningType"`
	ProgImage        sql.NullString `db:"prog_image" json:"progImage"`
	ProgPrice        sql.NullInt32  `db:"prog_price" json:"progPrice"`
	ProgDuration     sql.NullInt32  `db:"prog_duration" json:"progDuration"`
}

const getProgramEntity = `-- name: Getprogram_entity :one
SELECT prog_title, prog_headline, prog_learning_type, prog_image, prog_price, prog_duration FROM curriculum.program_entity
WHERE prog_title = $1
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
