package salesController

import (
	"log"
	"net/http"
	"strconv"

	"codeid.revampacademy/services/salesServices"
	"github.com/gin-gonic/gin"
)

type FintechController struct {
	fintechService *salesServices.FintechService
}

func NewFintechController(fintechService *salesServices.FintechService) *FintechController {
	return &FintechController{
		fintechService: fintechService,
	}
}

func (fintechController FintechController) GetPaymentFintech(ctx *gin.Context) {

	FintEntityID, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := fintechController.fintechService.GetPaymentFintech(ctx, int32(FintEntityID))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
