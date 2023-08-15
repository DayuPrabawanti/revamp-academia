package dbContext

import (
	"context"
	"database/sql"

	"codeid.revampacademy/models"
)

type CurriculumProgramEntityParams struct {
	ProgEntityID     int32          `db:"prog_entity_id" json:"progEntityId"`
	ProgTitle        string         `db:"prog_title" json:"progTitle"`
	ProgHeadline     sql.NullString `db:"prog_headline" json:"progHeadline"`
	ProgType         sql.NullString `db:"prog_type" json:"progType"`
	ProgLearningType sql.NullString `db:"prog_learning_type" json:"progLearningType"`
	ProgRating       sql.NullInt32  `db:"prog_rating" json:"progRating"`
	ProgTotalTrainee sql.NullInt32  `db:"prog_total_trainee" json:"progTotalTrainee"`
	ProgModifiedDate sql.NullTime   `db:"prog_modified_date" json:"progModifiedDate"`
	ProgImage        sql.NullString `db:"prog_image" json:"progImage"`
	ProgBestSeller   sql.NullString `db:"prog_best_seller" json:"progBestSeller"`
	ProgPrice        sql.NullInt32  `db:"prog_price" json:"progPrice"`
	ProgLanguage     sql.NullString `db:"prog_language" json:"progLanguage"`
	ProgDuration     sql.NullInt32  `db:"prog_duration" json:"progDuration"`
	ProgDurationType sql.NullString `db:"prog_duration_type" json:"progDurationType"`
	ProgTagSkill     sql.NullString `db:"prog_tag_skill" json:"progTagSkill"`
	ProgCityID       sql.NullInt32  `db:"prog_city_id" json:"progCityId"`
	ProgCateID       sql.NullInt32  `db:"prog_cate_id" json:"progCateId"`
	ProgCreatedBy    sql.NullInt32  `db:"prog_created_by" json:"progCreatedBy"`
	ProgStatus       sql.NullString `db:"prog_status" json:"progStatus"`
}

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

const updateProgramEntity = `-- name: UpdateProgramEntity :exec
UPDATE curriculum.program_entity
set prog_title = $2,
prog_headline = $3,
prog_type = $4,
prog_learning_type =$5,
prog_rating =$6,
prog_total_trainee = $7,
prog_image = $8,
prog_best_seller = $9,
prog_price = $10,
prog_language = $11,
prog_modified_date = $12,
prog_duration = $13,
prog_duration_type = $14,
prog_tag_skill = $15,
prog_city_id = $16,
prog_cate_id =$17,
prog_created_by = $18,
prog_status = $19
WHERE prog_entity_id = $1
`

func (q *Queries) UpdateProgramEntity(ctx context.Context, arg CurriculumProgramEntityParams) error {
	_, err := q.db.ExecContext(ctx, updateProgramEntity, arg.ProgEntityID, arg.ProgTitle, arg.ProgHeadline, arg.ProgType, arg.ProgLearningType, arg.ProgRating, arg.ProgTotalTrainee, arg.ProgImage, arg.ProgBestSeller, arg.ProgPrice, arg.ProgLanguage, arg.ProgModifiedDate, arg.ProgDuration, arg.ProgDurationType, arg.ProgTagSkill, arg.ProgCityID, arg.ProgCateID, arg.ProgCreatedBy, arg.ProgStatus)
	return err
}

const GetProgEntity = `-- name:  GetProgEntity :one
SELECT prog_entity_id, prog_title, prog_headline, prog_type, prog_learning_type, prog_rating, prog_total_trainee, prog_image, prog_best_seller, prog_price, prog_language, prog_modified_date, prog_duration, prog_duration_type, prog_tag_skill, prog_city_id, prog_cate_id, prog_created_by, prog_status 
, picture FROM progentityid
WHERE prog_entity_id = $1
`

func (q *Queries) GetProgEntity(ctx context.Context, progEntityId int32) (models.CurriculumProgramEntity, error) {
	row := q.db.QueryRowContext(ctx, GetProgEntity, progEntityId)
	var i models.CurriculumProgramEntity
	err := row.Scan(
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
	)
	return i, err
}
