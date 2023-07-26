package dbContext

import (
	"context"
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
)

type CreateMergeMock2 struct {
	CreateProgramApplyParams
	CreateBatchParams
	CreateInstructorProgramParams
}

const createProgramApply = `-- name: CreateProgramApply :one
INSERT INTO bootcamp.program_apply
(prap_user_entity_id, prap_prog_entity_id, prap_test_score, prap_gpa, prap_iq_test, prap_review, prap_modified_date, prap_status)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING prap_user_entity_id
`

type CreateProgramApplyParams struct {
	PrapReview sql.NullString `db:"prap_review" json:"prapReview"`
}

const createBatch = `-- name: CreateBatch :one
INSERT INTO bootcamp.batch
(batch_id, batch_entity_id, batch_name, batch_description, batch_start_date, batch_end_date, batch_reason, batch_type, batch_modified_date, batch_status, batch_pic_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING batch_id
`

func (q *Queries) CreateProgramApply(ctx context.Context, arg CreateProgramApplyParams) (*CreateProgramApplyParams, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createProgramApply,
		arg.PrapReview,
	)
	i := CreateProgramApplyParams{}
	err := row.Scan(
		&i.PrapReview,
	)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &CreateProgramApplyParams{
		PrapReview: i.PrapReview,
	}, nil
}

type CreateBatchParams struct {
	BatchDescription sql.NullString `db:"batch_description" json:"batchDescription"`
}

func (q *Queries) CreateBatch(ctx context.Context, arg CreateBatchParams) (*CreateBatchParams, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createBatch,
		arg.BatchDescription,
	)
	i := CreateBatchParams{}
	err := row.Scan(
		&i.BatchDescription,
	)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &CreateBatchParams{
		BatchDescription: i.BatchDescription,
	}, nil
}

const createInstructorProgram = `-- name: CreateInstructorProgram :one
INSERT INTO bootcamp.instructor_programs
(batch_id, inpro_entity_id, inpro_emp_entity_id, inpro_modified_date)
VALUES ($1, $2, $3, $4)
RETURNING batch_id
`

type CreateInstructorProgramParams struct {
	InproEntityID     int32        `db:"inpro_entity_id" json:"inproEntityId"`
	InproEmpEntityID  int32        `db:"inpro_emp_entity_id" json:"inproEmpEntityId"`
	InproModifiedDate sql.NullTime `db:"inpro_modified_date" json:"inproModifiedDate"`
}

func (q *Queries) CreateInstructorProgram(ctx context.Context, arg CreateInstructorProgramParams) (*CreateInstructorProgramParams, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createInstructorProgram,
		arg.InproEntityID,
		arg.InproEmpEntityID,
		arg.InproModifiedDate,
	)
	i := CreateInstructorProgramParams{}
	err := row.Scan(
		&i.InproEntityID,
		&i.InproEmpEntityID,
		&i.InproModifiedDate,
	)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &CreateInstructorProgramParams{
		InproEntityID:     i.InproEntityID,
		InproEmpEntityID:  i.InproEmpEntityID,
		InproModifiedDate: i.InproModifiedDate,
	}, nil
}

const listBootcampGrup = `-- name: ListBootcampGroup :many
select pa.prap_review,b.batch_description,ip.inpro_entity_id,ip.inpro_emp_entity_id
,ip.inpro_modified_date
from bootcamp.batch b
join bootcamp.instructor_programs ip 
on b.batch_id = ip.batch_id
join curriculum.program_entity pe
on b.batch_id = pe.prog_entity_id
join bootcamp.program_apply pa
on pe.prog_entity_id = pa.prap_prog_entity_id
`

func (q *Queries) ListBootcampGroup(ctx context.Context) ([]CreateMergeMock2, error) {
	rows, err := q.db.QueryContext(ctx, listBootcampGrup)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CreateMergeMock2
	for rows.Next() {
		var i CreateMergeMock2
		if err := rows.Scan(
			&i.PrapReview,
			&i.BatchDescription,
			&i.InproEntityID,
			&i.InproEmpEntityID,
			&i.InproModifiedDate,
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
