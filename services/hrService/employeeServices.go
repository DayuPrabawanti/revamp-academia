package hrService

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/hrRepository"
	"codeid.revampacademy/repositories/hrRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type EmployeeService struct {
	employeeRepository *hrRepository.EmployeeRepository
}

func NewEmployeeService(employeeRepository *hrRepository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{
		employeeRepository: employeeRepository,
	}
}

func (es EmployeeService) GetListEmployee(ctx *gin.Context) ([]*models.HrEmployee, *models.ResponseError) {
	return es.employeeRepository.GetListEmployee(ctx)
}

func (es EmployeeService) GetEmployee(ctx *gin.Context, id int64) (*models.HrEmployee, *models.ResponseError) {
	return es.employeeRepository.GetEmployee(ctx, id)
}

func (es EmployeeService) CreateEmployee(ctx *gin.Context, employeeParams *dbContext.CreateEmployeeParams) (*models.HrEmployee, *models.ResponseError) {
	responseErr := validateEmployee(employeeParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return es.employeeRepository.CreateEmployee(ctx, employeeParams)
}

func (es EmployeeService) UpdateEmployee(ctx *gin.Context, employeeParams *dbContext.UpdateEmployeeParams, id int64) *models.ResponseError {

	return es.employeeRepository.UpdateEmployee(ctx, employeeParams)
}

func (es EmployeeService) DeleteEmployee(ctx *gin.Context, id int64) *models.ResponseError {
	return es.employeeRepository.DeleteEmployee(ctx, id)
}

func validateEmployee(employeeParams *dbContext.CreateEmployeeParams) *models.ResponseError {
	if employeeParams.EmpEntityID == 0 {
		return &models.ResponseError{
			Message: "Invalid EmpEntityID",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}

// USER
func (cs EmployeeService) CreateUser(ctx *gin.Context, userParams *dbContext.CreateUsersParams) (*models.UsersUser, *models.ResponseError) {
	responseErr := validateUser(userParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return cs.employeeRepository.CreateUser(ctx, userParams)
}

func (cs EmployeeService) UpdateUser(ctx *gin.Context, userParams *dbContext.UpdateUsersParams, id int64) *models.ResponseError {

	return cs.employeeRepository.UpdateUser(ctx, userParams)
}

func validateUser(userParams *dbContext.CreateUsersParams) *models.ResponseError {
	if userParams.UserEntityID == 0 {
		return &models.ResponseError{
			Message: "Invalid User id",
			Status:  http.StatusBadRequest,
		}
	}
	return nil

}
