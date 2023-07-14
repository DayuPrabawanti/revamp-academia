package controllers

import (
	"net/http"

	"codeid.revampacademy/services"
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

func (sectionDetailController SectionDetailController) GetListSectionDetail(ctx *gin.Context) {
	response, responseErr := sectionDetailController.sectionDetailService.GetListSectionDetail(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
