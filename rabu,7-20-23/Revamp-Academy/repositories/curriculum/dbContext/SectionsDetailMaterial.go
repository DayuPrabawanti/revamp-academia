package dbContext

import (
	"context"
	"time"

	curi "codeid.revampacademy/models/curriculum"
)

const createsection_detail_material = `-- name: Createsection_detail_material :one

INSERT INTO curriculum.section_detail_material (sedm_id, 
sedm_filename, 
sedm_filesize, 
sedm_filetype, 
sedm_filelink, 
sedm_modified_date, 
sedm_secd_id)

VALUES($1,$2,$3,$4,$5,$6,$7)
RETURNING sedm_id
`

type Createsection_detail_materialParams struct {
	SedmID           int32     `db:"sedm_id" json:"sedmId"`
	SedmFilename     string    `db:"sedm_filename" json:"sedmFilename"`
	SedmFilesize     int32     `db:"sedm_filesize" json:"sedmFilesize"`
	SedmFiletype     string    `db:"sedm_filetype" json:"sedmFiletype"`
	SedmFilelink     string    `db:"sedm_filelink" json:"sedmFilelink"`
	SedmModifiedDate time.Time `db:"sedm_modified_date" json:"sedmModifiedDate"`
	SedmSecdID       int32     `db:"sedm_secd_id" json:"sedmSecdId"`
}

func (q *Queries) Createsection_detail_material(ctx context.Context, arg Createsection_detail_materialParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createsection_detail_material,
		arg.SedmID,
		arg.SedmFilename,
		arg.SedmFilesize,
		arg.SedmFiletype,
		arg.SedmFilelink,
		arg.SedmModifiedDate,
		arg.SedmSecdID,
	)
	var sedm_id int32
	err := row.Scan(&sedm_id)
	return sedm_id, err
}

const deletesection_detail_material = `-- name: Deletesection_detail_material :exec
DELETE FROM curriculum.section_detail_material
WHERE sedm_id = $1
`

func (q *Queries) Deletesection_detail_material(ctx context.Context, sedmID int32) error {
	_, err := q.db.ExecContext(ctx, deletesection_detail_material, sedmID)
	return err
}

const getsection_detail_material = `-- name: Getsection_detail_material :one

SELECT sedm_id, sedm_filename, sedm_filesize, sedm_filetype, sedm_filelink, sedm_modified_date, sedm_secd_id FROM curriculum.section_detail_material
WHERE sedm_id = $1
`

// curriculum.section_detail_material
func (q *Queries) Getsection_detail_material(ctx context.Context, sedmID int32) (curi.CurriculumSectionDetailMaterial, error) {
	row := q.db.QueryRowContext(ctx, getsection_detail_material, sedmID)
	var i curi.CurriculumSectionDetailMaterial
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

const listsection_detail_material = `-- name: Listsection_detail_material :many
SELECT sedm_id, sedm_filename, sedm_filesize, sedm_filetype, sedm_filelink, sedm_modified_date, sedm_secd_id FROM curriculum.section_detail_material
ORDER BY sedm_filename
`

func (q *Queries) Listsection_detail_material(ctx context.Context) ([]curi.CurriculumSectionDetailMaterial, error) {
	rows, err := q.db.QueryContext(ctx, listsection_detail_material)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []curi.CurriculumSectionDetailMaterial
	for rows.Next() {
		var i curi.CurriculumSectionDetailMaterial
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

const updatesection_detail_material = `-- name: Updatesection_detail_material :exec
UPDATE curriculum.section_detail_material
  set sedm_filename = $2,
  sedm_filesize = $3
WHERE sedm_id = $1
`

type Updatesection_detail_materialParams struct {
	SedmID       int32  `db:"sedm_id" json:"sedmId"`
	SedmFilename string `db:"sedm_filename" json:"sedmFilename"`
	SedmFilesize int32  `db:"sedm_filesize" json:"sedmFilesize"`
}

func (q *Queries) Updatesection_detail_material(ctx context.Context, arg Updatesection_detail_materialParams) error {
	_, err := q.db.ExecContext(ctx, updatesection_detail_material, arg.SedmID, arg.SedmFilename, arg.SedmFilesize)
	return err
}
