package services

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/dbContext"
	repositories "codeid.revampacademy/repositories/paymentRepositories"
	"github.com/gin-gonic/gin"
)

type PaymentTopupService struct {
	paymentTopupRepository *repositories.PaymentTopupRepository
}

func NewPaymentTopupService(paymentTopupRepository *repositories.PaymentTopupRepository) *PaymentTopupService {
	return &PaymentTopupService{
		paymentTopupRepository: paymentTopupRepository,
	}
}

func (pts PaymentTopupService) GetTopupDetail(ctx *gin.Context, sourceBankEntityID int32, targetFintechEntityID int32) (*dbContext.TopupDetail, *models.ResponseError) {
	return pts.paymentTopupRepository.GetTopupDetail(ctx, sourceBankEntityID, targetFintechEntityID)
}
