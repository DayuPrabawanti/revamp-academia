package hrs

import (
	"database/sql"
	"net/http"

	"codeid.revamptwo/models/hrsMdl"
	"codeid.revamptwo/repositories/hrs/dbContext"
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

func (cr EmployeeRepository) GetListEmployee(ctx *gin.Context) ([]*hrsMdl.HrEmployee, *hrsMdl.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	employees, err := store.GetListEmployee(ctx)

	listEmployee := make([]*hrsMdl.HrEmployee, 0)

	for _, v := range employees {
		employee := &hrsMdl.HrEmployee{
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
		listEmployee = append(listEmployee, employee)
	}

	if err != nil {
		return nil, &hrsMdl.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listEmployee, nil
}

func (cr EmployeeRepository) GetEmployee(ctx *gin.Context, id int64) (*hrsMdl.HrEmployee, *hrsMdl.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	employee, err := store.GetEmployee(ctx, int32(id))

	if err != nil {
		return nil, &hrsMdl.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &employee, nil
}

func (cr EmployeeRepository) CreateEmployee(ctx *gin.Context, employeeParams *dbContext.CreateEmployeeParams) (*hrsMdl.HrEmployee, *hrsMdl.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	employee, err := store.CreateEmployee(ctx, *employeeParams)

	if err != nil {
		return nil, &hrsMdl.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return employee, nil
}

func (cr EmployeeRepository) UpdateEmployee(ctx *gin.Context, employeeParams *dbContext.CreateEmployeeParams) *hrsMdl.ResponseError {

	store := dbContext.New(cr.dbHandler)
	err := store.UpdateEmployee(ctx, *employeeParams)

	if err != nil {
		return &hrsMdl.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &hrsMdl.ResponseError{
		Message: "data has been update",
		Status:  http.StatusOK,
	}
}

func (cr EmployeeRepository) DeleteEmployee(ctx *gin.Context, id int64) *hrsMdl.ResponseError {

	store := dbContext.New(cr.dbHandler)
	err := store.DeleteEmployee(ctx, int32(id))

	if err != nil {
		return &hrsMdl.ResponseError{
			Message: "error when deleted",
			Status:  http.StatusInternalServerError,
		}
	}
	return &hrsMdl.ResponseError{
		Message: "data has been deleted",
		Status:  http.StatusOK,
	}
}
