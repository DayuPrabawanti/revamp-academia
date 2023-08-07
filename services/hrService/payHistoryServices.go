package hrService

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/hrRepository"
	"codeid.revampacademy/repositories/hrRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type PayHistoryService struct {
	payHistoryRepository *hrRepository.PayHistoryRepository
}

func NewPayHistoryService(payHistoryRepository *hrRepository.PayHistoryRepository) *PayHistoryService {
	return &PayHistoryService{
		payHistoryRepository: payHistoryRepository,
	}
}

func (phs PayHistoryService) GetListPayHistory(ctx *gin.Context) ([]*models.HrEmployeePayHistory, *models.ResponseError) {
	return phs.payHistoryRepository.GetListPayHistory(ctx)
}

func (phs PayHistoryService) GetPayHistory(ctx *gin.Context, id int64) (*models.HrEmployeePayHistory, *models.ResponseError) {
	return phs.payHistoryRepository.GetPayHistory(ctx, id)
}

func (phs PayHistoryService) CreatePayHistory(ctx *gin.Context, payHistoryParams *dbContext.CreatePayHistoryParams) (*models.HrEmployeePayHistory, *models.ResponseError) {
	responseErr := validatePayHistory(payHistoryParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return phs.payHistoryRepository.CreatePayHistory(ctx, payHistoryParams)
}

func (phs PayHistoryService) UpdatePayHistory(ctx *gin.Context, payHistoryParams *dbContext.UpdatePayHistoryParams, id int64) *models.ResponseError {

	return phs.payHistoryRepository.UpdatePayHistory(ctx, payHistoryParams)
}

func (phs PayHistoryService) DeletePayHistory(ctx *gin.Context, id int64) *models.ResponseError {
	return phs.payHistoryRepository.DeletePayHistory(ctx, id)
}

func validatePayHistory(payHistoryParams *dbContext.CreatePayHistoryParams) *models.ResponseError {
	if payHistoryParams.EphiEntityID == 0 {
		return &models.ResponseError{
			Message: "Invalid Ephi Entity id",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}
