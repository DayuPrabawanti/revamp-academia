package curriculumcontrollers

import (
	"log"
	"net/http"
	"strconv"

	services "codeid.revampacademy/services/curriculumServices"
	"github.com/gin-gonic/gin"
)

type SectionDetailController struct {
	sectionDetailService *services.SectionDetailService
}

func NewSectionDetailController(sectionDetailService *services.SectionDetailService) *SectionDetailController {
	return &SectionDetailController{
		sectionDetailService: sectionDetailService,
	}
}

func (sectionDetailController SectionDetailController) GetSectionDetail(ctx *gin.Context) {

	sedmId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := sectionDetailController.sectionDetailService.GetSectionDetail(ctx, int64(sedmId))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
