package dbcontext

import (
	"context"

	models "codeid.revampacademy/models"
)

const group = `-- name: Group :many
select pe.prog_title, 
		ct.cate_name, 
		se.sect_id,
		sd.secd_id
from curriculum.program_entity pe
join master.category ct
on pe.prog_cate_id = ct.cate_id
join curriculum.sections se
on se.sect_prog_entity_id = pe.prog_entity_id
join curriculum.section_detail sd
on pe.prog_entity_id = se.sect_id
`

func (q *Queries) Group(ctx context.Context) ([]models.Group, error) {
	rows, err := q.db.QueryContext(ctx, group)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.Group
	for rows.Next() {
		var i models.Group
		if err := rows.Scan(
			&i.CurriculumProgramEntity.ProgTitle,
			&i.MasterCategory.CateName,
			&i.CurriculumSection.SectID,
			&i.CurriculumSectionDetail.SecdID,
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
