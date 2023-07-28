package salesControllers

import (
	"net/http"

	salesservice "codeid.revampacademy/services/salesService"
	"github.com/gin-gonic/gin"
)

type ControllerMock struct {
	serviceMock *salesservice.ServiceMock
}

func NewControllerMock(serviceMock *salesservice.ServiceMock) *ControllerMock {
	return &ControllerMock{
		serviceMock: serviceMock,
	}
}

func (controllerMock ControllerMock) GetMockup1(ctx *gin.Context) {
	mockUp := ctx.Query("nama")

	response, responseErr := controllerMock.serviceMock.GetMockup(ctx, string(mockUp))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (controllerMock ControllerMock) GetListProgram(ctx *gin.Context) {
	mockUp := ctx.Query("nama")
	response, responseErr := controllerMock.serviceMock.GetListProgram(ctx, string(mockUp))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
