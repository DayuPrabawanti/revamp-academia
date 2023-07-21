package salesControllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

func (controllerMock ControllerMock) GetListProgramMock(ctx *gin.Context) {
	mockUp := ctx.Query("orderby")
	response, responseErr := controllerMock.serviceMock.GetListProgramMock(ctx, string(mockUp))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
