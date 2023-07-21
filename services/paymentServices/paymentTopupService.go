package services

import (
	"codeid.revampacademy/models"
	repositories "codeid.revampacademy/repositories/paymentRepositories"
	"codeid.revampacademy/repositories/paymentRepositories/dbContext"
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

func (pts PaymentTopupService) GetListTopupDetail(ctx *gin.Context) ([]*dbContext.TopupDetail, *models.ResponseError) {
	return pts.paymentTopupRepository.GetListTopupDetail(ctx)
}
