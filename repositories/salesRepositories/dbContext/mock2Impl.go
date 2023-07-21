package dbContext

import (
	"context"

	"codeid.revampacademy/models"
)

const getProgramEntityId = `-- name: getProgramEntityId :one
SELECT prog_entity_id, prog_title, prog_headline, prog_type, prog_learning_type, prog_rating, prog_total_traniee, prog_modified_date, prog_image, prog_best_seller, prog_price, prog_language, prog_duration, prog_duration_type, prog_tag_skill, prog_city_id, prog_cate_id, prog_created_by, prog_status FROM curriculum.program_entity
WHERE prog_entity_id = $1
`

func (q *Queries) GetProgramEntityId(ctx context.Context, progEntityID int32) (models.CurriculumProgramEntity, error) {
	row := q.db.QueryRowContext(ctx, getProgramEntityId, progEntityID)
	var i models.CurriculumProgramEntity
	err := row.Scan(
		&i.ProgEntityID,
		&i.ProgTitle,
		&i.ProgHeadline,
		&i.ProgType,
		&i.ProgLearningType,
		&i.ProgRating,
		&i.ProgTotalTraniee,
		&i.ProgModifiedDate,
		&i.ProgImage,
		&i.ProgBestSeller,
		&i.ProgPrice,
		&i.ProgLanguage,
		&i.ProgDuration,
		&i.ProgDurationType,
		&i.ProgTagSkill,
		&i.ProgCityID,
		&i.ProgCateID,
		&i.ProgCreatedBy,
		&i.ProgStatus,
	)
	return i, err
}

const listProgram = `-- name: Listprogram_entity :many
SELECT prog_entity_id, prog_title, prog_headline, prog_type, prog_learning_type, prog_rating, prog_total_traniee, prog_modified_date, prog_image, prog_best_seller, prog_price, prog_language, prog_duration, prog_duration_type, prog_tag_skill, prog_city_id, prog_cate_id, prog_created_by, prog_status FROM curriculum.program_entity
where prog_learning_type = $1
`

func (q *Queries) ListProgram(ctx context.Context, orderby string) ([]models.CurriculumProgramEntity, error) {
	rows, err := q.db.QueryContext(ctx, listProgram, orderby)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.CurriculumProgramEntity
	for rows.Next() {
		var i models.CurriculumProgramEntity
		if err := rows.Scan(
			&i.ProgEntityID,
			&i.ProgTitle,
			&i.ProgHeadline,
			&i.ProgType,
			&i.ProgLearningType,
			&i.ProgRating,
			&i.ProgTotalTraniee,
			&i.ProgModifiedDate,
			&i.ProgImage,
			&i.ProgBestSeller,
			&i.ProgPrice,
			&i.ProgLanguage,
			&i.ProgDuration,
			&i.ProgDurationType,
			&i.ProgTagSkill,
			&i.ProgCityID,
			&i.ProgCateID,
			&i.ProgCreatedBy,
			&i.ProgStatus,
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
