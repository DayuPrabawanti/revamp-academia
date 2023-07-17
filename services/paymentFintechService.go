package services

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories"
	"github.com/gin-gonic/gin"
)

type PaymentFintechService struct {
	paymentFintechRepository *repositories.PaymentFintechRepository
}

func NewPaymentFintechService(paymentFintechRepository *repositories.PaymentFintechRepository) *PaymentFintechService {
	return &PaymentFintechService{
		paymentFintechRepository: paymentFintechRepository,
	}
}

func (pbs PaymentFintechService) GetListPaymentFintech(ctx *gin.Context) ([]*models.PaymentFintech, *models.ResponseError) {
	return pbs.paymentFintechRepository.GetListPaymentFintech(ctx)
}
