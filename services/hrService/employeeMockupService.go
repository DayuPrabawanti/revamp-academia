package hrService

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/hrRepository"
	"codeid.revampacademy/repositories/hrRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type EmployeeMockupService struct {
	employeeMockupRepository *hrRepository.EmployeeMockupRepository
}

func NewEmployeeMockupService(employeeMockupRepository *hrRepository.EmployeeMockupRepository) *EmployeeMockupService {
	return &EmployeeMockupService{
		employeeMockupRepository: employeeMockupRepository,
	}
}

func (es EmployeeMockupService) ListEmployeeMockup(ctx *gin.Context) ([]*models.EmployeeMockupList, *models.ResponseError) {
	return es.employeeMockupRepository.ListEmployeeMockup(ctx)
}

func (tdms EmployeeMockupService) SearchEmployee(ctx *gin.Context, userName string, status string) ([]models.EmployeeMockupList, *models.ResponseError) {
	return tdms.employeeMockupRepository.SearchEmployee(ctx, userName, status)
}

func (cs *EmployeeMockupService) EmployeeMockup(ctx *gin.Context, userParam *dbContext.CreateUsersParams, empParam *dbContext.CreateEmployeeParams, salaryParam *dbContext.CreatePayHistoryParams, assigmentParam *dbContext.CreateEmployeeDepartmentHistoryParams, clientParam *dbContext.CreateClientContractParams) (*models.EmployeeMockupModel, *models.ResponseError) {

	return cs.employeeMockupRepository.CreateEmployeeMockup(ctx, userParam, empParam, salaryParam, assigmentParam, clientParam)
}

func (cs *EmployeeMockupService) UpdateEmployeeMockup(ctx *gin.Context, userParam *dbContext.UpdateUsersParams, empParam *dbContext.UpdateEmployeeParams, salaryParam *dbContext.UpdatePayHistoryParams, assigmentParam *dbContext.UpdateEmployeeDepartmentHistoryParams, id int64) *models.ResponseError {

	return cs.employeeMockupRepository.UpdateEmployeeMockup(ctx, userParam, empParam, salaryParam, assigmentParam, id)
}
