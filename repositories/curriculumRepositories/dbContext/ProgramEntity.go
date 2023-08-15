package dbContext

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	mod "codeid.revampacademy/models"
	"codeid.revampacademy/models/features"
)

const createprogram_entity = `-- name: Createprogram_entity :one

INSERT INTO curriculum.program_entity (prog_entity_id, 
prog_title, 
prog_headline, 
prog_type, 
prog_learning_type, 
prog_rating, 
prog_total_traniee, 
prog_modified_date, 
prog_image, 
prog_best_seller, 
prog_price, 
prog_language, 
prog_duration, 
prog_duration_type, 
prog_tag_skill, 
prog_city_id, 
prog_cate_id, 
prog_created_by, 
prog_status) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19)
RETURNING *
`

type Createprogram_entityParams struct {
	ProgEntityID     int32     `db:"prog_entity_id" json:"progEntityId"`
	ProgTitle        string    `db:"prog_title" json:"progTitle"`
	ProgHeadline     string    `db:"prog_headline" json:"progHeadline"`
	ProgType         string    `db:"prog_type" json:"progType"`
	ProgLearningType string    `db:"prog_learning_type" json:"progLearningType"`
	ProgRating       string    `db:"prog_rating" json:"progRating"`
	ProgTotalTraniee int32     `db:"prog_total_trainee" json:"progTotalTrainee"`
	ProgModifiedDate time.Time `db:"prog_modified_date" json:"progModifiedDate"`
	ProgImage        string    `db:"prog_image" json:"progImage"`
	ProgBestSeller   string    `db:"prog_best_seller" json:"progBestSeller"`
	ProgPrice        int32     `db:"prog_price" json:"progPrice"`
	ProgLanguage     string    `db:"prog_language" json:"progLanguage"`
	ProgDuration     int32     `db:"prog_duration" json:"progDuration"`
	ProgDurationType string    `db:"prog_duration_type" json:"progDurationType"`
	ProgTagSkill     string    `db:"prog_tag_skill" json:"progTagSkill"`
	ProgCityID       int32     `db:"prog_city_id" json:"progCityId"`
	ProgCateID       int32     `db:"prog_cate_id" json:"progCateId"`
	ProgCreatedBy    int32     `db:"prog_created_by" json:"progCreatedBy"`
	ProgStatus       string    `db:"prog_status" json:"progStatus"`
}

