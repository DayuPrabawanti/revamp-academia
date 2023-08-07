package hrService

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/hrRepository"
	"codeid.revampacademy/repositories/hrRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type DepartmentHistoryService struct {
	departmentHistoryRepository *hrRepository.DepartmentHistoryRepository
}

func NewDepartmentHistoryService(departmentHistoryRepository *hrRepository.DepartmentHistoryRepository) *DepartmentHistoryService {
	return &DepartmentHistoryService{
		departmentHistoryRepository: departmentHistoryRepository,
	}
}

func (dhs DepartmentHistoryService) GetListDepartmentHistory(ctx *gin.Context) ([]*models.HrEmployeeDepartmentHistory, *models.ResponseError) {
	return dhs.departmentHistoryRepository.GetListDepartmentHistory(ctx)
}

func (dhs DepartmentHistoryService) GetDepartmentHistory(ctx *gin.Context, id int64) (*models.HrEmployeeDepartmentHistory, *models.ResponseError) {
	return dhs.departmentHistoryRepository.GetDepartmentHistory(ctx, id)
}

func (dhs DepartmentHistoryService) CreateDepartmentHistory(ctx *gin.Context, departmentHistoryParams *dbContext.CreateEmployeeDepartmentHistoryParams) (*models.HrEmployeeDepartmentHistory, *models.ResponseError) {
	responseErr := validateDepartmentHistory(departmentHistoryParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return dhs.departmentHistoryRepository.CreateDepartmentHistory(ctx, departmentHistoryParams)
}

func (dhs DepartmentHistoryService) UpdateDepartmentHistory(ctx *gin.Context, departmentHistoryParams *dbContext.UpdateEmployeeDepartmentHistoryParams, id int64) *models.ResponseError {

	return dhs.departmentHistoryRepository.UpdateDepartmentHistory(ctx, departmentHistoryParams)
}

func (dhs DepartmentHistoryService) DeleteDepartmentHistory(ctx *gin.Context, id int64) *models.ResponseError {
	return dhs.departmentHistoryRepository.DeleteDepartmentHistory(ctx, id)
}

func validateDepartmentHistory(departmentHistoryParams *dbContext.CreateEmployeeDepartmentHistoryParams) *models.ResponseError {
	if departmentHistoryParams.EdhiID == 0 {
		return &models.ResponseError{
			Message: "Invalid Edhi id",
			Status:  http.StatusBadRequest,
		}
	}

	if departmentHistoryParams.EdhiEntityID == 0 {
		return &models.ResponseError{
			Message: "Edhi Entity Id Required",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}
