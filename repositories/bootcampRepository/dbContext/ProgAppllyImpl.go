package dbContext

import (
	"context"
	"database/sql"
	"time"

	"codeid.revampacademy/models"
	"codeid.revampacademy/models/features"
)

type GetlistModifiedDate struct {
	PrapModifiedDate time.Time `db:"prap_modified_date" json:"prapModifiedDate"`
}

type UpdateStatus struct {
	PrapTestScore int32  `db:"prap_test_score" json:"prapTestScore"`
	PrapReview    string `db:"prap_review" json:"prapReview"`
	PrapStatus    string `db:"prap_status" json:"prapStatus"`
}

type ProgApply struct {
	PrapStatus       string `db:"prap_status" json:"prapStatus"`
	PrapUserEntityID int32  `db:"prap_user_entity_id" json:"prapUserEntityId"`
	PrapProgEntityID int32  `db:"prap_prog_entity_id" json:"prapProgEntityId"`
}
type StatusProgApply struct {
	PrapUserEntityID int32          `db:"prap_user_entity_id" json:"prapUserEntityId"`
	PrapProgEntityID int32          `db:"prap_prog_entity_id" json:"prapProgEntityId"`
	PrapStatus       sql.NullString `db:"prap_status" json:"prapStatus"`
}
type ProgApplyParams struct {
	PrapUserEntityID int32          `db:"prap_user_entity_id" json:"prapUserEntityId"`
	PrapProgEntityID int32          `db:"prap_prog_entity_id" json:"prapProgEntityId"`
	PrapTestScore    sql.NullInt32  `db:"prap_test_score" json:"prapTestScore"`
	PrapGpa          sql.NullInt32  `db:"prap_gpa" json:"prapGpa"`
	PrapIqTest       sql.NullInt32  `db:"prap_iq_test" json:"prapIqTest"`
	PrapReview       sql.NullString `db:"prap_review" json:"prapReview"`
	PrapModifiedDate sql.NullTime   `db:"prap_modified_date" json:"prapModifiedDate"`
	PrapStatus       string         `db:"prap_status" json:"prapStatus"`
}

const listProgApply = `-- name: listProgApply  :many
SELECT prap_user_entity_id, prap_prog_entity_id, prap_test_score, prap_gpa, prap_iq_test, prap_review, prap_modified_date, prap_status FROM bootcamp.program_apply
ORDER BY prap_user_entity_id
`

