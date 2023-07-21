package dbContext

import (
	"context"
	"net/http"
	"time"

	mod "codeid.revampacademy/models"
)

const createsections = `-- name: Createsections :many

INSERT INTO curriculum.sections (sect_id, 
sect_prog_entity_id, 
sect_title, 
sect_description, 
sect_total_section, 
sect_total_lecture, 
sect_total_minute, 
sect_modified_date)
VALUES($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING *
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

func (q *Queries) Createsections(ctx context.Context, arg CreatesectionsParams) (*mod.CurriculumSection, *mod.ResponseError) {
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
	i := mod.CurriculumSection{}
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
	if err != nil {
		return nil, &mod.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &mod.CurriculumSection{
		SectID:           i.SectID,
		SectProgEntityID: i.SectProgEntityID,
		SectTitle:        i.SectTitle,
		SectDescription:  i.SectDescription,
		SectTotalSection: i.SectTotalSection,
		SectTotalLecture: i.SectTotalLecture,
		SectTotalMinute:  i.SectTotalMinute,
		SectModifiedDate: i.SectModifiedDate,
	}, nil
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

SELECT  sect_prog_entity_id, sect_title, sect_description,sect_total_minute FROM curriculum.sections
WHERE sect_prog_entity_id = $1
`

// curriculum.sections
func (q *Queries) Getsections(ctx context.Context, sectID int32) ([]mod.CurriculumSectionGet, error) {
	rows, err := q.db.QueryContext(ctx, getsections, sectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []mod.CurriculumSectionGet
	for rows.Next() {
		var i mod.CurriculumSectionGet
		err := rows.Scan(
			&i.SectProgEntityID,
			&i.SectTitle,
			&i.SectDescription,
			&i.SectTotalMinute,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

const listsections = `-- name: Listsections :many
SELECT sect_id, 
sect_prog_entity_id, 
sect_title, 
sect_description, 
sect_total_section, 
sect_total_lecture, 
sect_total_minute, 
sect_modified_date 
FROM curriculum.sections
ORDER BY sect_id
`

func (q *Queries) Listsections(ctx context.Context) ([]mod.CurriculumSection, error) {
	rows, err := q.db.QueryContext(ctx, listsections)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []mod.CurriculumSection
	for rows.Next() {
		var i mod.CurriculumSection
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
