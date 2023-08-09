package salescontrollers

import (
	"net/http"

	saler "codeid.revampacademy/service/salesService"
	"github.com/gin-gonic/gin"
)

type ControlMockup8 struct {
	serviceMockup8 *saler.ServiceMockup8
}

func NewControlShoppingCart3(serviceMockup8 *saler.ServiceMockup8) *ControlMockup8 {
	return &ControlMockup8{
		serviceMockup8: serviceMockup8,
	}
}

func (controlMockup8 ControlMockup8) GetIdSummaryOrderMock8Control(ctx *gin.Context) {
	account := ctx.Query("number")
	response, responseErr := controlMockup8.serviceMockup8.GetIdSummaryOrderMock8Service(ctx, string(account))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
