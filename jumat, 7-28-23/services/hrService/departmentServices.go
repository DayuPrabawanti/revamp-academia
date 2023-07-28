package hrService

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/hrRepository"
	"codeid.revampacademy/repositories/hrRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type DepartmentService struct {
	departmentRepository *hrRepository.DepartmentRepository
}

func NewDepartmentService(departmentRepository *hrRepository.DepartmentRepository) *DepartmentService {
	return &DepartmentService{
		departmentRepository: departmentRepository,
	}
}

func (ds DepartmentService) GetListDepartment(ctx *gin.Context) ([]*models.HrDepartment, *models.ResponseError) {
	return ds.departmentRepository.GetListDepartment(ctx)
}

func (ds DepartmentService) GetDepartment(ctx *gin.Context, id int64) (*models.HrDepartment, *models.ResponseError) {
	return ds.departmentRepository.GetDepartment(ctx, id)
}

func (ds DepartmentService) CreateDepartment(ctx *gin.Context, departmentParams *dbContext.CreateDepartmentParams) (*models.HrDepartment, *models.ResponseError) {
	responseErr := validateDepartment(departmentParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return ds.departmentRepository.CreateDepartment(ctx, departmentParams)
}

func (ds DepartmentService) UpdateDepartment(ctx *gin.Context, departmentParams *dbContext.CreateDepartmentParams, id int64) *models.ResponseError {
	responseErr := validateDepartment(departmentParams)
	if responseErr != nil {
		return responseErr
	}

	return ds.departmentRepository.UpdateDepartment(ctx, departmentParams)
}

func (ds DepartmentService) DeleteDepartment(ctx *gin.Context, id int64) *models.ResponseError {
	return ds.departmentRepository.DeleteDepartment(ctx, id)
}

func validateDepartment(departmentParams *dbContext.CreateDepartmentParams) *models.ResponseError {
	if departmentParams.DeptID == 0 {
		return &models.ResponseError{
			Message: "Invalid department id",
			Status:  http.StatusBadRequest,
		}
	}

	if departmentParams.DeptName == "" {
		return &models.ResponseError{
			Message: "Department Name Required",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}
