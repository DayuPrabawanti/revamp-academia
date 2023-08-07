package hrRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/hrRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type EmployeeRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewEmployeeRepository(dbHandler *sql.DB) *EmployeeRepository {
	return &EmployeeRepository{
		dbHandler: dbHandler,
	}
}

func (er EmployeeRepository) GetListEmployee(ctx *gin.Context) ([]*models.HrEmployee, *models.ResponseError) {

	store := dbContext.New(er.dbHandler)
	employees, err := store.ListEmployees(ctx)

	listEmployees := make([]*models.HrEmployee, 0)

	for _, v := range employees {
		employee := &models.HrEmployee{
			EmpEntityID:       v.EmpEntityID,
			EmpEmpNumber:      v.EmpEmpNumber,
			EmpNationalID:     v.EmpNationalID,
			EmpBirthDate:      v.EmpBirthDate,
			EmpMaritalStatus:  v.EmpMaritalStatus,
			EmpGender:         v.EmpGender,
			EmpHireDate:       v.EmpHireDate,
			EmpSalariedFlag:   v.EmpSalariedFlag,
			EmpVacationHours:  v.EmpVacationHours,
			EmpSickleaveHours: v.EmpSickleaveHours,
			EmpCurrentFlag:    v.EmpCurrentFlag,
			EmpModifiedDate:   v.EmpModifiedDate,
			EmpType:           v.EmpType,
			EmpJoroID:         v.EmpJoroID,
			EmpEmpEntityID:    v.EmpEmpEntityID,
		}
		listEmployees = append(listEmployees, employee)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listEmployees, nil
}

func (er EmployeeRepository) GetEmployee(ctx *gin.Context, id int64) (*models.HrEmployee, *models.ResponseError) {

	store := dbContext.New(er.dbHandler)
	employee, err := store.GetEmployee(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &employee, nil
}

func (er EmployeeRepository) CreateEmployee(ctx *gin.Context, employeeParams *dbContext.CreateEmployeeParams) (*models.HrEmployee, *models.ResponseError) {

	store := dbContext.New(er.dbHandler)
	employee, err := store.CreateEmployee(ctx, *employeeParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return employee, nil
}

func (er EmployeeRepository) UpdateEmployee(ctx *gin.Context, employeeParams *dbContext.UpdateEmployeeParams) *models.ResponseError {

	store := dbContext.New(er.dbHandler)
	err := store.UpdateEmployee(ctx, *employeeParams)

	if err != nil {
		return &models.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been update",
		Status:  http.StatusOK,
	}
}

func (er EmployeeRepository) DeleteEmployee(ctx *gin.Context, id int64) *models.ResponseError {

	store := dbContext.New(er.dbHandler)
	err := store.DeleteEmployee(ctx, int32(id))

	if err != nil {
		return &models.ResponseError{
			Message: "error when delete",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been deleted",
		Status:  http.StatusOK,
	}
}

func (cr EmployeeRepository) CreateUser(ctx *gin.Context, userParams *dbContext.CreateUsersParams) (*models.UsersUser, *models.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	user, err := store.CreateUsers(ctx, *userParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return user, nil
}

func (cr EmployeeRepository) UpdateUser(ctx *gin.Context, userParams *dbContext.UpdateUsersParams) *models.ResponseError {

	store := dbContext.New(cr.dbHandler)
	err := store.UpdateUser(ctx, *userParams)

	if err != nil {
		return &models.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been update",
		Status:  http.StatusOK,
	}
}
