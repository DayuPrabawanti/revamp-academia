package controller

import (
	"net/http"

	"codeid.revampacademy/service"
	"github.com/gin-gonic/gin"
)

type MasterController struct {
	masterService *service.MasterService
}

func NewMasterController(masterService *service.MasterService) *MasterController {
	return &MasterController{
		masterService: masterService,
	}
}

func (mc MasterController) GetListAddressControl(ctx *gin.Context) {
	response, responseErr := mc.masterService.GetListMasterAddress(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
	}
	ctx.JSON(http.StatusOK, response)
}

func (mc MasterController) GetListCityControl(ctx *gin.Context) {
	response, responseErr := mc.masterService.GetListMasterCity(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
	}
	ctx.JSON(http.StatusOK, response)
}
