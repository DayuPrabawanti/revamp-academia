package hrsSc

import (
	"net/http"

	"codeid.revamptwo/models/hrsMdl"
	"codeid.revamptwo/repositories/hrs"
	"codeid.revamptwo/repositories/hrs/dbContext"
	"github.com/gin-gonic/gin"
)

type EmployeeService struct {
	employeeRepository *hrs.EmployeeRepository
}

func NewEmployeeService(employeeRepository *hrs.EmployeeRepository) *EmployeeService {
	return &EmployeeService{
		// struct				parameter
		employeeRepository: employeeRepository,
	}
}

// method
func (cs EmployeeService) GetListEmployee(ctx *gin.Context) ([]*hrsMdl.HrEmployee, *hrsMdl.ResponseError) {
	return cs.employeeRepository.GetListEmployee(ctx)
}

func (cs EmployeeService) GetEmployee(ctx *gin.Context, id int64) (*hrsMdl.HrEmployee, *hrsMdl.ResponseError) {
	return cs.employeeRepository.GetEmployee(ctx, id)
}

// CREATE
func (cs EmployeeService) CreateEmployee(ctx *gin.Context, employeeParams *dbContext.CreateEmployeeParams) (*hrsMdl.HrEmployee, *hrsMdl.ResponseError) {
	responseErr := validateEmployee(employeeParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return cs.employeeRepository.CreateEmployee(ctx, employeeParams)
}

// UPDATE
func (cs EmployeeService) UpdateEmployee(ctx *gin.Context, employeeParams *dbContext.CreateEmployeeParams, id int64) *hrsMdl.ResponseError {
	responseErr := validateEmployee(employeeParams)
	if responseErr != nil {
		return responseErr
	}

	return cs.employeeRepository.UpdateEmployee(ctx, employeeParams)
}

// DELETE
func (cs EmployeeService) DeleteEmployee(ctx *gin.Context, id int64) *hrsMdl.ResponseError {
	return cs.employeeRepository.DeleteEmployee(ctx, id)
}

// method validation, untuk yang diatas ini
func validateEmployee(employeeParams *dbContext.CreateEmployeeParams) *hrsMdl.ResponseError {
	if employeeParams.EmpEntityID == 0 {
		return &hrsMdl.ResponseError{
			Message: "Invalid category id",
			Status:  http.StatusBadRequest,
		}
	}

	if employeeParams.EmpEmpNumber.Valid == false {
		return &hrsMdl.ResponseError{
			Message: "Invalid category name",
			Status:  http.StatusBadRequest,
		}
	}

	return nil
}
