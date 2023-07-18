package dbContext

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"codeid.revamptwo/models/hrsMdl"
)

const listEmployeeDepartmentHistory = `-- name: ListEmployeeDepartmentHistory :many
SELECT edhi_id, edhi_entity_id, edhi_start_date, edhi_end_date, edhi_modified_date, edhi_dept_id FROM hr.employee_department_history
ORDER BY edhi_id
`

func (q *Queries) ListEmployeeDepartmentHistory(ctx context.Context) ([]hrsMdl.HrEmployeeDepartmentHistory, error) {
	rows, err := q.db.QueryContext(ctx, listEmployeeDepartmentHistory)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []hrsMdl.HrEmployeeDepartmentHistory
	for rows.Next() {
		var i hrsMdl.HrEmployeeDepartmentHistory
		if err := rows.Scan(
			&i.EdhiID,
			&i.EdhiEntityID,
			&i.EdhiStartDate,
			&i.EdhiEndDate,
			&i.EdhiModifiedDate,
			&i.EdhiDeptID,
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

const getEmployeeDepartmentHistory = `-- name: GetEmployeeDepartmentHistory :one

SELECT edhi_id, edhi_entity_id, edhi_start_date, edhi_end_date, edhi_modified_date, edhi_dept_id FROM hr.employee_department_history
WHERE edhi_id = $1
`

// hr.employee_department_history
func (q *Queries) GetEmployeeDepartmentHistory(ctx context.Context, edhiID int32) (hrsMdl.HrEmployeeDepartmentHistory, error) {
	row := q.db.QueryRowContext(ctx, getEmployeeDepartmentHistory, edhiID)
	var i hrsMdl.HrEmployeeDepartmentHistory
	err := row.Scan(
		&i.EdhiID,
		&i.EdhiEntityID,
		&i.EdhiStartDate,
		&i.EdhiEndDate,
		&i.EdhiModifiedDate,
		&i.EdhiDeptID,
	)
	return i, err
}

const createEmployeeDepartmentHistory = `-- name: CreateEmployeeDepartmentHistory :one

INSERT INTO hr.employee_department_history 
(edhi_id, edhi_entity_id, edhi_start_date, edhi_end_date, edhi_modified_date, edhi_dept_id)
VALUES($1,$2,$3,$4,$5,$6)
RETURNING *
`

type CreateEmployeeDepartmentHistoryParams struct {
	EdhiID           int32         `db:"edhi_id" json:"edhiId"`
	EdhiEntityID     int32         `db:"edhi_entity_id" json:"edhiEntityId"`
	EdhiStartDate    sql.NullTime  `db:"edhi_start_date" json:"edhiStartDate"`
	EdhiEndDate      sql.NullTime  `db:"edhi_end_date" json:"edhiEndDate"`
	EdhiModifiedDate sql.NullTime  `db:"edhi_modified_date" json:"edhiModifiedDate"`
	EdhiDeptID       sql.NullInt32 `db:"edhi_dept_id" json:"edhiDeptId"`
}

func (q *Queries) CreateEmployeeDepartmentHistory(ctx context.Context, arg CreateEmployeeDepartmentHistoryParams) (*hrsMdl.HrEmployeeDepartmentHistory, *hrsMdl.ResponseError) {
	row := q.db.QueryRowContext(ctx, createEmployeeDepartmentHistory,
		arg.EdhiID,
		arg.EdhiEntityID,
		arg.EdhiStartDate,
		arg.EdhiEndDate,
		sql.NullTime{Time: time.Now(), Valid: true},
		arg.EdhiDeptID,
	)
	i := hrsMdl.HrEmployeeDepartmentHistory{}
	err := row.Scan(
		&i.EdhiID,
		&i.EdhiEntityID,
		&i.EdhiStartDate,
		&i.EdhiEndDate,
		&i.EdhiModifiedDate,
		&i.EdhiDeptID,
	)

	if err != nil {
		return nil, &hrsMdl.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &hrsMdl.HrEmployeeDepartmentHistory{
		EdhiID:           i.EdhiID,
		EdhiEntityID:     i.EdhiEntityID,
		EdhiStartDate:    i.EdhiStartDate,
		EdhiEndDate:      i.EdhiEndDate,
		EdhiModifiedDate: i.EdhiModifiedDate,
		EdhiDeptID:       i.EdhiDeptID,
	}, nil
}

const updateEmployeeDepartmentHistory = `-- name: UpdateEmployeeDepartmentHistory :exec
UPDATE hr.employee_department_history
  set edhi_entity_id = $2
  edhi_start_date = $3,
  edhi_end_date = $4,
  edhi_modified_date = $5,
  edhi_dept_id = $6
WHERE edhi_id = $1
`

func (q *Queries) UpdateEmployeeDepartmentHistory(ctx context.Context, arg CreateEmployeeDepartmentHistoryParams) error {
	_, err := q.db.ExecContext(ctx, updateEmployeeDepartmentHistory,
		arg.EdhiID,
		arg.EdhiEntityID,
		arg.EdhiStartDate,
		arg.EdhiEndDate,
		sql.NullTime{Time: time.Now(), Valid: true},
		arg.EdhiDeptID,
	)
	return err
}

const deleteEmployeeDepartmentHistory = `-- name: DeleteEmployeeDepartmentHistory :exec
DELETE FROM hr.employee_department_history
WHERE edhi_id = $1
`

func (q *Queries) DeleteEmployeeDepartmentHistory(ctx context.Context, edhiID int32) error {
	_, err := q.db.ExecContext(ctx, deleteEmployeeDepartmentHistory, edhiID)
	return err
}
