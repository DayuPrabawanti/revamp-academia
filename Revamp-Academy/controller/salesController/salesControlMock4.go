package salesController

import (
	"log"
	"net/http"
	"strconv"

	saler "codeid.revampacademy/services/salesServices"
	"github.com/gin-gonic/gin"
)

type ControlMock4 struct {
	serviceMock4 *saler.ServiceMock4
}

func NewMockupApplyController4(serviceMock4 *saler.ServiceMock4) *ControlMock4 {
	return &ControlMock4{
		serviceMock4: serviceMock4,
	}
}

func (controlMock4 ControlMock4) GetListMock4(ctx *gin.Context) {
	response, responerr := controlMock4.serviceMock4.ListMock4Group(ctx)
	if responerr != nil {
		ctx.JSON(responerr.Status, responerr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (controlMock4 ControlMock4) GetMock4Group(ctx *gin.Context) {
	ParogID, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	response, responseErr := controlMock4.serviceMock4.GetMock4Group(ctx, int32(ParogID))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
