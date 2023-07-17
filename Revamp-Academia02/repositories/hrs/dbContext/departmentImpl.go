package dbContext

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"codeid.revamptwo/models/hrsMdl"
)

const listDepartment = `-- name: ListDepartment :many
SELECT dept_id, dept_name, dept_modified_date FROM hr.department
ORDER BY dept_id
`

func (q *Queries) ListDepartment(ctx context.Context) ([]hrsMdl.HrDepartment, error) {
	rows, err := q.db.QueryContext(ctx, listDepartment)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []hrsMdl.HrDepartment
	for rows.Next() {
		var i hrsMdl.HrDepartment
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
func (q *Queries) GetDepartment(ctx context.Context, deptID int32) (hrsMdl.HrDepartment, error) {
	row := q.db.QueryRowContext(ctx, getDepartment, deptID)
	var i hrsMdl.HrDepartment
	err := row.Scan(&i.DeptID, &i.DeptName, &i.DeptModifiedDate)
	return i, err
}

const createDepartment = `-- name: CreateDepartment :one

INSERT INTO hr.department 
(dept_id, dept_name, dept_modified_date)
VALUES($1,$2,$3)
RETURNING *
`

type CreateDepartmentParams struct {
	DeptID           sql.NullInt32  `db:"dept_id" json:"deptId"`
	DeptName         sql.NullString `db:"dept_name" json:"deptName"`
	DeptModifiedDate sql.NullTime   `db:"dept_modified_date" json:"deptModifiedDate"`
}

func (q *Queries) CreateDepartment(ctx context.Context, arg CreateDepartmentParams) (*hrsMdl.HrDepartment, *hrsMdl.ResponseError) {
	row := q.db.QueryRowContext(ctx, createDepartment,
		arg.DeptID,
		arg.DeptName,
		arg.DeptModifiedDate,
	)
	i := hrsMdl.HrDepartment{}
	err := row.Scan(
		&i.DeptID,
		&i.DeptName,
		&i.DeptModifiedDate,
	)

	if err != nil {
		return nil, &hrsMdl.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &hrsMdl.HrDepartment{
		DeptID:           i.DeptID,
		DeptName:         i.DeptName,
		DeptModifiedDate: sql.NullTime{Time: time.Now(), Valid: true},
	}, nil
}

const updateDepartment = `-- name: UpdateDepartment :exec
UPDATE hr.department
  set dept_name = $2,
  dept_modified_date = $3
WHERE dept_id = $1
`

func (q *Queries) UpdateDepartment(ctx context.Context, arg CreateDepartmentParams) error {
	_, err := q.db.ExecContext(ctx, updateDepartment,
		arg.DeptID,
		arg.DeptName,
		sql.NullTime{Time: time.Now(), Valid: true})
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
