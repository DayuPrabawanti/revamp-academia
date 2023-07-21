package dbContext

import (
	"context"
	"net/http"
	"time"

	"codeid.revampacademy/models"
)

const createEmployee = `-- name: CreateEmployee :one

INSERT INTO hr.employee (emp_entity_id, emp_emp_number, emp_national_id, emp_birth_date, emp_marital_status, emp_gender, emp_hire_date, emp_salaried_flag, emp_vacation_hours, emp_sickleave_hours, emp_current_flag, emp_modified_date, emp_type, emp_joro_id, emp_emp_entity_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
RETURNING *
`

type CreateEmployeeParams struct {
	EmpEntityID       int32     `db:"emp_entity_id" json:"empEntityId"`
	EmpEmpNumber      string    `db:"emp_emp_number" json:"empEmpNumber"`
	EmpNationalID     string    `db:"emp_national_id" json:"empNationalId"`
	EmpBirthDate      time.Time `db:"emp_birth_date" json:"empBirthDate"`
	EmpMaritalStatus  string    `db:"emp_marital_status" json:"empMaritalStatus"`
	EmpGender         string    `db:"emp_gender" json:"empGender"`
	EmpHireDate       time.Time `db:"emp_hire_date" json:"empHireDate"`
	EmpSalariedFlag   string    `db:"emp_salaried_flag" json:"empSalariedFlag"`
	EmpVacationHours  int16     `db:"emp_vacation_hours" json:"empVacationHours"`
	EmpSickleaveHours int16     `db:"emp_sickleave_hours" json:"empSickleaveHours"`
	EmpCurrentFlag    int16     `db:"emp_current_flag" json:"empCurrentFlag"`
	EmpModifiedDate   time.Time `db:"emp_modified_date" json:"empModifiedDate"`
	EmpType           string    `db:"emp_type" json:"empType"`
	EmpJoroID         int32     `db:"emp_joro_id" json:"empJoroId"`
	EmpEmpEntityID    int32     `db:"emp_emp_entity_id" json:"empEmpEntityId"`
}

func (q *Queries) CreateEmployee(ctx context.Context, arg CreateEmployeeParams) (*models.HrEmployee, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createEmployee,
		arg.EmpEntityID,
		arg.EmpEmpNumber,
		arg.EmpNationalID,
		arg.EmpBirthDate,
		arg.EmpMaritalStatus,
		arg.EmpGender,
		arg.EmpHireDate,
		arg.EmpSalariedFlag,
		arg.EmpVacationHours,
		arg.EmpSickleaveHours,
		arg.EmpCurrentFlag,
		arg.EmpModifiedDate,
		arg.EmpType,
		arg.EmpJoroID,
		arg.EmpEmpEntityID,
	)
	i := models.HrEmployee{}
	err := row.Scan(
		&i.EmpEntityID,
		&i.EmpEmpNumber,
		&i.EmpNationalID,
		&i.EmpBirthDate,
		&i.EmpMaritalStatus,
		&i.EmpGender,
		&i.EmpHireDate,
		&i.EmpSalariedFlag,
		&i.EmpVacationHours,
		&i.EmpSickleaveHours,
		&i.EmpCurrentFlag,
		&i.EmpModifiedDate,
		&i.EmpType,
		&i.EmpJoroID,
		&i.EmpEmpEntityID,
	)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.HrEmployee{
		EmpEntityID:       i.EmpEntityID,
		EmpEmpNumber:      i.EmpEmpNumber,
		EmpNationalID:     i.EmpNationalID,
		EmpBirthDate:      i.EmpBirthDate,
		EmpMaritalStatus:  i.EmpMaritalStatus,
		EmpGender:         i.EmpGender,
		EmpHireDate:       i.EmpHireDate,
		EmpSalariedFlag:   i.EmpSalariedFlag,
		EmpVacationHours:  i.EmpVacationHours,
		EmpSickleaveHours: i.EmpSickleaveHours,
		EmpCurrentFlag:    i.EmpCurrentFlag,
		EmpModifiedDate:   i.EmpModifiedDate,
		EmpType:           i.EmpType,
		EmpJoroID:         i.EmpJoroID,
		EmpEmpEntityID:    i.EmpEmpEntityID,
	}, nil
}

