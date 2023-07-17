package controller

import (
	"net/http"

	"codeid.revampacademy/service"
	"github.com/gin-gonic/gin"
)

type ControllerMock1 struct {
	serviceMock1 *service.ServiceMock1
}

func NewRepositoryMock1(serviceMock1 *service.ServiceMock1) *ControllerMock1 {
	return &ControllerMock1{
		serviceMock1: serviceMock1,
	}
}

func (controllerMock1 ControllerMock1) GetMockup1(ctx *gin.Context) {
	mockup := ctx.Param("nama")

	response, responseErr := controllerMock1.serviceMock1.GetMockup1(ctx, string(mockup))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (controllerMock1 ControllerMock1) GetListMock1(ctx *gin.Context) {
	response, responseErr := controllerMock1.serviceMock1.GetListMock1(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)

}
