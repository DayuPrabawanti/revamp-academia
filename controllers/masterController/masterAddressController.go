package masterController

import (
	"net/http"

	"codeid.revampacademy/services/masterService"
	"github.com/gin-gonic/gin"
)

type MasterAddressController struct {
	masterAddressService *masterService.MasterAddressService
}

// Declare constructor
func NewMasterAddressController(masterAddressService *masterService.MasterAddressService) *MasterAddressController {
	return &MasterAddressController{
		masterAddressService: masterAddressService,
	}
}

func (masterAddressController MasterAddressController) GetListMasterAddress(ctx *gin.Context){

	response, responseErr := masterAddressController.masterAddressService.GetListMasterAddress(ctx)

	if responseErr != nil{
		ctx.JSON(responseErr.Status,responseErr)
		return 
	}

	ctx.JSON(http.StatusOK,response)
	
}