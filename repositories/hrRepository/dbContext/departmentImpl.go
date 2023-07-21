package dbContext

import (
	"context"
	"net/http"
	"time"

	"codeid.revampacademy/models"
)

const createDepartment = `-- name: CreateDepartment :one

INSERT INTO hr.department 
(dept_id, dept_name, dept_modified_date)
VALUES($1,$2,$3)
RETURNING *
`

type CreateDepartmentParams struct {
	DeptID           int32     `db:"dept_id" json:"deptId"`
	DeptName         string    `db:"dept_name" json:"deptName"`
	DeptModifiedDate time.Time `db:"dept_modified_date" json:"deptModifiedDate"`
}

func (q *Queries) CreateDepartment(ctx context.Context, arg CreateDepartmentParams) (*models.HrDepartment, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createDepartment,
		arg.DeptID,
		arg.DeptName,
		arg.DeptModifiedDate,
	)
	i := models.HrDepartment{}
	err := row.Scan(
		&i.DeptID,
		&i.DeptName,
		&i.DeptModifiedDate,
	)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.HrDepartment{
		DeptID:           i.DeptID,
		DeptName:         i.DeptName,
		DeptModifiedDate: i.DeptModifiedDate,
	}, nil
}

const listDepartment = `-- name: ListDepartment :many
SELECT dept_id, dept_name, dept_modified_date FROM hr.department
ORDER BY dept_id
`

func (q *Queries) ListDepartment(ctx context.Context) ([]models.HrDepartment, error) {
	rows, err := q.db.QueryContext(ctx, listDepartment)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.HrDepartment
	for rows.Next() {
		var i models.HrDepartment
		if err := rows.Scan(&i.DeptID, &i.DeptName, &i.DeptModifiedDate); err != nil {
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

const getDepartment = `-- name: GetDepartment :one

SELECT dept_id, dept_name, dept_modified_date FROM hr.department
WHERE dept_id = $1
`

// hr.department
func (q *Queries) GetDepartment(ctx context.Context, deptID int32) (models.HrDepartment, error) {
	row := q.db.QueryRowContext(ctx, getDepartment, deptID)
	var i models.HrDepartment
	err := row.Scan(&i.DeptID, &i.DeptName, &i.DeptModifiedDate)
	return i, err
}

const updateDepartment = `-- name: UpdateDepartment :exec
UPDATE hr.department
  set dept_name = $2,
  dept_modified_date = $3
WHERE dept_id = $1
`

func (q *Queries) UpdateDepartment(ctx context.Context, arg CreateDepartmentParams) error {
	_, err := q.db.ExecContext(ctx, updateDepartment, arg.DeptID, arg.DeptName, arg.DeptModifiedDate)
	return err
}

const deleteDepartment = `-- name: DeleteDepartment :exec
DELETE FROM hr.department
WHERE dept_id = $1
`

func (q *Queries) DeleteDepartment(ctx context.Context, deptID int32) error {
	_, err := q.db.ExecContext(ctx, deleteDepartment, deptID)
	return err
}
