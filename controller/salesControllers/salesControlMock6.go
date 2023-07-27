package salescontrollers

import (
	"log"
	"net/http"
	"strconv"

	saler "codeid.revampacademy/service/salesService"
	"github.com/gin-gonic/gin"
)

type ControlMock6 struct {
	serviceMock6 *saler.ServiceMock6
}

func NewControlShoppingCart1(serviceMock6 *saler.ServiceMock6) *ControlMock6 {
	return &ControlMock6{
		serviceMock6: serviceMock6,
	}
}
func (controlMock6 ControlMock6) GetUserIdShoppingCart1(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	response, responseErr := controlMock6.serviceMock6.GetUsersIdShoopingCart1Service(ctx, int64(userId))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (controlMock6 ControlMock6) ListSalesOrderControl(ctx *gin.Context) {
	response, responseErr := controlMock6.serviceMock6.ListSalesOrderService(ctx)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (controlMock6 ControlMock6) GetAccountNumbersControl(ctx *gin.Context) {
	account := ctx.Query("accountNumber")
	response, responseErr := controlMock6.serviceMock6.GetAccountNumbersService(ctx, string(account))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
