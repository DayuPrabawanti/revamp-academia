package dbcontext

import (
	"context"

	models "codeid.revampacademy/db-generator/gen"
)

const Group = `-- name: Group :many
SELECT 
curriculum.program_entity.prog_entity_id,
curriculum.sections.sect_title, 
curriculum.sections.sect_description, 
curriculum.sections.sect_total_section, 
curriculum.sections.sect_total_lecture, 
curriculum.sections.sect_total_minute, 
curriculum.sections.sect_modified_date,
curriculum.program_entity.prog_title, 
curriculum.program_entity.prog_headline, 
curriculum.program_entity.prog_type, 
curriculum.program_entity.prog_learning_type, 
curriculum.program_entity.prog_rating, 
curriculum.program_entity.prog_total_trainee, 
curriculum.program_entity.prog_modified_date, 
curriculum.program_entity.prog_image, 
curriculum.program_entity.prog_best_seller, 
curriculum.program_entity.prog_price, 
curriculum.program_entity.prog_language, 
curriculum.program_entity.prog_duration, 
curriculum.program_entity.prog_duration_type, 
curriculum.program_entity.prog_tag_skill, 
curriculum.program_entity.prog_city_id, 
curriculum.program_entity.prog_cate_id, 
curriculum.program_entity.prog_created_by
FROM curriculum.sections
JOIN curriculum.program_entity ON curriculum.program_entity.prog_entity_id = curriculum.sections.sect_prog_entity_id
ORDER BY curriculum.program_entity;
`

func (q *Queries) GroupImpl(ctx context.Context) ([]models.Group, error) {
	rows, err := q.db.QueryContext(ctx, Group)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.Group
	for rows.Next() {
		var i models.Group
		if err := rows.Scan(
			&i.CurriculumProgramEntity.ProgEntityID,
			&i.CurriculumSection.SectTitle,
			&i.CurriculumSection.SectDescription,
			&i.CurriculumSection.SectTotalSection,
			&i.CurriculumSection.SectTotalLecture,
			&i.CurriculumSection.SectTotalMinute,
			&i.CurriculumSection.SectModifiedDate,
			&i.CurriculumProgramEntity.ProgTitle,
			&i.CurriculumProgramEntity.ProgHeadline,
			&i.CurriculumProgramEntity.ProgType,
			&i.CurriculumProgramEntity.ProgLearningType,
			&i.CurriculumProgramEntity.ProgRating,
			&i.CurriculumProgramEntity.ProgTotalTrainee,
			&i.CurriculumProgramEntity.ProgModifiedDate,
			&i.CurriculumProgramEntity.ProgImage,
			&i.CurriculumProgramEntity.ProgBestSeller,
			&i.CurriculumProgramEntity.ProgPrice,
			&i.CurriculumProgramEntity.ProgLanguage,
			&i.CurriculumProgramEntity.ProgDuration,
			&i.CurriculumProgramEntity.ProgDurationType,
			&i.CurriculumProgramEntity.ProgTagSkill,
			&i.CurriculumProgramEntity.ProgCityID,
			&i.CurriculumProgramEntity.ProgCateID,
			&i.CurriculumProgramEntity.ProgCreatedBy,
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
