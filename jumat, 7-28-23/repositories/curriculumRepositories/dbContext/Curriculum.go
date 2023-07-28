package dbContext

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	mod "codeid.revampacademy/models"
)

const getprogramentity = `-- name: Getprogram_entity :one
SELECT prog_entity_id, 
prog_title, 
prog_headline, 
prog_type, 
prog_learning_type, 
prog_modified_date, 
prog_image, 
prog_price, 
prog_language, 
prog_duration, 
prog_duration_type, 
prog_tag_skill, 
prog_cate_id,
pred_prog_entity_id,
pred_item_learning,
pred_description
FROM curriculum.program_entity cpe
LEFT JOIN curriculum.program_entity_description cped
ON cpe.prog_entity_id = cped.pred_prog_entity_id
WHERE prog_entity_id = $1
`

func (q *Queries) Getprogramentity(ctx context.Context, progEntityID int32) (mod.CurriculumProgramEntityMockup, error) {
	row := q.db.QueryRowContext(ctx, getprogramentity, progEntityID)
	var i mod.CurriculumProgramEntityMockup
	err := row.Scan(
		&i.ProgEntityID,
		&i.ProgTitle,
		&i.ProgHeadline,
		&i.ProgType,
		&i.ProgLearningType,
		&i.ProgModifiedDate,
		&i.ProgImage,
		&i.ProgPrice,
		&i.ProgLanguage,
		&i.ProgDuration,
		&i.ProgDurationType,
		&i.ProgTagSkill,
		&i.ProgCateID,
		&i.PredProgEntityID,
		&i.PredItemLearning,
		&i.PredDescription,
	)
	yearMonth := time.Now().Format("200601")
	seqNo := i.ProgEntityID // Atur nomor sekuen sesuai dengan kondisi Anda

	// Format nomor registrasi CURR-TAHUN-BULAN-SEQNO
	i.RegistrasiNumber = fmt.Sprintf("CURR%s#%03d", yearMonth, seqNo)
	return i, err
}

// impl.go

const getsection = `-- name: Getsections :one
SELECT  sect_prog_entity_id, 
        sect_id,
        sect_title, 
        sect_total_minute 
FROM curriculum.sections
WHERE sect_prog_entity_id = $1
`

