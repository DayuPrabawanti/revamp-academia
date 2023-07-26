package salesController

import (
	"net/http"

	saler "codeid.revampacademy/services/salesServices"
	"github.com/gin-gonic/gin"
)

type ControlMock2 struct {
	serviceMock2 *saler.ServiceMock2
}

func NewMockupApplyController2(serviceMock2 *saler.ServiceMock2) *ControlMock2 {
	return &ControlMock2{
		serviceMock2: serviceMock2,
	}
}

func (controlMock2 ControlMock2) GetListGroup(ctx *gin.Context) {
	response, responerr := controlMock2.serviceMock2.ListBootcampGroup(ctx)
	if responerr != nil {
		ctx.JSON(responerr.Status, responerr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}
