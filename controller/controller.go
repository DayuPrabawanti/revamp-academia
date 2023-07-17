package controller

import (
	"log"
	"net/http"
	"strconv"

	"codeid.revampacademy/service"
	"github.com/gin-gonic/gin"
)

type SalesController struct {
	salesService *service.SalesService
}

func NewSalesRepository(salesService *service.SalesService) *SalesController {
	return &SalesController{
		salesService: salesService,
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

func (salesController SalesController) GetListCart_item(ctx *gin.Context) {
	caitId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := salesController.salesService.GetListCart_item(ctx, int64(caitId))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
