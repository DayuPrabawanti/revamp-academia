package dbcontext

import (
	"context"
	"database/sql"
	"net/http"

	models "codeid.revampacademy/models"
)

const listSectionDetail = `-- name: ListSectionDetail :many
SELECT secd_id, secd_title, secd_preview, secd_score, secd_note, secd_minute, secd_modified_date, secd_sect_id 
FROM curriculum.section_detail
ORDER BY secd_id
`

func (q *Queries) ListSectionDetail(ctx context.Context) ([]models.CurriculumSectionDetail, error) {
	rows, err := q.db.QueryContext(ctx, listSectionDetail)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.CurriculumSectionDetail
	for rows.Next() {
		var i models.CurriculumSectionDetail
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

const getSectionDetail = `-- name: GetSectionDetail :one
SELECT secd_id, secd_title, secd_preview, secd_score, secd_note, secd_minute, secd_modified_date, secd_sect_id 
FROM curriculum.section_detail
	WHERE secd_id = $1
`

func (q *Queries) GetSectionDetail(ctx context.Context, secdId int16) (models.CurriculumSectionDetail, error) {
	row := q.db.QueryRowContext(ctx, getSectionDetail, secdId)
	var i models.CurriculumSectionDetail
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

type CreatesectionDetailParams struct {
	SecdID           int32          `db:"secd_id" json:"secdId"`
	SecdTitle        sql.NullString `db:"secd_title" json:"secdTitle"`
	SecdPreview      sql.NullString `db:"secd_preview" json:"secdPreview"`
	SecdScore        sql.NullInt32  `db:"secd_score" json:"secdScore"`
	SecdNote         sql.NullString `db:"secd_note" json:"secdNote"`
	SecdMinute       sql.NullInt32  `db:"secd_minute" json:"secdMinute"`
	SecdModifiedDate sql.NullTime   `db:"secd_modified_date" json:"secdModifiedDate"`
	SecdSectID       sql.NullInt32  `db:"secd_sect_id" json:"secdSectId"`
}

const createSectionDetail = `-- name: CreateSectionDetail :many

INSERT INTO curriculum.section_detail 
	(secd_id, 
	secd_title, 
	secd_preview, 
	secd_score, 
	secd_note, 
	secd_minute, 
	secd_modified_date, 
	secd_sect_id)
	VALUES($1,$2,$3,$4,$5,$6,$7,$8)
	RETURNING *
	`

func (q *Queries) CreateSectionDetail(ctx context.Context, arg CreatesectionDetailParams) (*models.CurriculumSectionDetail, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createSectionDetail,
		arg.SecdID,
		arg.SecdTitle,
		arg.SecdPreview,
		arg.SecdScore,
		arg.SecdNote,
		arg.SecdMinute,
		arg.SecdModifiedDate,
		arg.SecdSectID,
	)
	i := models.CurriculumSectionDetail{}
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
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.CurriculumSectionDetail{
		SecdID:           i.SecdID,
		SecdTitle:        i.SecdTitle,
		SecdPreview:      i.SecdPreview,
		SecdScore:        i.SecdScore,
		SecdNote:         i.SecdNote,
		SecdMinute:       i.SecdMinute,
		SecdModifiedDate: i.SecdModifiedDate,
		SecdSectID:       i.SecdSectID,
	}, nil
}
