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

// const getGroup = `-- name: GetGroup :one
// SELECT secd_id, secd_title, secd_preview, secd_score, secd_note, secd_minute, secd_modified_date, secd_sect_id
// FROM curriculum.section_detail
// 	WHERE secd_id = $1
// `

// func (q *Queries) GetSectionDetail(ctx context.Context, secdId int16) (models.CurriculumSectionDetail, error) {
// 	row := q.db.QueryRowContext(ctx, getSectionDetail, secdId)
// 	var i models.CurriculumSectionDetail
// 	err := row.Scan(
// 		&i.SecdID,
// 		&i.SecdTitle,
// 		&i.SecdPreview,
// 		&i.SecdScore,
// 		&i.SecdNote,
// 		&i.SecdMinute,
// 		&i.SecdModifiedDate,
// 		&i.SecdSectID,
// 	)
// 	return i, err
// }

type CreateGroupParams struct {
	CreateProgramEntityParams CreateProgramEntityParams
	CreatesectionsParams      CreatesectionsParams
}

// const createGroup = `-- name: CreateGroup :many

// INSERT INTO curriculum.sections (sect_id,
// sect_prog_entity_id,
// sect_title,
// sect_description,
// sect_total_section,
// sect_total_lecture,
// sect_total_minute,
// sect_modified_date)
// VALUES($1,$2,$3,$4,$5,$6,$7,$8)
// RETURNING *
// `

// func (q *Queries) Createsections(ctx context.Context, arg CreatesectionsParams) (*models.CurriculumSection, *models.ResponseError) {
// 	row := q.db.QueryRowContext(ctx, createsections,
// 		arg.SectID,
// 		arg.SectProgEntityID,
// 		arg.SectTitle,
// 		arg.SectDescription,
// 		arg.SectTotalSection,
// 		arg.SectTotalLecture,
// 		arg.SectTotalMinute,
// 		arg.SectModifiedDate,
// 	)
// 	i := models.CurriculumSection{}
// 	err := row.Scan(
// 		&i.SectID,
// 		&i.SectProgEntityID,
// 		&i.SectTitle,
// 		&i.SectDescription,
// 		&i.SectTotalSection,
// 		&i.SectTotalLecture,
// 		&i.SectTotalMinute,
// 		&i.SectModifiedDate,
// 	)
// 	if err != nil {
// 		return nil, &models.ResponseError{
// 			Message: err.Error(),
// 			Status:  http.StatusInternalServerError,
// 		}
// 	}
// 	return &models.CurriculumSection{
// 		SectID:           i.SectID,
// 		SectProgEntityID: i.SectProgEntityID,
// 		SectTitle:        i.SectTitle,
// 		SectDescription:  i.SectDescription,
// 		SectTotalSection: i.SectTotalSection,
// 		SectTotalLecture: i.SectTotalLecture,
// 		SectTotalMinute:  i.SectTotalMinute,
// 		SectModifiedDate: i.SectModifiedDate,
// 	}, nil
// }
