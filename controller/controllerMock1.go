package controller

import (
	"log"
	"net/http"
	"strconv"

	"codeid.revampacademy/service"
	"github.com/gin-gonic/gin"
)

type ControllerMock struct {
	serviceMock *service.ServiceMock
}

func NewRepositoryMock(serviceMock *service.ServiceMock) *ControllerMock {
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

func (controllerMock ControllerMock) GetMockupId(ctx *gin.Context) {
	progEntityID, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	response, responseErr := controllerMock.serviceMock.GetMockupId(ctx, int64(progEntityID))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