const listEmployees = `-- name: ListEmployees :many
SELECT emp_entity_id, emp_emp_number, emp_national_id, emp_birth_date, emp_marital_status, emp_gender, emp_hire_date, emp_salaried_flag, emp_vacation_hours, emp_sickleave_hours, emp_current_flag, emp_modified_date, emp_type, emp_joro_id, emp_emp_entity_id FROM hr.employee
ORDER BY emp_emp_number
`

func (q *Queries) ListEmployees(ctx context.Context) ([]models.HrEmployee, error) {
	rows, err := q.db.QueryContext(ctx, listEmployees)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.HrEmployee
	for rows.Next() {
		var i models.HrEmployee
		if err := rows.Scan(
			&i.EmpEntityID,
			&i.EmpEmpNumber,
			&i.EmpNationalID,
			&i.EmpBirthDate,
			&i.EmpMaritalStatus,
			&i.EmpGender,
			&i.EmpHireDate,
			&i.EmpSalariedFlag,
			&i.EmpVacationHours,
			&i.EmpSickleaveHours,
			&i.EmpCurrentFlag,
			&i.EmpModifiedDate,
			&i.EmpType,
			&i.EmpJoroID,
			&i.EmpEmpEntityID,
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

const getEmployee = `-- name: GetEmployee :one

SELECT emp_entity_id, emp_emp_number, emp_national_id, emp_birth_date, emp_marital_status, emp_gender, emp_hire_date, emp_salaried_flag, emp_vacation_hours, emp_sickleave_hours, emp_current_flag, emp_modified_date, emp_type, emp_joro_id, emp_emp_entity_id FROM hr.employee
WHERE emp_entity_id = $1
`

// hr.employee
func (q *Queries) GetEmployee(ctx context.Context, empEntityID int32) (models.HrEmployee, error) {
	row := q.db.QueryRowContext(ctx, getEmployee, empEntityID)
	var i models.HrEmployee
	err := row.Scan(
		&i.EmpEntityID,
		&i.EmpEmpNumber,
		&i.EmpNationalID,
		&i.EmpBirthDate,
		&i.EmpMaritalStatus,
		&i.EmpGender,
		&i.EmpHireDate,
		&i.EmpSalariedFlag,
		&i.EmpVacationHours,
		&i.EmpSickleaveHours,
		&i.EmpCurrentFlag,
		&i.EmpModifiedDate,
		&i.EmpType,
		&i.EmpJoroID,
		&i.EmpEmpEntityID,
	)
	return i, err
}

const updateEmployee = `-- name: UpdateEmployee :exec
UPDATE hr.employee
  set emp_emp_number = $2,
  emp_national_id = $3,
  emp_birth_date = $4,
  emp_marital_status = $5,
  emp_gender = $6,
  emp_hire_date = $7,
  emp_salaried_flag = $8,
  emp_vacation_hours = $9,
  emp_sickleave_hours = $10,
  emp_current_flag = $11,
  emp_modified_date = $12,
  emp_type = $13,
  emp_joro_id = $14,
  emp_emp_entity_id = 15
WHERE emp_entity_id = $1
`

func (q *Queries) UpdateEmployee(ctx context.Context, arg CreateEmployeeParams) error {
	_, err := q.db.ExecContext(ctx, updateEmployee, arg.EmpEntityID, arg.EmpEmpNumber, arg.EmpNationalID, arg.EmpBirthDate, arg.EmpMaritalStatus, arg.EmpGender, arg.EmpHireDate, arg.EmpSalariedFlag, arg.EmpVacationHours, arg.EmpSickleaveHours, arg.EmpCurrentFlag, arg.EmpModifiedDate, arg.EmpType, arg.EmpJoroID, arg.EmpEmpEntityID)
	return err
}

const deleteEmployee = `-- name: DeleteEmployee :exec
DELETE FROM hr.employee
WHERE emp_entity_id = $1
`

func (q *Queries) DeleteEmployee(ctx context.Context, empEntityID int32) error {
	_, err := q.db.ExecContext(ctx, deleteEmployee, empEntityID)
	return err
}
