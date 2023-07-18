package hrsSc

import (
	"net/http"

	"codeid.revamptwo/models/hrsMdl"
	"codeid.revamptwo/repositories/hrs"
	"codeid.revamptwo/repositories/hrs/dbContext"
	"github.com/gin-gonic/gin"
)

type EmployeeDepartmentHistoryService struct {
	departmentHistoryRepository *hrs.EmployeeDepartmentHistoryRepository
}

func NewEmployeeDepartmentHistoryService(departmentHistoryRepository *hrs.EmployeeDepartmentHistoryRepository) *EmployeeDepartmentHistoryService {
	return &EmployeeDepartmentHistoryService{
		// struct				parameter
		departmentHistoryRepository: departmentHistoryRepository,
	}
}

// method
func (cs EmployeeDepartmentHistoryService) ListEmployeeDepartmentHistory(ctx *gin.Context) ([]*hrsMdl.HrEmployeeDepartmentHistory, *hrsMdl.ResponseError) {
	return cs.departmentHistoryRepository.ListEmployeeDepartmentHistory(ctx)
}

func (cs EmployeeDepartmentHistoryService) GetEmployeeDepartmentHistory(ctx *gin.Context, id int64) (*hrsMdl.HrEmployeeDepartmentHistory, *hrsMdl.ResponseError) {
	return cs.departmentHistoryRepository.GetEmployeeDepartmentHistory(ctx, id)
}

// CREATE
func (cs EmployeeDepartmentHistoryService) CreateEmployeeDepartmentHistory(ctx *gin.Context, departmentHistoryParams *dbContext.CreateEmployeeDepartmentHistoryParams) (*hrsMdl.HrEmployeeDepartmentHistory, *hrsMdl.ResponseError) {
	responseErr := validateDepartmentHistory(departmentHistoryParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return cs.departmentHistoryRepository.CreateEmployeeDepartmentHistory(ctx, departmentHistoryParams)
}

// UPDATE
func (cs EmployeeDepartmentHistoryService) UpdateEmployeeDepartmentHistory(ctx *gin.Context, departmentHistoryParams *dbContext.UpdateEmployeeDepartmentHistoryParams, id int64) *hrsMdl.ResponseError {
	responseErr := validateDepartmentHistory(departmentHistoryParams)
	if responseErr != nil {
		return responseErr
	}

	return cs.departmentHistoryRepository.UpdateEmployeeDepartmentHistory(ctx, departmentHistoryParams)
}

// DELETE
func (cs EmployeeDepartmentHistoryService) DeleteEmployeeDepartmentHistory(ctx *gin.Context, id int64) *hrsMdl.ResponseError {
	return cs.departmentHistoryRepository.DeleteEmployeeDepartmentHistory(ctx, id)
}

// method validation, untuk yang diatas ini
func validateDepartmentHistory(departmentHistoryParams *dbContext.UpdateEmployeeDepartmentHistoryParams) *hrsMdl.ResponseError {
	if departmentHistoryParams.EdhiID == 0 {
		return &hrsMdl.ResponseError{
			Message: "Invalid category id",
			Status:  http.StatusBadRequest,
		}
	}

	if departmentHistoryParams.EdhiEntityID == 0 {
		return &hrsMdl.ResponseError{
			Message: "Invalid category name",
			Status:  http.StatusBadRequest,
		}
	}

	return nil
}
