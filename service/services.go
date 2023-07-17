package service

import (
	"net/http"

	models "codeid.revampacademy/models"
	"codeid.revampacademy/repositories"
	dbcontext "codeid.revampacademy/repositories/dbContext"
	"github.com/gin-gonic/gin"
)

type SalesService struct {
	salesRepository *repositories.SalesRepository
}

func NewSalesService(salesRepository *repositories.SalesRepository) *SalesService {
	return &SalesService{
		salesRepository: salesRepository,
	}
}

func (cs SalesService) GetListSales(ctx *gin.Context) ([]*models.SalesSpecialOffer, *models.ResponseError) {
	return cs.salesRepository.GetListSales(ctx)
}

func validateSales(salesParams *dbcontext.CreateSales_special_offerParams) *models.ResponseError {
	if salesParams.SpofID == 0 {
		return &models.ResponseError{
			Message: "Invalid category id",
			Status:  http.StatusBadRequest,
		}
	}

	if salesParams.SpofDescription == "" {
		return &models.ResponseError{
			Message: "Invalid category name",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}
