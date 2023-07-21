package dbContext

import (
	"context"
	"time"

	curi "codeid.revampacademy/models"
)

const createsection_detail = `-- name: Createsection_detail :one

INSERT INTO curriculum.section_detail (secd_id, 
secd_title, 
secd_preview, 
secd_score, 
secd_note, 
secd_minute, 
secd_modified_date, 
secd_sect_id)
VALUES($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING secd_id
`

type Createsection_detailParams struct {
	SecdID           int32     `db:"secd_id" json:"secdId"`
	SecdTitle        string    `db:"secd_title" json:"secdTitle"`
	SecdPreview      string    `db:"secd_preview" json:"secdPreview"`
	SecdScore        int32     `db:"secd_score" json:"secdScore"`
	SecdNote         string    `db:"secd_note" json:"secdNote"`
	SecdMinute       int32     `db:"secd_minute" json:"secdMinute"`
	SecdModifiedDate time.Time `db:"secd_modified_date" json:"secdModifiedDate"`
	SecdSectID       int32     `db:"secd_sect_id" json:"secdSectId"`
}

func (q *Queries) Createsection_detail(ctx context.Context, arg Createsection_detailParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createsection_detail,
		arg.SecdID,
		arg.SecdTitle,
		arg.SecdPreview,
		arg.SecdScore,
		arg.SecdNote,
		arg.SecdMinute,
		arg.SecdModifiedDate,
		arg.SecdSectID,
	)
	var secd_id int32
	err := row.Scan(&secd_id)
	return secd_id, err
}

const deletesection_detail = `-- name: Deletesection_detail :exec
DELETE FROM curriculum.section_detail
WHERE secd_id = $1
`

func (q *Queries) Deletesection_detail(ctx context.Context, secdID int32) error {
	_, err := q.db.ExecContext(ctx, deletesection_detail, secdID)
	return err
}

const getsection_detail = `-- name: Getsection_detail :one

SELECT secd_id, secd_title, secd_preview, secd_score, secd_note, secd_minute, secd_modified_date, secd_sect_id FROM curriculum.section_detail
WHERE secd_id = $1
`

// curriculum.section_detail
func (q *Queries) Getsection_detail(ctx context.Context, secdID int32) (curi.CurriculumSectionDetail, error) {
	row := q.db.QueryRowContext(ctx, getsection_detail, secdID)
	var i curi.CurriculumSectionDetail
	err := row.Scan(
		&i.SecdID,
		&i.SecdTitle,
		&i.SecdPreview,
		&i.SecdScore,
		&i.SecdNote,
		&i.SecdMinute,
		&i.SecdModifiedDate,
		&i.SecdSectID,
	)
	return i, err
}

const listsection_detail = `-- name: Listsection_detail :many
SELECT secd_id, secd_title, secd_preview, secd_score, secd_note, secd_minute, secd_modified_date, secd_sect_id FROM curriculum.section_detail
ORDER BY secd_title
`

func (q *Queries) Listsection_detail(ctx context.Context) ([]curi.CurriculumSectionDetail, error) {
	rows, err := q.db.QueryContext(ctx, listsection_detail)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []curi.CurriculumSectionDetail
	for rows.Next() {
		var i curi.CurriculumSectionDetail
		if err := rows.Scan(
			&i.SecdID,
			&i.SecdTitle,
			&i.SecdPreview,
			&i.SecdScore,
			&i.SecdNote,
			&i.SecdMinute,
			&i.SecdModifiedDate,
			&i.SecdSectID,
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

const updatesection_detail = `-- name: Updatesection_detail :exec
UPDATE curriculum.section_detail
  set secd_title = $2,
  secd_preview = $3
WHERE secd_id = $1
`

type Updatesection_detailParams struct {
	SecdID      int32  `db:"secd_id" json:"secdId"`
	SecdTitle   string `db:"secd_title" json:"secdTitle"`
	SecdPreview string `db:"secd_preview" json:"secdPreview"`
}

func (q *Queries) Updatesection_detail(ctx context.Context, arg Updatesection_detailParams) error {
	_, err := q.db.ExecContext(ctx, updatesection_detail, arg.SecdID, arg.SecdTitle, arg.SecdPreview)
	return err
}
