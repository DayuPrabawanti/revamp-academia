package dbContext

import (
	"context"
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
)

type CreateMergeMock4 struct {
	Createprogram_entityParams
	CreateProgramApplyParam
	CreateProgramApplyProgressParams
}

const createprogram_entity = `-- name: Createprogram_entity :one

INSERT INTO curriculum.program_entity (prog_entity_id, 
prog_title, 
prog_headline, 
prog_type, 
prog_learning_type, 
prog_rating, 
prog_total_trainee, 
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
RETURNING prog_entity_id
`

type Createprogram_entityParams struct {
	ProgTitle sql.NullString `db:"prog_title" json:"progTitle"`
	ProgImage sql.NullString `db:"prog_image" json:"progImage"`
}

func (q *Queries) Createprogram_entity(ctx context.Context, arg Createprogram_entityParams) (*Createprogram_entityParams, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createprogram_entity,
		arg.ProgTitle,
		arg.ProgImage,
	)
	i := Createprogram_entityParams{}
	err := row.Scan(
		&i.ProgTitle,
		&i.ProgImage,
	)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &Createprogram_entityParams{
		ProgTitle: i.ProgTitle,
		ProgImage: i.ProgImage,
	}, nil

}

const createProgramApplys = `-- name: CreateProgramApply :one
INSERT INTO bootcamp.program_apply
(prap_user_entity_id, prap_prog_entity_id, prap_test_score, prap_gpa, prap_iq_test, prap_review, prap_modified_date, prap_status)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING prap_user_entity_id
`

type CreateProgramApplyParam struct {
	PrapStatus sql.NullString `db:"prap_status" json:"prapStatus"`
}

func (q *Queries) CreateProgramApplys(ctx context.Context, arg CreateProgramApplyParam) (*CreateProgramApplyParam, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createProgramApplys,
		arg.PrapStatus,
	)
	i := CreateProgramApplyParam{}
	err := row.Scan(
		&i.PrapStatus,
	)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &CreateProgramApplyParam{
		PrapStatus: i.PrapStatus,
	}, nil

}

const createProgramApplyProgress = `-- name: CreateProgramApplyProgress :one
INSERT INTO bootcamp.program_apply_progress
(parog_id, parog_user_entity_id, parog_prog_entity_id, parog_action_date, parog_modified_date, parog_comment, parog_progress_name, parog_emp_entity_id, parog_status)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING parog_id
`

type CreateProgramApplyProgressParams struct {
	ParogID           int32          `db:"parog_id" json:"parogId"`
	ParogActionDate   sql.NullTime   `db:"parog_action_date" json:"parogActionDate"`
	ParogProgressName sql.NullString `db:"parog_progress_name" json:"parogProgressName"`
}

func (q *Queries) CreateProgramApplyProgress(ctx context.Context, arg CreateProgramApplyProgressParams) (*CreateProgramApplyProgressParams, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createProgramApplyProgress,
		arg.ParogID,
		arg.ParogActionDate,
		arg.ParogProgressName,
	)
	i := CreateProgramApplyProgressParams{}
	err := row.Scan(
		&i.ParogID,
		&i.ParogActionDate,
		&i.ParogProgressName,
	)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &CreateProgramApplyProgressParams{
		ParogID:           i.ParogID,
		ParogActionDate:   i.ParogActionDate,
		ParogProgressName: i.ParogProgressName,
	}, nil
}

const listMock4Grup = `-- name: ListMock4Group :many
select pap.parog_id,pe.prog_image,pe.prog_title,pa.prap_status,pap.parog_action_date,pap.parog_progress_name
from bootcamp.program_apply pa
join bootcamp.program_apply_progress pap
on pa.prap_prog_entity_id = pap.parog_prog_entity_id
join curriculum.program_entity pe
on pa.prap_prog_entity_id = pe.prog_entity_id
`

func (q *Queries) ListMock4Group(ctx context.Context) ([]CreateMergeMock4, error) {
	rows, err := q.db.QueryContext(ctx, listMock4Grup)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CreateMergeMock4
	for rows.Next() {
		var i CreateMergeMock4
		if err := rows.Scan(
			&i.ParogID,
			&i.ProgImage,
			&i.ProgTitle,
			&i.PrapStatus,
			&i.ParogActionDate,
			&i.ParogProgressName,
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

const getMock4Grup = `-- name: GetMock4Group :many
select pap.parog_id,pe.prog_image,pe.prog_title,pa.prap_status,pap.parog_action_date,pap.parog_progress_name
from bootcamp.program_apply pa
join bootcamp.program_apply_progress pap
on pa.prap_prog_entity_id = pap.parog_prog_entity_id
join curriculum.program_entity pe
on pa.prap_prog_entity_id = pe.prog_entity_id
WHERE pap.parog_id = $1
`

func (q *Queries) GetMock4Group(ctx context.Context, ParogID int32) (models.MergeMock4, error) {
	row := q.db.QueryRowContext(ctx, getMock4Grup, ParogID)
	var i models.MergeMock4
	err := row.Scan(
		&i.Progress.ParogID,
		&i.Curriculum.ProgImage,
		&i.Curriculum.ProgTitle,
		&i.Program.PrapStatus,
		&i.Progress.ParogActionDate,
		&i.Progress.ParogProgressName,
	)
	return i, err
}
