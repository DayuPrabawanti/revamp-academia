package dbContext

import (
	"context"
	"net/http"
	"time"

	models "codeid.revampacademy/models"
)

const listSectionDetailMaterial = `-- name: ListSectionDetailMaterial :many
SELECT sedm_id, sedm_filename, sedm_filesize, sedm_filetype, sedm_filelink, sedm_modified_date, sedm_secd_id 
FROM curriculum.section_detail_material
ORDER BY sedm_id
`

func (q *Queries) ListSectionDetailMaterial(ctx context.Context) ([]models.CurriculumSectionDetailMaterial, error) {
	rows, err := q.db.QueryContext(ctx, listSectionDetailMaterial)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.CurriculumSectionDetailMaterial
	for rows.Next() {
		var i models.CurriculumSectionDetailMaterial
		if err := rows.Scan(
			&i.SedmID,
			&i.SedmFilename,
			&i.SedmFilesize,
			&i.SedmFiletype,
			&i.SedmFilelink,
			&i.SedmModifiedDate,
			&i.SedmSecdID,
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

const getSectionDetailMaterial = `-- name: GetSectionDetailMaterial :one
SELECT sedm_id, sedm_filename, sedm_filesize, sedm_filetype, sedm_filelink, sedm_modified_date, sedm_secd_id 
FROM curriculum.section_detail_material
	WHERE sedm_id = $1
`

func (q *Queries) GetSectionDetailMaterial(ctx context.Context, sedmID int16) (models.CurriculumSectionDetailMaterial, error) {
	row := q.db.QueryRowContext(ctx, getSectionDetailMaterial, sedmID)
	var i models.CurriculumSectionDetailMaterial
	err := row.Scan(
		&i.SedmID,
		&i.SedmFilename,
		&i.SedmFilesize,
		&i.SedmFiletype,
		&i.SedmFilelink,
		&i.SedmModifiedDate,
		&i.SedmSecdID,
	)
	return i, err
}

type CreatesectionDetailMaterialParams struct {
	SedmID           int32     `db:"sedm_id" json:"sedmId"`
	SedmFilename     string    `db:"sedm_filename" json:"sedmFilename"`
	SedmFilesize     int32     `db:"sedm_filesize" json:"sedmFilesize"`
	SedmFiletype     string    `db:"sedm_filetype" json:"sedmFiletype"`
	SedmFilelink     string    `db:"sedm_filelink" json:"sedmFilelink"`
	SedmModifiedDate time.Time `db:"sedm_modified_date" json:"sedmModifiedDate"`
	SedmSecdID       int32     `db:"sedm_secd_id" json:"sedmSecdId"`
}

const createsectiondetailMaterial = `-- name: CreatesectiondetailMaterial :many

INSERT INTO curriculum.section_detail_material (sedm_id, 
	sedm_filename, 
	sedm_filesize, 
	sedm_filetype, 
	sedm_filelink, 
	sedm_modified_date, 
	sedm_secd_id)
	
	VALUES($1,$2,$3,$4,$5,$6,$7)
	RETURNING *
	`

func (q *Queries) CreatesectiondetailMaterial(ctx context.Context, arg CreatesectionDetailMaterialParams) (*models.CurriculumSectionDetailMaterial, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createsectiondetailMaterial,
		arg.SedmID,
		arg.SedmFilename,
		arg.SedmFilesize,
		arg.SedmFiletype,
		arg.SedmFilelink,
		arg.SedmModifiedDate,
		arg.SedmSecdID,
	)
	i := models.CurriculumSectionDetailMaterial{}
	err := row.Scan(
		&i.SedmID,
		&i.SedmFilename,
		&i.SedmFilesize,
		&i.SedmFiletype,
		&i.SedmFilelink,
		&i.SedmModifiedDate,
		&i.SedmSecdID,
	)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.CurriculumSectionDetailMaterial{
		SedmID:           i.SedmID,
		SedmFilename:     i.SedmFilename,
		SedmFilesize:     i.SedmFilesize,
		SedmFiletype:     i.SedmFiletype,
		SedmFilelink:     i.SedmFilelink,
		SedmModifiedDate: i.SedmModifiedDate,
		SedmSecdID:       i.SedmSecdID,
	}, nil
}

const updateSectionDetailMaterial = `-- name: UpdateSectionDetailMaterial :exec
UPDATE curriculum.section_detail_material
SET
sedm_filename = $2,
sedm_filesize = $3,
sedm_filetype = $4,
sedm_filelink = $5,
sedm_modified_date = $6,
sedm_secd_id = $7
WHERE sedm_id = $1
`

func (q *Queries) UpdateSectionDetailMaterial(ctx context.Context, arg CreatesectionDetailMaterialParams) error {
	_, err := q.db.ExecContext(ctx, updateSectionDetailMaterial,
		arg.SedmID,
		arg.SedmFilename,
		arg.SedmFilesize,
		arg.SedmFiletype,
		arg.SedmFilelink,
		arg.SedmModifiedDate,
		arg.SedmSecdID)
	return err
}

const deleteSectionDetailMaterial = `-- name: DeleteSectionDetailMaterial :exec
DELETE FROM curriculum.section_detail_material
WHERE sedm_id = $1
`

func (q *Queries) DeleteSectionDetailMaterial(ctx context.Context, sedmID int16) error {
	_, err := q.db.ExecContext(ctx, deleteSectionDetailMaterial, sedmID)
	return err
}
