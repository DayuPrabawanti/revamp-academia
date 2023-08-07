package dbContext

import (
	"context"
	"net/http"
	"time"

	"codeid.revampacademy/models"
)

const createEmployeeDepartmentHistory = `-- name: CreateEmployeeDepartmentHistory :one
WITH inserted_entity AS (
	SELECT * FROM hr.employee
	ORDER BY emp_entity_id DESC
	LIMIT 1
  )
INSERT INTO hr.employee_department_history 
(edhi_entity_id, edhi_id, edhi_start_date, edhi_end_date, edhi_modified_date, edhi_dept_id)
SELECT emp_entity_id, $1, $2, $3, Now(), $4
FROM inserted_entity
RETURNING edhi_id, edhi_entity_id, edhi_start_date, edhi_end_date, edhi_modified_date, edhi_dept_id
`

type CreateEmployeeDepartmentHistoryParams struct {
	EdhiID           int32     `db:"edhi_id" json:"edhiId"`
	EdhiEntityID     int32     `db:"edhi_entity_id" json:"edhiEntityId"`
	EdhiStartDate    time.Time `db:"edhi_start_date" json:"edhiStartDate"`
	EdhiEndDate      time.Time `db:"edhi_end_date" json:"edhiEndDate"`
	EdhiModifiedDate time.Time `db:"edhi_modified_date" json:"edhiModifiedDate"`
	EdhiDeptID       int32     `db:"edhi_dept_id" json:"edhiDeptId"`
}

func (q *Queries) CreateEmployeeDepartmentHistory(ctx context.Context, arg CreateEmployeeDepartmentHistoryParams) (*models.HrEmployeeDepartmentHistory, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createEmployeeDepartmentHistory,
		arg.EdhiID,
		arg.EdhiStartDate,
		arg.EdhiEndDate,
		arg.EdhiDeptID,
	)
	i := models.HrEmployeeDepartmentHistory{}
	err := row.Scan(
		&i.EdhiID,
		&i.EdhiEntityID,
		&i.EdhiStartDate,
		&i.EdhiEndDate,
		&i.EdhiModifiedDate,
		&i.EdhiDeptID,
	)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.HrEmployeeDepartmentHistory{
		EdhiID:           i.EdhiID,
		EdhiEntityID:     i.EdhiEntityID,
		EdhiStartDate:    i.EdhiStartDate,
		EdhiEndDate:      i.EdhiEndDate,
		EdhiModifiedDate: i.EdhiModifiedDate,
		EdhiDeptID:       i.EdhiDeptID,
	}, nil
}

const listEmployeeDepartmentHistory = `-- name: ListEmployeeDepartmentHistory :many
SELECT edhi_id, edhi_entity_id, edhi_start_date, edhi_end_date, edhi_modified_date, edhi_dept_id FROM hr.employee_department_history
ORDER BY edhi_id
`

func (q *Queries) ListEmployeeDepartmentHistory(ctx context.Context) ([]models.HrEmployeeDepartmentHistory, error) {
	rows, err := q.db.QueryContext(ctx, listEmployeeDepartmentHistory)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.HrEmployeeDepartmentHistory
	for rows.Next() {
		var i models.HrEmployeeDepartmentHistory
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
func (q *Queries) GetEmployeeDepartmentHistory(ctx context.Context, edhiID int32) (models.HrEmployeeDepartmentHistory, error) {
	row := q.db.QueryRowContext(ctx, getEmployeeDepartmentHistory, edhiID)
	var i models.HrEmployeeDepartmentHistory
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

const updateEmployeeDepartmentHistory = `-- name: UpdateEmployeeDepartmentHistory :exec
UPDATE hr.employee_department_history
  set edhi_entity_id = $2,
  edhi_start_date = $3,
  edhi_end_date = $4,
  edhi_modified_date = Now(),
  edhi_dept_id = $5
WHERE edhi_id = $1
`

type UpdateEmployeeDepartmentHistoryParams struct {
	EdhiID           int32     `db:"edhi_id" json:"edhiId"`
	EdhiEntityID     int32     `db:"edhi_entity_id" json:"edhiEntityId"`
	EdhiStartDate    time.Time `db:"edhi_start_date" json:"edhiStartDate"`
	EdhiEndDate      time.Time `db:"edhi_end_date" json:"edhiEndDate"`
	EdhiModifiedDate time.Time `db:"edhi_modified_date" json:"edhiModifiedDate"`
	EdhiDeptID       int32     `db:"edhi_dept_id" json:"edhiDeptId"`
}

func (q *Queries) UpdateEmployeeDepartmentHistory(ctx context.Context, arg UpdateEmployeeDepartmentHistoryParams) error {
	_, err := q.db.ExecContext(ctx, updateEmployeeDepartmentHistory, arg.EdhiID, arg.EdhiEntityID, arg.EdhiStartDate, arg.EdhiEndDate, arg.EdhiDeptID)
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
