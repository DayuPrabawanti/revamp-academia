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

func (cs *EmployeeMockupService) EmployeeMockup(ctx *gin.Context, employeeMockup *dbContext.EmployeeMockupParams) (*models.EmployeeMockupModel, *models.ResponseError) {

	return cs.employeeMockupRepository.CreateEmployeeMockup(ctx, employeeMockup)
}
