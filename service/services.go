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

func (sr SalesService) GetListSales(ctx *gin.Context) ([]*models.SalesSpecialOffer, *models.ResponseError) {
	return sr.salesRepository.GetListSales(ctx)
}

func (sr SalesService) GetListCart_item(ctx *gin.Context, id int64) (*models.SalesCartItem, *models.ResponseError) {
	return sr.salesRepository.GetListCart_item(ctx, id)
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
