package dbContext

import (
	"context"
	"time"

	curi "codeid.revampacademy/models/curriculum"
)

const createsections = `-- name: Createsections :one

INSERT INTO curriculum.sections (sect_id, 
sect_prog_entity_id, 
sect_title, 
sect_description, 
sect_total_section, 
sect_total_lecture, 
sect_total_minute, 
sect_modified_date)
VALUES($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING sect_id
`

type CreatesectionsParams struct {
	SectID           int32     `db:"sect_id" json:"sectId"`
	SectProgEntityID int32     `db:"sect_prog_entity_id" json:"sectProgEntityId"`
	SectTitle        string    `db:"sect_title" json:"sectTitle"`
	SectDescription  string    `db:"sect_description" json:"sectDescription"`
	SectTotalSection int32     `db:"sect_total_section" json:"sectTotalSection"`
	SectTotalLecture int32     `db:"sect_total_lecture" json:"sectTotalLecture"`
	SectTotalMinute  int32     `db:"sect_total_minute" json:"sectTotalMinute"`
	SectModifiedDate time.Time `db:"sect_modified_date" json:"sectModifiedDate"`
}

func (q *Queries) Createsections(ctx context.Context, arg CreatesectionsParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createsections,
		arg.SectID,
		arg.SectProgEntityID,
		arg.SectTitle,
		arg.SectDescription,
		arg.SectTotalSection,
		arg.SectTotalLecture,
		arg.SectTotalMinute,
		arg.SectModifiedDate,
	)
	var sect_id int32
	err := row.Scan(&sect_id)
	return sect_id, err
}

const deletesections = `-- name: Deletesections :exec
DELETE FROM curriculum.sections
WHERE sect_id = $1
`

func (q *Queries) Deletesections(ctx context.Context, sectID int32) error {
	_, err := q.db.ExecContext(ctx, deletesections, sectID)
	return err
}

const getsections = `-- name: Getsections :one

SELECT sect_id, sect_prog_entity_id, sect_title, sect_description, sect_total_section, sect_total_lecture, sect_total_minute, sect_modified_date FROM curriculum.sections
WHERE sect_id = $1
`

// curriculum.sections
func (q *Queries) Getsections(ctx context.Context, sectID int32) (curi.CurriculumSection, error) {
	row := q.db.QueryRowContext(ctx, getsections, sectID)
	var i curi.CurriculumSection
	err := row.Scan(
		&i.SectID,
		&i.SectProgEntityID,
		&i.SectTitle,
		&i.SectDescription,
		&i.SectTotalSection,
		&i.SectTotalLecture,
		&i.SectTotalMinute,
		&i.SectModifiedDate,
	)
	return i, err
}

const listsections = `-- name: Listsections :many
SELECT sect_id, sect_prog_entity_id, sect_title, sect_description, sect_total_section, sect_total_lecture, sect_total_minute, sect_modified_date FROM curriculum.sections
`

func (q *Queries) Listsections(ctx context.Context) ([]curi.CurriculumSection, error) {
	rows, err := q.db.QueryContext(ctx, listsections)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []curi.CurriculumSection
	for rows.Next() {
		var i curi.CurriculumSection
		if err := rows.Scan(
			&i.SectID,
			&i.SectProgEntityID,
			&i.SectTitle,
			&i.SectDescription,
			&i.SectTotalSection,
			&i.SectTotalLecture,
			&i.SectTotalMinute,
			&i.SectModifiedDate,
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

const updatesections = `-- name: Updatesections :exec
UPDATE curriculum.sections
  set sect_title = $2,
  sect_description = $3
WHERE sect_id = $1
`

type UpdatesectionsParams struct {
	SectID          int32  `db:"sect_id" json:"sectId"`
	SectTitle       string `db:"sect_title" json:"sectTitle"`
	SectDescription string `db:"sect_description" json:"sectDescription"`
}

func (q *Queries) Updatesections(ctx context.Context, arg UpdatesectionsParams) error {
	_, err := q.db.ExecContext(ctx, updatesections, arg.SectID, arg.SectTitle, arg.SectDescription)
	return err
}
