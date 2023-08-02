package dbContext

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"codeid.revampacademy/models"
)

const createEmployee = `-- name: CreateEmployee :one

WITH inserted_entity AS (
	INSERT INTO hr.employee 
	(emp_entity_id, emp_emp_number, emp_national_id, emp_birth_date, emp_marital_status, emp_gender, emp_hire_date, emp_salaried_flag, emp_vacation_hours, emp_sickleave_hours, emp_current_flag, emp_modified_date, emp_type, emp_joro_id, emp_emp_entity_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
	RETURNING *
)

INSERT INTO hr.employee_client_contract (ecco_entity_id)
SELECT emp_entity_id
	FROM inserted_entity
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
	EmpJoroID         sql.NullInt32  `db:"emp_joro_id" json:"empJoroId"`
	EmpEmpEntityID    sql.NullInt32  `db:"emp_emp_entity_id" json:"empEmpEntityId"`
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
		sql.NullTime{Time: time.Now(), Valid: true},
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

type UpdateEmployeeParams struct {
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
	EmpJoroID         sql.NullInt32  `db:"emp_joro_id" json:"empJoroId"`
	EmpEmpEntityID    sql.NullInt32  `db:"emp_emp_entity_id" json:"empEmpEntityId"`
}

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
		arg.EmpModifiedDate,
		arg.EmpType,
		arg.EmpJoroID,
		arg.EmpEmpEntityID,
	)
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

type CreateUsersParams struct {
	UserEntityID       int32          `db:"user_entity_id" json:"userEntityId"`
	UserName           sql.NullString `db:"user_name" json:"userName"`
	UserPassword       sql.NullString `db:"user_password" json:"userPassword"`
	UserFirstName      sql.NullString `db:"user_first_name" json:"userFirstName"`
	UserLastName       sql.NullString `db:"user_last_name" json:"userLastName"`
	UserBirthDate      sql.NullTime   `db:"user_birth_date" json:"userBirthDate"`
	UserEmailPromotion sql.NullInt32  `db:"user_email_promotion" json:"userEmailPromotion"`
	UserDemographic    sql.NullString `db:"user_demographic" json:"userDemographic"`
	UserModifiedDate   sql.NullTime   `db:"user_modified_date" json:"userModifiedDate"`
	UserPhoto          sql.NullString `db:"user_photo" json:"userPhoto"`
	UserCurrentRole    sql.NullInt32  `db:"user_current_role" json:"userCurrentRole"`
}

const createUsers = `-- name: CreateUsers :one

WITH inserted_entity AS (
  INSERT INTO users.business_entity 
  (entity_modified_date)
  VALUES (Now())
  RETURNING entity_id
)
INSERT INTO users.users 
(user_entity_id, user_name, user_password, user_first_name, 
user_last_name, user_birth_date, user_email_promotion, user_demographic, 
user_modified_date, user_photo, user_current_role)
SELECT  entity_id, $1, $2, $3, $4, $5, $6, $7, $8, $9, $10 FROM inserted_entity
RETURNING user_entity_id, user_name, user_password, user_first_name, 
user_last_name, user_birth_date, user_email_promotion, user_demographic, 
user_modified_date, user_photo, user_current_role
`

func (q *Queries) CreateUsers(ctx context.Context, arg CreateUsersParams) (*models.UsersUser, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createUsers,
		arg.UserName,
		arg.UserPassword,
		arg.UserFirstName,
		arg.UserLastName,
		arg.UserBirthDate,
		arg.UserEmailPromotion,
		arg.UserDemographic,
		sql.NullTime{Time: time.Now(), Valid: true},
		arg.UserPhoto,
		sql.NullInt64{Int64: 1, Valid: true},
	)
	i := models.UsersUser{}
	err := row.Scan(
		&i.UserEntityID,
		&i.UserName,
		&i.UserPassword,
		&i.UserFirstName,
		&i.UserLastName,
		&i.UserBirthDate,
		&i.UserEmailPromotion,
		&i.UserDemographic,
		&i.UserModifiedDate,
		&i.UserPhoto,
		&i.UserCurrentRole,
	)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.UsersUser{
		UserEntityID:       i.UserEntityID,
		UserName:           i.UserName,
		UserPassword:       i.UserPassword,
		UserFirstName:      i.UserFirstName,
		UserLastName:       i.UserLastName,
		UserBirthDate:      i.UserBirthDate,
		UserEmailPromotion: i.UserEmailPromotion,
		UserDemographic:    i.UserDemographic,
		UserModifiedDate:   i.UserModifiedDate,
		UserPhoto:          i.UserPhoto,
		UserCurrentRole:    i.UserCurrentRole,
	}, nil
}

const updateUsers = `-- name: UpdateUsers :exec
UPDATE users.users
  set user_name = $2,
  user_password=$3,
  user_first_name= $4,
  user_last_name =$5,
  user_birth_date=$6,
  user_email_promotion=$7,
  user_demographic=$8,
  user_modified_date=$9,
  user_photo=$10,
  user_current_role=$11
WHERE user_entity_id = $1
`

type UpdateUsersParams struct {
	UserEntityID       int32          `db:"user_entity_id" json:"userEntityId"`
	UserName           sql.NullString `db:"user_name" json:"userName"`
	UserPassword       sql.NullString `db:"user_password" json:"userPassword"`
	UserFirstName      sql.NullString `db:"user_first_name" json:"userFirstName"`
	UserLastName       sql.NullString `db:"user_last_name" json:"userLastName"`
	UserBirthDate      sql.NullTime   `db:"user_birth_date" json:"userBirthDate"`
	UserEmailPromotion sql.NullInt32  `db:"user_email_promotion" json:"userEmailPromotion"`
	UserDemographic    sql.NullString `db:"user_demographic" json:"userDemographic"`
	UserModifiedDate   sql.NullTime   `db:"user_modified_date" json:"userModifiedDate"`
	UserPhoto          sql.NullString `db:"user_photo" json:"userPhoto"`
	UserCurrentRole    sql.NullInt32  `db:"user_current_role" json:"userCurrentRole"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg CreateUsersParams) error {
	_, err := q.db.ExecContext(ctx, updateUsers,
		arg.UserEntityID,
		arg.UserName,
		arg.UserPassword,
		arg.UserFirstName,
		arg.UserLastName,
		arg.UserBirthDate,
		arg.UserEmailPromotion,
		arg.UserDemographic,
		sql.NullTime{Time: time.Now(), Valid: true},
		arg.UserPhoto,
		arg.UserCurrentRole)
	return err
}
