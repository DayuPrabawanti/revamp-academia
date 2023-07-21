package jobhireController

import (
	"net/http"

	// "codeid.revampacademy/service"
	"codeid.revampacademy/services/jobhireService"
	"github.com/gin-gonic/gin"
)

type MasterController struct {
	masterService *jobhireService.MasterService
}

func NewMasterController(masterService *jobhireService.MasterService) *MasterController {
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
