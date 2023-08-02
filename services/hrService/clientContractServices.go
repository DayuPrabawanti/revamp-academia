package hrService

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/hrRepository"
	"codeid.revampacademy/repositories/hrRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type ClientContractService struct {
	clientContractRepository *hrRepository.ClientContractRepository
}

func NewClientContractService(clientContractRepository *hrRepository.ClientContractRepository) *ClientContractService {
	return &ClientContractService{
		clientContractRepository: clientContractRepository,
	}
}

func (ccs ClientContractService) GetListClientContract(ctx *gin.Context) ([]*models.HrEmployeeClientContract, *models.ResponseError) {
	return ccs.clientContractRepository.GetListClientContract(ctx)
}

func (ccs ClientContractService) GetClientContract(ctx *gin.Context, id int64) (*models.HrEmployeeClientContract, *models.ResponseError) {
	return ccs.clientContractRepository.GetClientContract(ctx, id)
}

func (ccs ClientContractService) CreateClientContract(ctx *gin.Context, clientContractParams *dbContext.CreateClientContractParams) (*models.HrEmployeeClientContract, *models.ResponseError) {
	responseErr := validateClientContract(clientContractParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return ccs.clientContractRepository.CreateClientContract(ctx, clientContractParams)
}

func (ccs ClientContractService) UpdateClientContract(ctx *gin.Context, clientContractParams *dbContext.CreateClientContractParams, id int64) *models.ResponseError {
	responseErr := validateClientContract(clientContractParams)
	if responseErr != nil {
		return responseErr
	}

	return ccs.clientContractRepository.UpdateClientContract(ctx, clientContractParams)
}

func (ccs ClientContractService) DeleteClientContract(ctx *gin.Context, id int64) *models.ResponseError {
	return ccs.clientContractRepository.DeleteClientContract(ctx, id)
}

func validateClientContract(clientContractParams *dbContext.CreateClientContractParams) *models.ResponseError {
	if clientContractParams.EccoID == 0 {
		return &models.ResponseError{
			Message: "Invalid ecco id",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}