// curriculum.sections
func (q *Queries) Getsection(ctx context.Context, sectID int64) ([]mod.CurriculumSectionGet, error) {
	rows, err := q.db.QueryContext(ctx, getsection, sectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []mod.CurriculumSectionGet
	for rows.Next() {
		var i mod.CurriculumSectionGet
		err := rows.Scan(
			&i.SectProgEntityID,
			&i.SectID,
			&i.SectTitle,
			&i.SectTotalMinute,
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

// impl.go

const getsectiondetail = `-- name: Getsection_detail :one
SELECT 
    secd_sect_id,
    secd_title,
    secd_score,
    secd_minute,
	sedm_filename,
	sedm_filesize,
	sedm_filetype,
	sedm_filelink
FROM curriculum.section_detail csd
LEFT JOIN curriculum.section_detail_material csdm
ON csd.secd_id = csdm.sedm_secd_id
WHERE secd_sect_id = $1
`

// curriculum.section_detail
func (q *Queries) Getsectiondetail(ctx context.Context, secdID int64) ([]mod.CurriculumSectionDetailMockup, error) {
	rows, err := q.db.QueryContext(ctx, getsectiondetail, secdID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []mod.CurriculumSectionDetailMockup
	for rows.Next() {
		var i mod.CurriculumSectionDetailMockup
		err := rows.Scan(
			&i.SecdSectID,
			&i.SecdTitle,
			&i.SecdScore,
			&i.SecdMinute,
			&i.SedmFilename,
			&i.SedmFilesize,
			&i.SedmFiletype,
			&i.SedmFilelink,
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

const getmscategory = `-- name: ListCategories :many
SELECT 
prog_entity_id,
cate_id,
cate_name 
FROM curriculum.program_entity cpe 
JOIN master.category mc 
ON cpe.prog_cate_id = mc.cate_id
WHERE cpe.prog_entity_id = $1
`

func (q *Queries) GetMsCategory(ctx context.Context, masterID int32) ([]mod.MasterCategoryProgramEntity, error) {
	rows, err := q.db.QueryContext(ctx, getmscategory, masterID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []mod.MasterCategoryProgramEntity
	for rows.Next() {
		var i mod.MasterCategoryProgramEntity
		err := rows.Scan(
			&i.ProgEntityID,
			&i.CateID,
			&i.CateName,
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

const getprogramentitydescription = `-- name: Getprogram_entity_description :one

SELECT pred_prog_entity_id, 
pred_item_learning, 
pred_description
FROM curriculum.program_entity_description
WHERE pred_prog_entity_id = $1
`

// curriculum.programentitydescription
func (q *Queries) Getprogramentitydescription(ctx context.Context, predProgEntityID int32) (mod.CurriculumProgramEntityDescriptionMockup, error) {
	row := q.db.QueryRowContext(ctx, getprogramentitydescription, predProgEntityID)
	var i mod.CurriculumProgramEntityDescriptionMockup
	err := row.Scan(
		&i.PredProgEntityID,
		&i.PredItemLearning,
		&i.PredDescription,
	)
	return i, err
}

// update
type UpdateCurriculum struct {
	UpdateprogramentityParams            UpdateprogramentityParams
	UpdateprogramentitydescriptionParams UpdateprogramentitydescriptionParams
	UpdateScoreParams                    UpdateScoreParams
}

// const updatecurriculum = `UPDATE curriculum.program_entity AS t1
// SET set prog_title =$2,
// prog_headline=$3,
// prog_type=$4,
// prog_learning_type=$5,
// prog_modified_date=$6,
// prog_image=$7,
// prog_price=$8,
// prog_language=$9,
// prog_duration=$10,
// prog_duration_type=$11,
// prog_tag_skill=$12,
// prog_cate_id=$13
// WHERE prog_entity_id= $1;

// UPDATE curriculum.program_entity_description AS t2
// SET pred_item_learning= $15,
// pred_description = $16
// WHERE pred_prog_entity_id= $1;
// `

// type UpdateCurriculumParams struct {
// 	ProgEntityID     int32          `db:"prog_entity_id" json:"progEntityId"`
// 	ProgTitle        string         `db:"prog_title" json:"progTitle"`
// 	ProgHeadline     string         `db:"prog_headline" json:"progHeadline"`
// 	ProgType         string         `db:"prog_type" json:"progType"`
// 	ProgLearningType string         `db:"prog_learning_type" json:"progLearningType"`
// 	ProgModifiedDate time.Time      `db:"prog_modified_date" json:"progModifiedDate"`
// 	ProgImage        string         `db:"prog_image" json:"progImage"`
// 	ProgPrice        int32          `db:"prog_price" json:"progPrice"`
// 	ProgLanguage     string         `db:"prog_language" json:"progLanguage"`
// 	ProgDuration     int32          `db:"prog_duration" json:"progDuration"`
// 	ProgDurationType string         `db:"prog_duration_type" json:"progDurationType"`
// 	ProgTagSkill     string         `db:"prog_tag_skill" json:"progTagSkill"`
// 	ProgCateID       int32          `db:"prog_cate_id" json:"progCateId"`
// 	PredProgEntityID int32          `db:"pred_prog_entity_id" json:"predProgEntityId"`
// 	PredItemLearning sql.NullString `db:"pred_item_learning" json:"predItemLearning"`
// 	PredDescription  sql.NullString `db:"pred_description" json:"predDescription"`
// }

// func (q *Queries) UpdateCurriculum(ctx context.Context, arg UpdateCurriculumParams) error {
// 	_, err := q.db.ExecContext(ctx, updatecurriculum,
// 		arg.ProgEntityID,
// 		arg.ProgTitle,
// 		arg.ProgHeadline,
// 		arg.ProgType,
// 		arg.ProgLearningType,
// 		arg.ProgModifiedDate,
// 		arg.ProgImage,
// 		arg.ProgPrice,
// 		arg.ProgLanguage,
// 		arg.ProgDuration,
// 		arg.ProgDurationType,
// 		arg.ProgTagSkill,
// 		arg.ProgCateID,
// 		arg.PredProgEntityID,
// 		arg.PredItemLearning,
// 		arg.PredDescription,
// 	)
// 	return err
// }

const updateprogramentity = `-- name: Updateprogramentity :exec
UPDATE curriculum.program_entity
  set prog_title =$2, 
  prog_headline=$3, 
  prog_type=$4, 
  prog_learning_type=$5, 
  prog_modified_date=$6, 
  prog_image=$7, 
  prog_price=$8, 
  prog_language=$9, 
  prog_duration=$10, 
  prog_duration_type=$11, 
  prog_tag_skill=$12, 
  prog_cate_id=$13
WHERE prog_entity_id= $1
`

type UpdateprogramentityParams struct {
	ProgEntityID     int32     `db:"prog_entity_id" json:"progEntityId"`
	ProgTitle        string    `db:"prog_title" json:"progTitle"`
	ProgHeadline     string    `db:"prog_headline" json:"progHeadline"`
	ProgType         string    `db:"prog_type" json:"progType"`
	ProgLearningType string    `db:"prog_learning_type" json:"progLearningType"`
	ProgModifiedDate time.Time `db:"prog_modified_date" json:"progModifiedDate"`
	ProgImage        string    `db:"prog_image" json:"progImage"`
	ProgPrice        int32     `db:"prog_price" json:"progPrice"`
	ProgLanguage     string    `db:"prog_language" json:"progLanguage"`
	ProgDuration     int32     `db:"prog_duration" json:"progDuration"`
	ProgDurationType string    `db:"prog_duration_type" json:"progDurationType"`
	ProgTagSkill     string    `db:"prog_tag_skill" json:"progTagSkill"`
	ProgCateID       int32     `db:"prog_cate_id" json:"progCateId"`
}

func (q *Queries) Updateprogramentity(ctx context.Context, arg UpdateprogramentityParams) error {
	_, err := q.db.ExecContext(ctx, updateprogramentity,
		arg.ProgEntityID,
		arg.ProgTitle,
		arg.ProgHeadline,
		arg.ProgType,
		arg.ProgLearningType,
		arg.ProgModifiedDate,
		arg.ProgImage,
		arg.ProgPrice,
		arg.ProgLanguage,
		arg.ProgDuration,
		arg.ProgDurationType,
		arg.ProgTagSkill,
		arg.ProgCateID,
	)
	return err
}

const updateprogramentitydescription = `-- name: Updateprogram_entity_description :exec
UPDATE curriculum.program_entity_description
  set pred_item_learning= $2,
  pred_description = $3
WHERE pred_prog_entity_id= $1
`

type UpdateprogramentitydescriptionParams struct {
	PredProgEntityID int32          `db:"pred_prog_entity_id" json:"predProgEntityId"`
	PredItemLearning sql.NullString `db:"pred_item_learning" json:"predItemLearning"`
	PredDescription  sql.NullString `db:"pred_description" json:"predDescription"`
}

func (q *Queries) Updateprogramentitydescription(ctx context.Context, arg UpdateprogramentitydescriptionParams) error {
	_, err := q.db.ExecContext(ctx, updateprogramentitydescription,
		arg.PredProgEntityID,
		arg.PredItemLearning,
		arg.PredDescription)
	return err
}

const updatescore = `-- name: Updateprogram_entity :exec
UPDATE curriculum.section_detail AS sd
SET secd_score = $2
WHERE sd.secd_sect_id IN (
    SELECT s.sect_id
    FROM curriculum.sections AS s
    WHERE s.sect_prog_entity_id = $1
);
`

type UpdateScoreParams struct {
	SectProgEntityID int32 `db:"sect_prog_entity_id" json:"sectProgEntityId"`
	SecdScore        int32 `db:"secd_score" json:"secdScore"`
}

func (q *Queries) UpdateScore(ctx context.Context, arg UpdateScoreParams) error {
	_, err := q.db.ExecContext(ctx, updatescore,
		arg.SectProgEntityID,
		arg.SecdScore,
	)
	return err
}
