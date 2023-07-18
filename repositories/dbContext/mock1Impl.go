package dbcontext

import (
	"context"

	"codeid.revampacademy/models"
)

type CreateprogramEntityParams struct {
	ProgTitle        string `db:"prog_title" json:"progTitle"`
	ProgHeadline     string `db:"prog_headline" json:"progHeadline"`
	ProgLearningType string `db:"prog_learning_type" json:"progLearningType"`
	ProgImage        string `db:"prog_image" json:"progImage"`
	ProgPrice        int32  `db:"prog_price" json:"progPrice"`
	ProgDuration     int32  `db:"prog_duration" json:"progDuration"`
}

const getProgramEntity = `-- name: getProgramEntity :one
SELECT prog_title, prog_headline, prog_learning_type, prog_image, prog_price, prog_duration FROM curriculum.program_entity
WHERE prog_title = $1
`

func (q *Queries) GetProgramEntity(ctx context.Context, nama string) (CreateprogramEntityParams, error) {
	row := q.db.QueryRowContext(ctx, getProgramEntity, nama)
	var i CreateprogramEntityParams
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

const getProgramEntityId = `-- name: getProgramEntityId :one
SELECT prog_entity_id, prog_title, prog_headline, prog_type, prog_learning_type, prog_rating, prog_total_trainee, prog_modified_date, prog_image, prog_best_seller, prog_price, prog_language, prog_duration, prog_duration_type, prog_tag_skill, prog_city_id, prog_cate_id, prog_created_by, prog_status FROM curriculum.program_entity
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
		&i.ProgTotalTrainee,
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
