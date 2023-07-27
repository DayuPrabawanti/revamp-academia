package jobhireController

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	// "codeid.revampacademy/service"
	"codeid.revampacademy/repositories/jobhireRepositories/dbContext"
	"codeid.revampacademy/service/jobhireService"
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

func (mc MasterController) CreateAddress(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error While reading create address request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var address dbContext.CreateMasterAddressParams
	err = json.Unmarshal(body, &address)

	if err != nil {
		log.Println("Error while unmarshalling create address request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := mc.masterService.CreateMasterAddressService(ctx, &address)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}