func (q *Queries) Createprogram_entity(ctx context.Context, arg Createprogram_entityParams) (*mod.CurriculumProgramEntity, *mod.ResponseError) {
	row := q.db.QueryRowContext(ctx, createprogram_entity,
		arg.ProgEntityID,
		arg.ProgTitle,
		arg.ProgHeadline,
		arg.ProgType,
		arg.ProgLearningType,
		arg.ProgRating,
		arg.ProgTotalTraniee,
		arg.ProgModifiedDate,
		arg.ProgImage,
		arg.ProgBestSeller,
		arg.ProgPrice,
		arg.ProgLanguage,
		arg.ProgDuration,
		arg.ProgDurationType,
		arg.ProgTagSkill,
		arg.ProgCityID,
		arg.ProgCateID,
		arg.ProgCreatedBy,
		arg.ProgStatus,
	)
	i := mod.CurriculumProgramEntity{}
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

	if err != nil {
		return nil, &mod.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &mod.CurriculumProgramEntity{
		ProgEntityID:     i.ProgEntityID,
		ProgTitle:        i.ProgTitle,
		ProgHeadline:     i.ProgHeadline,
		ProgType:         i.ProgType,
		ProgLearningType: i.ProgLearningType,
		ProgRating:       i.ProgRating,
		ProgTotalTraniee: i.ProgTotalTraniee,
		ProgModifiedDate: i.ProgModifiedDate,
		ProgImage:        i.ProgImage,
		ProgBestSeller:   i.ProgBestSeller,
		ProgPrice:        i.ProgPrice,
		ProgLanguage:     i.ProgLanguage,
		ProgDuration:     i.ProgDuration,
		ProgDurationType: i.ProgDurationType,
		ProgTagSkill:     i.ProgTagSkill,
		ProgCityID:       i.ProgCityID,
		ProgCateID:       i.ProgCateID,
		ProgCreatedBy:    i.ProgCreatedBy,
		ProgStatus:       i.ProgStatus,
	}, nil
}

const createCategory = `-- name: CreateCategory :one

INSERT INTO master.category (cate_id, 
cate_name, 
cate_cate_id, 
cate_modified_date, 
prog_status) VALUES ($1,$2,$3,$4)
RETURNING *
`

type CreateCategoryParams struct {
	CateID           int32         `db:"cate_id" json:"cateId"`
	CateName         string        `db:"cate_name" json:"cateName"`
	CateCateID       sql.NullInt32 `db:"cate_cate_id" json:"cateCateId"`
	CateModifiedDate sql.NullTime  `db:"cate_modified_date" json:"cateModifiedDate"`
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) (*mod.MasterCategory, *mod.ResponseError) {
	row := q.db.QueryRowContext(ctx, createCategory,
		arg.CateID,
		arg.CateName,
		arg.CateCateID,
		arg.CateModifiedDate,
	)
	i := mod.MasterCategory{}
	err := row.Scan(
		&i.CateID,
		&i.CateName,
		&i.CateCateID,
		&i.CateModifiedDate,
	)

	if err != nil {
		return nil, &mod.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &mod.MasterCategory{
		CateID:           i.CateID,
		CateName:         i.CateName,
		CateCateID:       i.CateCateID,
		CateModifiedDate: i.CateModifiedDate,
	}, nil
}

const deleteprogram_entity = `-- name: Deleteprogram_entity :exec
DELETE FROM curriculum.program_entity
WHERE prog_entity_id = $1
`

func (q *Queries) Deleteprogram_entity(ctx context.Context, progEntityID int32) error {
	_, err := q.db.ExecContext(ctx, deleteprogram_entity, progEntityID)
	return err
}

const getprogram_entity = `-- name: Getprogram_entity :one
SELECT prog_entity_id, prog_title, prog_headline, prog_type, prog_learning_type, prog_rating, prog_total_traniee, prog_modified_date, prog_image, prog_best_seller, prog_price, prog_language, prog_duration, prog_duration_type, prog_tag_skill, prog_city_id, prog_cate_id, prog_created_by, prog_status FROM curriculum.program_entity
WHERE prog_entity_id = $1
`

func (q *Queries) Getprogram_entity(ctx context.Context, progEntityID int32) (mod.CurriculumProgramEntity, error) {
	row := q.db.QueryRowContext(ctx, getprogram_entity, progEntityID)
	var i mod.CurriculumProgramEntity
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
	yearMonth := time.Now().Format("200601")
	seqNo := i.ProgEntityID // Atur nomor sekuen sesuai dengan kondisi Anda

	// Format nomor registrasi CURR-TAHUN-BULAN-SEQNO
	i.RegistrasiNumber = fmt.Sprintf("CURR%s#%03d", yearMonth, seqNo)

	return i, err
}

const listprogram_entity = `-- name: Listprogram_entity :many
SELECT prog_entity_id, 
prog_title, 
prog_headline, 
prog_type, 
prog_learning_type, 
prog_rating, 
prog_total_traniee, 
prog_modified_date, 
prog_image, 
prog_best_seller, 
prog_price, 
prog_language, 
prog_duration, 
prog_duration_type, 
prog_tag_skill, 
prog_city_id, 
prog_cate_id, 
prog_created_by, 
prog_status FROM curriculum.program_entity
ORDER BY prog_entity_id
limit $1 offset $2
`

func (q *Queries) Listprogram_entity(ctx context.Context, metadata *features.Metadata) ([]mod.CurriculumProgramEntity, error) {
	rows, err := q.db.QueryContext(ctx, listprogram_entity, metadata.PageSize, metadata.PageNo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []mod.CurriculumProgramEntity
	for rows.Next() {
		var i mod.CurriculumProgramEntity
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

const listMasterCategories = `-- name: ListCategories :many
SELECT cate_id, cate_name,cate_cate_id, cate_modified_date FROM master.category
ORDER BY cate_name
`

func (q *Queries) ListMasterCategories(ctx context.Context) ([]mod.MasterCategory, error) {
	rows, err := q.db.QueryContext(ctx, listMasterCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []mod.MasterCategory
	for rows.Next() {
		var i mod.MasterCategory
		if err := rows.Scan(
			&i.CateID,
			&i.CateName,
			&i.CateCateID,
			&i.CateModifiedDate); err != nil {
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

const getCategories = `-- name: GetMasterCategories :one

SELECT cate_id, cate_name,cate_cate_id, cate_modified_date FROM master.category
WHERE cate_id = $1
`

func (q *Queries) GetCategories(ctx context.Context, cateId int32) ([]mod.MasterCategory, error) {
	rows, err := q.db.QueryContext(ctx, getCategories, cateId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []mod.MasterCategory
	for rows.Next() {
		var i mod.MasterCategory
		err := rows.Scan(
			&i.CateID,
			&i.CateName,
			&i.CateCateID,
			&i.CateModifiedDate,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

const updateprogram_entity = `-- name: Updateprogram_entity :exec
UPDATE curriculum.program_entity
  set prog_title =$2, 
  prog_headline=$3, 
  prog_type=$4, 
  prog_learning_type=$5, 
  prog_rating=$6, 
  prog_total_traniee=$7, 
  prog_modified_date=$8, 
  prog_image=$9, 
  prog_best_seller=$10, 
  prog_price=$11, 
  prog_language=$12, 
  prog_duration=$13, 
  prog_duration_type=$14, 
  prog_tag_skill=$15, 
  prog_city_id=$16, 
  prog_cate_id=$17, 
  prog_created_by=$18, 
  prog_status=$19
WHERE prog_entity_id= $1 
`

type Updateprogram_entityParams struct {
	ProgEntityID     int32     `db:"prog_entity_id" json:"progEntityId"`
	ProgTitle        string    `db:"prog_title" json:"progTitle"`
	ProgHeadline     string    `db:"prog_headline" json:"progHeadline"`
	ProgType         string    `db:"prog_type" json:"progType"`
	ProgLearningType string    `db:"prog_learning_type" json:"progLearningType"`
	ProgRating       int32     `db:"prog_rating" json:"progRating"`
	ProgTotalTraniee int32     `db:"prog_total_traniee" json:"progTotalTraniee"`
	ProgModifiedDate time.Time `db:"prog_modified_date" json:"progModifiedDate"`
	ProgImage        string    `db:"prog_image" json:"progImage"`
	ProgBestSeller   string    `db:"prog_best_seller" json:"progBestSeller"`
	ProgPrice        int32     `db:"prog_price" json:"progPrice"`
	ProgLanguage     string    `db:"prog_language" json:"progLanguage"`
	ProgDuration     int32     `db:"prog_duration" json:"progDuration"`
	ProgDurationType string    `db:"prog_duration_type" json:"progDurationType"`
	ProgTagSkill     string    `db:"prog_tag_skill" json:"progTagSkill"`
	ProgCityID       int32     `db:"prog_city_id" json:"progCityId"`
	ProgCateID       int32     `db:"prog_cate_id" json:"progCateId"`
	ProgCreatedBy    int32     `db:"prog_created_by" json:"progCreatedBy"`
	ProgStatus       string    `db:"prog_status" json:"progStatus"`
}

func (q *Queries) Updateprogram_entity(ctx context.Context, arg Createprogram_entityParams) error {
	_, err := q.db.ExecContext(ctx, updateprogram_entity,
		arg.ProgEntityID,
		arg.ProgTitle,
		arg.ProgHeadline,
		arg.ProgType,
		arg.ProgLearningType,
		arg.ProgRating,
		arg.ProgTotalTraniee,
		arg.ProgModifiedDate,
		arg.ProgImage,
		arg.ProgBestSeller,
		arg.ProgPrice,
		arg.ProgLanguage,
		arg.ProgDuration,
		arg.ProgDurationType,
		arg.ProgTagSkill,
		arg.ProgCityID,
		arg.ProgCateID,
		arg.ProgCreatedBy,
		arg.ProgStatus,
	)
	return err
}
