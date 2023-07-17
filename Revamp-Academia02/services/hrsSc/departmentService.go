package hrsSc

import (
	"net/http"

	"codeid.revamptwo/models/hrsMdl"
	"codeid.revamptwo/repositories/hrs"
	"codeid.revamptwo/repositories/hrs/dbContext"
	"github.com/gin-gonic/gin"
)

type DepartmentService struct {
	departmentRepository *hrs.DepartmentRepository
}

func NewDepartmentService(departmentRepository *hrs.DepartmentRepository) *DepartmentService {
	return &DepartmentService{
		// struct				parameter
		departmentRepository: departmentRepository,
	}
}

// method
func (cs DepartmentService) GetListDepartment(ctx *gin.Context) ([]*hrsMdl.HrDepartment, *hrsMdl.ResponseError) {
	return cs.departmentRepository.GetListDepartment(ctx)
}

func (cs DepartmentService) GetDepartment(ctx *gin.Context, id int64) (*hrsMdl.HrDepartment, *hrsMdl.ResponseError) {
	return cs.departmentRepository.GetDepartment(ctx, id)
}

// CREATE
func (cs DepartmentService) CreateDepartment(ctx *gin.Context, departmentParams *dbContext.CreateDepartmentParams) (*hrsMdl.HrDepartment, *hrsMdl.ResponseError) {
	responseErr := validateDepartment(departmentParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return cs.departmentRepository.CreateDepartment(ctx, departmentParams)
}

// UPDATE
func (cs DepartmentService) UpdateDepartment(ctx *gin.Context, departmentParams *dbContext.CreateDepartmentParams, id int64) *hrsMdl.ResponseError {
	responseErr := validateDepartment(departmentParams)
	if responseErr != nil {
		return responseErr
	}

	return cs.departmentRepository.UpdateDepartment(ctx, departmentParams)
}

// DELETE
func (cs DepartmentService) DeleteDepartment(ctx *gin.Context, id int64) *hrsMdl.ResponseError {
	return cs.departmentRepository.DeleteDepartment(ctx, id)
}

// method validation, untuk yang diatas ini
func validateDepartment(departmentParams *dbContext.CreateDepartmentParams) *hrsMdl.ResponseError {
	if departmentParams.DeptID.Valid == false {
		return &hrsMdl.ResponseError{
			Message: "Invalid category id",
			Status:  http.StatusBadRequest,
		}
	}

	if departmentParams.DeptName.Valid == false {
		return &hrsMdl.ResponseError{
			Message: "Invalid category name",
			Status:  http.StatusBadRequest,
		}
	}

	return nil
}
