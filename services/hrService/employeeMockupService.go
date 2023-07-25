package hrService

import (
	"net/http"

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

func (cs *EmployeeMockupService) EmployeeMockup(ctx *gin.Context, employeeMockup *dbContext.EmployeeMockupParams) (*models.EmployeeMockupModel, *models.ResponseError) {
	responseErr := validateEmployeeMockup(employeeMockup)
	if responseErr != nil {
		return nil, responseErr
	}

	return cs.employeeMockupRepository.CreateEmployeeMockup(ctx, employeeMockup)
}

func validateEmployeeMockup(employeemockupParams *dbContext.EmployeeMockupParams) *models.ResponseError {
	if employeemockupParams.General.EmpEntityID != 0 {
		return &models.ResponseError{
			Message: "ID Pengguna tidak valid",
			Status:  http.StatusBadRequest,
		}
	}

	return nil
}
