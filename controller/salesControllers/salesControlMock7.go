package salescontrollers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	dbcontext "codeid.revampacademy/repositories/salesRepositories/dbContext"
	saser "codeid.revampacademy/service/salesService"
	"github.com/gin-gonic/gin"
)

type ControlMock7 struct {
	serviceMock7 *saser.ServiceMock7
}

func NewControlShoppingCart2(serviceMock7 *saser.ServiceMock7) *ControlMock7 {
	return &ControlMock7{
		serviceMock7: serviceMock7,
	}
}

func (controlMock7 ControlMock7) GetUsersIdShoopingCart2Control(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	response, responseErr := controlMock7.serviceMock7.GetUsersIdShoopingCart2Service(ctx, int64(userId))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (controlMock7 ControlMock7) CreateOrderDetailControl(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create OrderDetail request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var orderParams dbcontext.CreateSales_order_detailParams
	err = json.Unmarshal(body, &orderParams)
	if err != nil {
		log.Println("Error while unmarshaling create OrderDetail request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := controlMock7.serviceMock7.CreateOrderDetailService(ctx, &orderParams)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (controlMock7 ControlMock7) CancelOrderDetailControl(ctx *gin.Context) {

	orderId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseErr := controlMock7.serviceMock7.CancelOrderDetailService(ctx, int64(orderId))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (controlMock7 ControlMock7) GetAccountNumbersMock7Control(ctx *gin.Context) {
	account := ctx.Query("accountNumber")
	response, responseErr := controlMock7.serviceMock7.GetAccountNumberMock7Service(ctx, string(account))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
