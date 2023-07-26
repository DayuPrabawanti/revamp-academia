package salesController

import (
	"log"
	"net/http"
	"strconv"

	saler "codeid.revampacademy/services/salesServices"
	"github.com/gin-gonic/gin"
)

type ControlMock8 struct {
	serviceMock8 *saler.ServiceMock8
}

func NewMockupApplyController8(serviceMock8 *saler.ServiceMock8) *ControlMock8 {
	return &ControlMock8{
		serviceMock8: serviceMock8,
	}
}

func (controlMock8 ControlMock8) GetListMock8(ctx *gin.Context) {
	response, responerr := controlMock8.serviceMock8.ListMock8Group(ctx)
	if responerr != nil {
		ctx.JSON(responerr.Status, responerr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (controlMock8 ControlMock8) GetMock8Group(ctx *gin.Context) {
	TrpaCodeNumber, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	response, responseErr := controlMock8.serviceMock8.GetMock8Group(ctx, int32(TrpaCodeNumber))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
