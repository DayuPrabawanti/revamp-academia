package dbContext

import (
	"context"

	"codeid.revampacademy/models"
)

const listProgramEntity = `-- name: ListProgramEntity :many
SELECT prog_entity_id, prog_title, prog_headline, prog_type, prog_learning_type, prog_rating, prog_total_trainee, prog_image, prog_best_seller, prog_price, prog_language, prog_modified_date, prog_duration, prog_duration_type, prog_tag_skill, prog_city_id, prog_cate_id, prog_created_by, prog_status 
	FROM curriculum.program_entity
	ORDER BY prog_title
`

func (q *Queries) ListProgramEntitys(ctx context.Context) ([]models.CurriculumProgramEntity, error) {
	rows, err := q.db.QueryContext(ctx, listProgramEntity)
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
			&i.ProgTotalTrainee,
			&i.ProgImage,
			&i.ProgBestSeller,
			&i.ProgPrice,
			&i.ProgLanguage,
			&i.ProgModifiedDate,
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
