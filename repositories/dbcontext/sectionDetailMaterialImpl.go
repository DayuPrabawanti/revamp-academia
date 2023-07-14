package dbcontext

import (
	"context"

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
