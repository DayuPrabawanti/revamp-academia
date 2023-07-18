package controller

import (
	"net/http"

	"codeid.revampacademy/services"
	"github.com/gin-gonic/gin"
)

type ControllerMock struct {
	serviceMock *services.ServiceMock
}

func NewRepositoryMock(serviceMock *services.ServiceMock) *ControllerMock {
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

// func (controllerMock1 ControllerMock1) GetListMock1(ctx *gin.Context) {
// 	response, responseErr := controllerMock1.serviceMock1.GetListMock1(ctx)

// 	if responseErr != nil {
// 		ctx.JSON(responseErr.Status, responseErr)
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, response)

// }
