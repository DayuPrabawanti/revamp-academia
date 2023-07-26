package controllers

import (
	"log"
	"net/http"
	"strconv"

	"codeid.revampacademy/services"
	"github.com/gin-gonic/gin"
)

type MasterIndustryController struct {
	MasterIndustryService *services.MasterIndustryService
}

func NewMasterIndustryController(masterIndustryController *services.MasterIndustryService) *MasterIndustryController {
	return &MasterIndustryController{
		MasterIndustryService: masterIndustryController,
	}
}


func (masterIndustryController MasterIndustryController) GetMasterIndustryHttp(ctx *gin.Context) {

				InduCodeID, err := strconv.Atoi(ctx.Param("id"))

				if err != nil {
					log.Println("Error while reading paramater id", err)
					ctx.AbortWithError(http.StatusBadRequest, err)
					return
				}

				response, responseErr := masterIndustryController.MasterIndustryService.GetMasterIndustryService(ctx, int32(InduCodeID))
				if responseErr != nil {

					ctx.JSON(responseErr.Status, responseErr)
					return
				}

				ctx.JSON(http.StatusOK, response)
}

	func (masterIndustryController MasterIndustryController) ListMasterIndustryHttp(ctx *gin.Context) {
						response, responseErr := masterIndustryController.MasterIndustryService.ListMasterIndustryService(ctx)

						if responseErr != nil {
							ctx.JSON(responseErr.Status, responseErr)
							return
						}

						ctx.JSON(http.StatusOK, response)
	}