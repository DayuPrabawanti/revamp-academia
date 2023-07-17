-- hr.employee

-- name: GetEmployee :one
SELECT * FROM hr.employee
WHERE emp_entity_id = $1;

-- name: ListEmployees :many
SELECT * FROM hr.employee
ORDER BY emp_emp_number;

-- name: CreateEmployee :one

INSERT INTO hr.employee (emp_entity_id, emp_emp_number, emp_national_id, emp_birth_date, emp_marital_status, emp_gender, emp_hire_date, emp_salaried_flag, emp_vacation_hours, emp_sickleave_hours, emp_current_flag, emp_modified_date, emp_type, emp_joro_id, emp_emp_entity_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
RETURNING emp_entity_id;

-- name: DeleteEmployee :exec
DELETE FROM hr.employee
WHERE emp_entity_id = $1;

UPDATE hr.employee
  set emp_marital_status = $2,
  emp_gender = $3
WHERE emp_entity_id = $1;

-- hr.employee_department_history

-- name: GetEmployeeDepartmentHistory :one
SELECT * FROM hr.employee_department_history
WHERE edhi_id = $1;

-- name: ListEmployeeDepartmentHistory :many
SELECT * FROM hr.employee_department_history
ORDER BY edhi_id;

-- name: CreateEmployeeDepartmentHistory :one

INSERT INTO hr.employee_department_history 
(edhi_id, edhi_entity_id, edhi_start_date, edhi_end_date, edhi_modified_date, edhi_dept_id)
VALUES($1,$2,$3,$4,$5,$6)
RETURNING edhi_id;

-- name: DeleteEmployeeDepartmentHistory :exec
DELETE FROM hr.employee_department_history
WHERE edhi_id = $1;

-- name: UpdateEmployeeDepartmentHistory :exec
UPDATE hr.employee_department_history
  set edhi_start_date = $2,
  edhi_end_date = $3
WHERE edhi_id = $1;

-- hr.department

-- name: GetDepartment :one
SELECT * FROM hr.department
WHERE dept_id = $1;

-- name: ListDepartment :many
SELECT * FROM hr.department
ORDER BY dept_id;

-- name: CreateDepartment :one

INSERT INTO hr.department 
(dept_id, dept_name, dept_modified_date)
VALUES($1,$2,$3)
RETURNING dept_id;

-- name: DeleteDepartment :exec
DELETE FROM hr.department
WHERE dept_id = $1;

-- name: UpdateDepartment :exec
UPDATE hr.department
  set dept_name = $2
WHERE dept_id = $1;

-- hr.employee_client_contract

-- name: GetClientContract :one
SELECT * FROM hr.employee_client_contract
WHERE ecco_id = $1;

-- name: ListClientContract :many
SELECT * FROM hr.employee_client_contract
ORDER BY ecco_id;

-- name: CreateClientContract :one

INSERT INTO hr.employee_client_contract (ecco_id, ecco_entity_id, ecco_contract_no, ecco_contract_date, ecco_start_date, ecco_end_date, ecco_notes, ecco_modified_date, ecco_media_link, ecco_joty_id, ecco_account_manager, ecco_clit_id, ecco_status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
RETURNING ecco_id;

-- name: DeleteClientContract :exec
DELETE FROM hr.employee_client_contract
WHERE ecco_id = $1;

-- name: UpdateClientContract :exec
UPDATE hr.employee_client_contract
  set ecco_start_date = $2,
  ecco_end_date = $3
WHERE ecco_id = $1;

-- hr.employee_pay_history

-- name: GetPayHistory :one
SELECT * FROM hr.employee_pay_history
WHERE ephi_entity_id = $1;

-- name: ListPayHistory :many
SELECT * FROM hr.employee_pay_history
ORDER BY ephi_entity_id;

-- name: CreatePayHistory :one

INSERT INTO hr.employee_pay_history (ephi_entity_id, ephi_rate_change_date, ephi_rate_salary, ephi_pay_frequence, ephi_modified_date) VALUES ($1, $2, $3, $4, $5)
RETURNING ephi_entity_id;

-- name: DeletePayHistory :exec
DELETE FROM hr.employee_pay_history
WHERE ephi_entity_id = $1;

-- name: UpdatePayHistory :exec
UPDATE hr.employee_pay_history
  set ephi_rate_salary = $2,
  ephi_pay_frequence = $3
WHERE ephi_entity_id = $1;