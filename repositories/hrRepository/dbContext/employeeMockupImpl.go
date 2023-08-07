package dbContext

import (
	"context"

	"codeid.revampacademy/models"
)

type EmployeeMockupParams struct {
	User      CreateUsersParams
	General   CreateEmployeeParams
	Salary    CreatePayHistoryParams
	Assigment CreateEmployeeDepartmentHistoryParams
	Client    CreateClientContractParams
}

type EmployeeUpdateMockupParams struct {
	UpdateUser      UpdateUsersParams
	UpdateGeneral   UpdateEmployeeParams
	UpdateSalary    UpdatePayHistoryParams
	UpdateAssigment UpdateEmployeeDepartmentHistoryParams
}

const listEmployeeMockup = `-- name: ListEmployeeMockup :many
SELECT 
emp.emp_entity_id, emp.emp_national_id, 
us.user_first_name, us.user_last_name,
emp.emp_birth_date, emp.emp_hire_date, emp.emp_current_flag
FROM hr.employee emp
JOIN users.users us
ON emp.emp_entity_id = us.user_entity_id
ORDER BY emp.emp_entity_id
`

func (q *Queries) ListEmployeeMockup(ctx context.Context) ([]models.EmployeeMockupList, error) {
	rows, err := q.db.QueryContext(ctx, listEmployeeMockup)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.EmployeeMockupList
	for rows.Next() {
		var i models.EmployeeMockupList
		if err := rows.Scan(
			&i.HrEmployee.EmpEntityID, &i.HrEmployee.EmpNationalID,
			&i.UsersUser.UserFirstName, &i.UsersUser.UserLastName,
			&i.HrEmployee.EmpBirthDate,
			&i.HrEmployee.EmpHireDate,
			&i.HrEmployee.EmpCurrentFlag,
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

const searchEmployee = `-- name: SearchEmployee :many
SELECT 
us.user_first_name, us.user_last_name,
emp.emp_current_flag
FROM hr.employee emp
JOIN users.users us
ON emp.emp_entity_id= us.user_entity_id
WHERE us.user_name like '%' || $1 || '%' AND emp.emp_current_flag = $2
`

func (q *Queries) SearchEmployee(ctx context.Context, userName string, status string) ([]models.EmployeeMockupList, error) {
	rows, err := q.db.QueryContext(ctx, searchEmployee, userName, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var employee []models.EmployeeMockupList
	for rows.Next() {
		var i models.EmployeeMockupList
		if err := rows.Scan(
			&i.UsersUser.UserFirstName, &i.UsersUser.UserLastName,
			&i.HrEmployee.EmpEmpEntityID,
		); err != nil {
			return nil, err
		}
		employee = append(employee, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return employee, nil
}
