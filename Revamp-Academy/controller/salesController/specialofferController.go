package salesController

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"codeid.revampacademy/repositories/salesRepositories/dbContext"
	"codeid.revampacademy/services/salesServices"
	"github.com/gin-gonic/gin"
)

type SpecialOfferController struct {
	specialOfferService *salesServices.SpecialOfferService
}

// declare Constructor
func NewSpecialController(specialOfferService *salesServices.SpecialOfferService) *SpecialOfferController {
	return &SpecialOfferController{
		specialOfferService: specialOfferService,
	}
}

// create method
func (specialOfferController SpecialOfferController) GetListSpecialOffer(ctx *gin.Context) {
	response, resposeErr := specialOfferController.specialOfferService.GetListSpecialOffer(ctx)

	if resposeErr != nil {
		ctx.JSON(resposeErr.Status, resposeErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

	// ctx.JSON(http.StatusOK, "Hello gin Framework")
}

func (specialOfferController SpecialOfferController) GetSpecial_offer(ctx *gin.Context) {

	spof_id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := specialOfferController.specialOfferService.GetSpecialOffer(ctx, int64(spof_id))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (specialOfferController SpecialOfferController) CreateSpecialOffer(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var specialoffer dbContext.CreateSales_special_offerParams
	err = json.Unmarshal(body, &specialoffer)
	if err != nil {
		log.Println("Error while unmarshaling create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := specialOfferController.specialOfferService.CreateSales_special_offer(ctx, &specialoffer)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}
