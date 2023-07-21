package salesServices

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/salesRepositories"
	"github.com/gin-gonic/gin"
)

type FintechService struct {
	fintechRepository *salesRepositories.FintechRepository
}

func NewFintechService(fintechRepository *salesRepositories.FintechRepository) *FintechService {
	return &FintechService{
		fintechRepository: fintechRepository,
	}
}

func (cs FintechService) GetPaymentFintech(ctx *gin.Context, id int32) (*models.PaymentFintech, *models.ResponseError) {
	return cs.fintechRepository.GetPaymentFintech(ctx, id)
}
