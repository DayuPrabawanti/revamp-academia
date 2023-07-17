package dbcontext

import (
	"context"

	"codeid.revampacademy/models"
)

type Createprogram_entityParams struct {
	ProgTitle        string `db:"prog_title" json:"progTitle"`
	ProgHeadline     string `db:"prog_headline" json:"progHeadline"`
	ProgLearningType string `db:"prog_learning_type" json:"progLearningType"`
	ProgImage        string `db:"prog_image" json:"progImage"`
	ProgPrice        int32  `db:"prog_price" json:"progPrice"`
	ProgDuration     int32  `db:"prog_duration" json:"progDuration"`
}

const getprogram_entity = `-- name: Getprogram_entity :one
SELECT prog_entity_id, prog_title, prog_headline, prog_type, prog_learning_type, prog_rating, prog_total_trainee, prog_modified_date, prog_image, prog_best_seller, prog_price, prog_language, prog_duration, prog_duration_type, prog_tag_skill, prog_city_id, prog_cate_id, prog_created_by, prog_status FROM curriculum.program_entity
WHERE prog_entity_id = $1
`

func (q *Queries) Getprogram_entity(ctx context.Context, progTitle string) (models.CurriculumProgramEntity, error) {
	row := q.db.QueryRowContext(ctx, getprogram_entity, progTitle)
	var i models.CurriculumProgramEntity
	err := row.Scan(
		&i.ProgTitle,
		&i.ProgHeadline,
		&i.ProgLearningType,
		&i.ProgImage,
		&i.ProgPrice,
		&i.ProgDuration,
	)
	return i, err
}

const listprogram_entity = `-- name: Listprogram_entity :many
SELECT prog_entity_id, prog_title, prog_headline, prog_type, prog_learning_type, prog_rating, prog_total_trainee, prog_modified_date, prog_image, prog_best_seller, prog_price, prog_language, prog_duration, prog_duration_type, prog_tag_skill, prog_city_id, prog_cate_id, prog_created_by, prog_status FROM curriculum.program_entity
ORDER BY prog_title
`

func (q *Queries) Listprogram_entity(ctx context.Context) ([]*models.CurriculumProgramEntity, error) {
	rows, err := q.db.QueryContext(ctx, listprogram_entity)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*models.CurriculumProgramEntity
	for rows.Next() {
		var i *models.CurriculumProgramEntity
		if err := rows.Scan(
			&i.ProgTitle,
			&i.ProgHeadline,
			&i.ProgLearningType,
			&i.ProgImage,
			&i.ProgPrice,
			&i.ProgDuration,
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