func (q *Queries) GetlistProgApply(ctx context.Context) ([]models.BootcampProgramApply, error) {
	rows, err := q.db.QueryContext(ctx, listProgApply)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.BootcampProgramApply
	for rows.Next() {
		var i models.BootcampProgramApply
		if err := rows.Scan(
			&i.PrapUserEntityID,
			&i.PrapProgEntityID,
			&i.PrapTestScore,
			&i.PrapGpa,
			&i.PrapIqTest,
			&i.PrapReview,
			&i.PrapModifiedDate,
			&i.PrapStatus,
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

const UpdateProgApply = `-- name: UpdateProgApply :exec
UPDATE Bootcamp.program_apply
set prap_prog_entity_id = $2,
prap_status = $3
WHERE prap_user_entity_id = $1`

func (q *Queries) UpdateProgApply(ctx context.Context, arg ProgApply) error {
	_, err := q.db.ExecContext(ctx, UpdateProgApply,
		arg.PrapUserEntityID,
		arg.PrapProgEntityID,
		arg.PrapStatus,
	)
	return err
}

// func (q *Queries) UpdateProgApply(ctx context.Context, arg StatusProgApply) error {
// 	// Set the default value "Idle" for BatchStatus if it is not provided or empty
// 	if !arg.PrapStatus.Valid || arg.PrapStatus.String == "" {
// 		arg.PrapStatus.String = "Passed"
// 		arg.PrapStatus.Valid = true
// 	}

// 	_, err := q.db.ExecContext(ctx, UpdateProgApply,
// 		arg.PrapUserEntityID,
// 		arg.PrapProgEntityID,
// 		arg.PrapStatus,
// 	)
// 	return err
// }

const GetProgApply = `-- name: GetProgApply :one
select prap_user_entity_id ,prap_status from bootcamp.program_apply
where prap_user_entity_id =$1
`

func (q *Queries) GetProgApply(ctx context.Context, prapUserEntityId int32) (ProgApply, error) {
	row := q.db.QueryRowContext(ctx, GetProgApply, prapUserEntityId)
	var i ProgApply
	err := row.Scan(&i.PrapProgEntityID, &i.PrapStatus)
	return i, err
}

const GetTestScore = `-- name: GetTestScore :one
SELECT prap_test_score,prap_status,prap_review
FROM bootcamp.program_apply
WHERE prap_user_entity_id = $1
`

func (q *Queries) GetTestScore(ctx context.Context, PrapUserEntityID int32) (UpdateStatus, error) {
	row := q.db.QueryRowContext(ctx, GetTestScore, PrapUserEntityID)
	var i UpdateStatus
	err := row.Scan(
		&i.PrapTestScore,
		&i.PrapStatus,
		&i.PrapReview,
	)
	return i, err
}

const GetReview = `-- name: GetReview :one
SELECT prap_test_score,prap_status,prap_review
FROM bootcamp.program_apply
WHERE prap_user_entity_id = $1
`

func (q *Queries) GetReview(ctx context.Context, PrapUserEntityID int32) (UpdateStatus, error) {
	row := q.db.QueryRowContext(ctx, GetTestScore, PrapUserEntityID)
	var i UpdateStatus
	err := row.Scan(
		&i.PrapTestScore,
		&i.PrapStatus,
		&i.PrapReview,
	)
	return i, err
}

const GetStatus = `-- name: GetStatus :one
SELECT prap_test_score,prap_status,prap_review
FROM bootcamp.program_apply
WHERE prap_user_entity_id = $1
`

func (q *Queries) GetStatus(ctx context.Context, PrapUserEntityID int32) (UpdateStatus, error) {
	row := q.db.QueryRowContext(ctx, GetTestScore, PrapUserEntityID)
	var i UpdateStatus
	err := row.Scan(
		&i.PrapTestScore,
		&i.PrapStatus,
		&i.PrapReview,
	)
	return i, err
}

const listProgApplyStatus = `-- name: listProgApplyStatus :one
SELECT prap_user_entity_id, prap_prog_entity_id, prap_test_score, prap_gpa, prap_iq_test, prap_review, prap_modified_date, prap_status FROM bootcamp.program_apply
where prap_status = $1
`

func (q *Queries) GetlistProgApplyStatus(ctx context.Context, status string) ([]ProgApplyParams, error) {
	rows, err := q.db.QueryContext(ctx, listProgApplyStatus, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProgApplyParams
	for rows.Next() {
		var i ProgApplyParams
		if err := rows.Scan(
			&i.PrapUserEntityID,
			&i.PrapProgEntityID,
			&i.PrapTestScore,
			&i.PrapGpa,
			&i.PrapIqTest,
			&i.PrapReview,
			&i.PrapModifiedDate,
			&i.PrapStatus,
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

const Listprogapplyfiltering = `-- name: Listprogapplyfiltering :one
SELECT prap_user_entity_id, prap_prog_entity_id, prap_test_score, prap_gpa, prap_iq_test, prap_review, prap_modified_date, prap_status FROM bootcamp.program_apply
where prap_status = $1
`

func (q *Queries) GetlistProgApplyfiltering(ctx context.Context, status string) ([]ProgApplyParams, error) {
	rows, err := q.db.QueryContext(ctx, Listprogapplyfiltering, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProgApplyParams
	for rows.Next() {
		var i ProgApplyParams
		if err := rows.Scan(
			&i.PrapUserEntityID,
			&i.PrapProgEntityID,
			&i.PrapTestScore,
			&i.PrapGpa,
			&i.PrapIqTest,
			&i.PrapReview,
			&i.PrapModifiedDate,
			&i.PrapStatus,
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

const Listprogapplycontract = `-- name: Listprogapplycontract :one
SELECT prap_user_entity_id, prap_prog_entity_id, prap_test_score, prap_gpa, prap_iq_test, prap_review, prap_modified_date, prap_status FROM bootcamp.program_apply
where prap_status = $1
`

func (q *Queries) GetlistProgApplycontract(ctx context.Context, status string) ([]ProgApplyParams, error) {
	rows, err := q.db.QueryContext(ctx, Listprogapplycontract, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProgApplyParams
	for rows.Next() {
		var i ProgApplyParams
		if err := rows.Scan(
			&i.PrapUserEntityID,
			&i.PrapProgEntityID,
			&i.PrapTestScore,
			&i.PrapGpa,
			&i.PrapIqTest,
			&i.PrapReview,
			&i.PrapModifiedDate,
			&i.PrapStatus,
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

const Listprogapplyfailed = `-- name: Listprogapplyfailed :one
SELECT prap_user_entity_id, prap_prog_entity_id, prap_test_score, prap_gpa, prap_iq_test, prap_review, prap_modified_date, prap_status FROM bootcamp.program_apply
where prap_status = $1
`

func (q *Queries) GetlistProgApplyfailed(ctx context.Context, status string) ([]ProgApplyParams, error) {
	rows, err := q.db.QueryContext(ctx, Listprogapplyfailed, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProgApplyParams
	for rows.Next() {
		var i ProgApplyParams
		if err := rows.Scan(
			&i.PrapUserEntityID,
			&i.PrapProgEntityID,
			&i.PrapTestScore,
			&i.PrapGpa,
			&i.PrapIqTest,
			&i.PrapReview,
			&i.PrapModifiedDate,
			&i.PrapStatus,
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

const Listprogapplyidle = `-- name: Listprogapplyidle :one
SELECT prap_user_entity_id, prap_prog_entity_id, prap_test_score, prap_gpa, prap_iq_test, prap_review, prap_modified_date, prap_status FROM bootcamp.program_apply
where prap_status = $1
`

func (q *Queries) GetlistProgApplyidle(ctx context.Context, status string) ([]ProgApplyParams, error) {
	rows, err := q.db.QueryContext(ctx, Listprogapplyidle, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProgApplyParams
	for rows.Next() {
		var i ProgApplyParams
		if err := rows.Scan(
			&i.PrapUserEntityID,
			&i.PrapProgEntityID,
			&i.PrapTestScore,
			&i.PrapGpa,
			&i.PrapIqTest,
			&i.PrapReview,
			&i.PrapModifiedDate,
			&i.PrapStatus,
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

const UpdateTestScore = `-- name: UpdateTestScore :exec
UPDATE Bootcamp.program_apply
set prap_status = $2,
prap_user_entity_id = $3
WHERE prap_prog_entity_id = $1`

func (q *Queries) UpdateTestScore(ctx context.Context, arg UpdateStatus, id int32) error {
	_, err := q.db.ExecContext(ctx, UpdateTestScore, id,
		arg.PrapTestScore,
		arg.PrapStatus,
		arg.PrapReview,
	)
	return err
}

const UpdatePrapStatus = `-- name: UpdatePrapStatus :exec
UPDATE Bootcamp.program_apply
set prap_status = $2,
prap_user_entity_id = $3
WHERE prap_prog_entity_id = $1`

func (q *Queries) UpdatePrapStatus(ctx context.Context, arg UpdateStatus, id int32) error {
	_, err := q.db.ExecContext(ctx, UpdatePrapStatus, id,
		arg.PrapTestScore,
		arg.PrapStatus,
		arg.PrapReview,
	)
	return err
}

const UpdatePrapReview = `-- name: UpdatePrapReview :exec
UPDATE Bootcamp.program_apply
SET prap_test_score = $2,
    prap_review = $3,
    prap_status = CASE
        WHEN $2 < 25 THEN 'Failed'
        WHEN $2 >= 25 AND $2 < 50 THEN 'Running'
        WHEN $2 >= 50 THEN 'Passed'
        ELSE prap_status
    END
WHERE prap_user_entity_id = $1`

func (q *Queries) UpdatePrapReview(ctx context.Context, arg UpdateStatus, id int32) error {
	_, err := q.db.ExecContext(ctx, UpdatePrapReview, id,
		arg.PrapTestScore,
		arg.PrapReview,
	)

	return err
}

const getlistModifieddate = `-- name: GetlistModifieddate :one
select prap_user_entity_id, 
prap_prog_entity_id, 
prap_test_score, 
prap_gpa, 
prap_iq_test, 
prap_review,
prap_modified_date, 
prap_status  
from bootcamp.program_apply 
where extract(month from prap_modified_date) = $1 and extract(year from prap_modified_date) = $2
`

func (q *Queries) GetlistModifiedDate(ctx context.Context, metadata *features.Metadata) ([]ProgApplyParams, error) {
	rows, err := q.db.QueryContext(ctx, getlistModifieddate, metadata.Month, metadata.Year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProgApplyParams
	for rows.Next() {
		var i ProgApplyParams
		if err := rows.Scan(
			&i.PrapUserEntityID,
			&i.PrapProgEntityID,
			&i.PrapTestScore,
			&i.PrapGpa,
			&i.PrapIqTest,
			&i.PrapReview,
			&i.PrapModifiedDate,
			&i.PrapStatus,
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
