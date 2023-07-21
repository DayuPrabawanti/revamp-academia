package dbcontext

import (
	"context"
	"net/http"
	"time"

	models "codeid.revampacademy/models"
)

const listSections = `-- name: ListSections :many
SELECT sect_title, sect_description, sect_total_section, sect_total_lecture, sect_total_minute, sect_modified_date, sect_prog_entity_id, sect_id 
FROM curriculum.sections
ORDER BY sect_title
`

func (q *Queries) ListSections(ctx context.Context) ([]models.CurriculumSection, error) {
	rows, err := q.db.QueryContext(ctx, listSections)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.CurriculumSection
	for rows.Next() {
		var i models.CurriculumSection
		if err := rows.Scan(
			&i.SectTitle,
			&i.SectDescription,
			&i.SectTotalSection,
			&i.SectTotalLecture,
			&i.SectTotalMinute,
			&i.SectModifiedDate,
			&i.SectProgEntityID,
			&i.SectID,
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

const getSections = `-- name: GetSections :one
SELECT sect_title, sect_description, sect_total_section, sect_total_lecture, sect_total_minute, sect_modified_date, sect_prog_entity_id, sect_id 
FROM curriculum.sections
	WHERE sect_id = $1
`

func (q *Queries) GetSections(ctx context.Context, sectId int16) (models.CurriculumSection, error) {
	row := q.db.QueryRowContext(ctx, getSections, sectId)
	var i models.CurriculumSection
	err := row.Scan(
		&i.SectTitle,
		&i.SectDescription,
		&i.SectTotalSection,
		&i.SectTotalLecture,
		&i.SectTotalMinute,
		&i.SectModifiedDate,
		&i.SectProgEntityID,
		&i.SectID,
	)
	return i, err
}

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

func (q *Queries) Createsections(ctx context.Context, arg CreatesectionsParams) (*models.CurriculumSection, *models.ResponseError) {
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
	i := models.CurriculumSection{}
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
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.CurriculumSection{
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

const updateSections = `-- name: updateSections :exec
UPDATE curriculum.sections
SET
sect_prog_entity_id = $2, 
sect_title = $3, 
sect_description = $4, 
sect_total_section = $5, 
sect_total_lecture = $6, 
sect_total_minute = $7, 
sect_modified_date = $8
WHERE sect_id = $1
`

func (q *Queries) UpdateSections(ctx context.Context, arg CreatesectionsParams) error {
	_, err := q.db.ExecContext(ctx, updateSections,
		arg.SectID,
		arg.SectProgEntityID,
		arg.SectTitle,
		arg.SectDescription,
		arg.SectTotalSection,
		arg.SectTotalLecture,
		arg.SectTotalMinute,
		arg.SectModifiedDate)
	return err
}

const deleteSections = `-- name: DeleteSections :exec
DELETE FROM curriculum.sections
WHERE sect_id = $1
`

func (q *Queries) DeleteSections(ctx context.Context, sectId int16) error {
	_, err := q.db.ExecContext(ctx, deleteSections, sectId)
	return err
}
