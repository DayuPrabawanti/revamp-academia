package dbcontext

import (
	"context"

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
