package dbcontext

import (
	"context"

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
