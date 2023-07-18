package dbContext

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"codeid.revamptwo/models/hrsMdl"
)

const listEmployee = `-- name: ListEmployee :many
SELECT emp_entity_id, emp_emp_number, emp_national_id, emp_birth_date, emp_marital_status, emp_gender, emp_hire_date, emp_salaried_flag, emp_vacation_hours, emp_sickleave_hours, emp_current_flag, emp_modified_date, emp_type, emp_joro_id, emp_emp_entity_id FROM hr.employee
ORDER BY emp_entity_id
`

func (q *Queries) GetListEmployee(ctx context.Context) ([]hrsMdl.HrEmployee, error) {
	rows, err := q.db.QueryContext(ctx, listEmployee)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []hrsMdl.HrEmployee
	for rows.Next() {
		var i hrsMdl.HrEmployee
		if err := rows.Scan(&i.EmpEntityID, &i.EmpEmpNumber, &i.EmpNationalID, &i.EmpBirthDate, &i.EmpMaritalStatus, &i.EmpGender, &i.EmpHireDate, &i.EmpSalariedFlag, &i.EmpVacationHours, &i.EmpSickleaveHours, &i.EmpCurrentFlag, &i.EmpModifiedDate, &i.EmpType, &i.EmpJoroID, &i.EmpEmpEntityID); err != nil {
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

// hr.department
func (q *Queries) GetEmployee(ctx context.Context, empEntityID int32) (hrsMdl.HrEmployee, error) {
	row := q.db.QueryRowContext(ctx, getEmployee, empEntityID)
	var i hrsMdl.HrEmployee
	err := row.Scan(&i.EmpEntityID, &i.EmpEmpNumber, &i.EmpNationalID, &i.EmpBirthDate, &i.EmpMaritalStatus, &i.EmpGender, &i.EmpHireDate, &i.EmpSalariedFlag, &i.EmpVacationHours, &i.EmpSickleaveHours, &i.EmpCurrentFlag, &i.EmpModifiedDate, &i.EmpType, &i.EmpJoroID, &i.EmpEmpEntityID)
	return i, err
}

const createEmployee = `-- name: CreateEmployee :one

INSERT INTO hr.employee 
(emp_entity_id, emp_emp_number, emp_national_id, emp_birth_date, emp_marital_status, emp_gender, emp_hire_date, emp_salaried_flag, emp_vacation_hours, emp_sickleave_hours, emp_current_flag, emp_modified_date, emp_type, emp_joro_id, emp_emp_entity_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
RETURNING *
`

type CreateEmployeeParams struct {
	EmpEntityID       int32          `db:"emp_entity_id" json:"empEntityId"`
	EmpEmpNumber      sql.NullString `db:"emp_emp_number" json:"empEmpNumber"`
	EmpNationalID     sql.NullString `db:"emp_national_id" json:"empNationalId"`
	EmpBirthDate      sql.NullTime   `db:"emp_birth_date" json:"empBirthDate"`
	EmpMaritalStatus  sql.NullString `db:"emp_marital_status" json:"empMaritalStatus"`
	EmpGender         sql.NullString `db:"emp_gender" json:"empGender"`
	EmpHireDate       sql.NullTime   `db:"emp_hire_date" json:"empHireDate"`
	EmpSalariedFlag   sql.NullString `db:"emp_salaried_flag" json:"empSalariedFlag"`
	EmpVacationHours  sql.NullInt16  `db:"emp_vacation_hours" json:"empVacationHours"`
	EmpSickleaveHours sql.NullInt16  `db:"emp_sickleave_hours" json:"empSickleaveHours"`
	EmpCurrentFlag    sql.NullInt16  `db:"emp_current_flag" json:"empCurrentFlag"`
	EmpModifiedDate   sql.NullTime   `db:"emp_modified_date" json:"empModifiedDate"`
	EmpType           sql.NullString `db:"emp_type" json:"empType"`
	EmpJoroID         int32          `db:"emp_joro_id" json:"empJoroId"`
	EmpEmpEntityID    int32          `db:"emp_emp_entity_id" json:"empEmpEntityId"`
}

func (q *Queries) CreateEmployee(ctx context.Context, arg CreateEmployeeParams) (*hrsMdl.HrEmployee, *hrsMdl.ResponseError) {
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
		sql.NullTime{Time: time.Now(), Valid: true},
		arg.EmpType,
		arg.EmpJoroID,
		arg.EmpEmpEntityID,
	)
	i := hrsMdl.HrEmployee{}
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
		return nil, &hrsMdl.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &hrsMdl.HrEmployee{
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

const updateEmployee = `-- name: UpdateEmployee :exec
UPDATE hr.employee
  set emp_emp_number_id = $2,
  emp_national_id = $3,
  emp_birthdate = $4,
  emp_maritalstatus = $5,
  emp_gender = $6,
  emp_hiredate = $7,
  emp_salariedflag = $8,
  emp_vacationhours = $9,
  emp_sickleaveshours = $10,
  emp_currentflag = $11,
  emp_modifieddate = $12,
  emp_type = $13,
  emp_joro_id = $14,
  emp_entity_id = $15
WHERE emp_entity_id = $1
`

func (q *Queries) UpdateEmployee(ctx context.Context, arg CreateEmployeeParams) error {
	_, err := q.db.ExecContext(ctx, updateEmployee,
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
		sql.NullTime{Time: time.Now(), Valid: true},
		arg.EmpType,
		arg.EmpJoroID,
		arg.EmpEmpEntityID)
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
