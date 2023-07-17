package controller

import (
	"net/http"

	"codeid.revampacademy/service"
	"github.com/gin-gonic/gin"
)

type SalesController struct {
	salesService *service.SalesService
}

func NewSalesRepository(categoryService *service.SalesService) *SalesController {
	return &SalesController{
		salesService: categoryService,
	}
}

func (salesController SalesController) GetListSales(ctx *gin.Context) {
	response, responseErr := salesController.salesService.GetListSales(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)

}
