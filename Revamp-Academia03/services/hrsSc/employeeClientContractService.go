package hrsSc

import (
	"net/http"

	"codeid.revamptwo/models/hrsMdl"
	"codeid.revamptwo/repositories/hrs"
	"codeid.revamptwo/repositories/hrs/dbContext"
	"github.com/gin-gonic/gin"
)

type EmployeeClientContractService struct {
	employeeClientContractRepository *hrs.EmployeeClientContractRepository
}

func NewEmployeeClientContractService(employeeClientContractRepository *hrs.EmployeeClientContractRepository) *EmployeeClientContractService {
	return &EmployeeClientContractService{
		// struct									parameter
		employeeClientContractRepository: employeeClientContractRepository,
	}
}

// method
func (cs EmployeeClientContractService) GetListEmployeeClientContract(ctx *gin.Context) ([]*hrsMdl.HrEmployeeClientContract, *hrsMdl.ResponseError) {
	return cs.employeeClientContractRepository.GetListEmployeeClientContract(ctx)
}

func (cs EmployeeClientContractService) GetEmployeeClientContract(ctx *gin.Context, id int64) (*hrsMdl.HrEmployeeClientContract, *hrsMdl.ResponseError) {
	return cs.employeeClientContractRepository.GetEmployeeClientContract(ctx, id)
}

// CREATE
func (cs EmployeeClientContractService) CreateEmployeeClientContract(ctx *gin.Context, employeeClientContractParams *dbContext.CreateEmployeeClientContractParams) (*hrsMdl.HrEmployeeClientContract, *hrsMdl.ResponseError) {
	responseErr := validateEmployeeClientContract(employeeClientContractParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return cs.employeeClientContractRepository.CreateEmployeeClientContract(ctx, employeeClientContractParams)
}

// method validation, untuk yang diatas ini
func validateEmployeeClientContract(employeeClientContractParams *dbContext.CreateEmployeeClientContractParams) *hrsMdl.ResponseError {
	if employeeClientContractParams.EccoID == 0 {
		return &hrsMdl.ResponseError{
			Message: "Invalid category id",
			Status:  http.StatusBadRequest,
		}
	}

	if employeeClientContractParams.EccoEntityID == 0 {
		return &hrsMdl.ResponseError{
			Message: "Invalid category name",
			Status:  http.StatusBadRequest,
		}
	}

	return nil
}

// UPDATE
func (cs EmployeeClientContractService) UpdateClientContract(ctx *gin.Context, employeeClientContractParams *dbContext.CreateEmployeeClientContractParams, id int64) *hrsMdl.ResponseError {
	responseErr := validateEmployeeClientContract(employeeClientContractParams)
	if responseErr != nil {
		return responseErr
	}

	return cs.employeeClientContractRepository.UpdateClientContract(ctx, employeeClientContractParams)
}

// DELETE
func (cs EmployeeClientContractService) DeleteClientContract(ctx *gin.Context, id int64) *hrsMdl.ResponseError {
	return cs.employeeClientContractRepository.DeleteClientContract(ctx, id)
}
