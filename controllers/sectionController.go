package controllers

import (
	"net/http"

	"codeid.revampacademy/services"
	"github.com/gin-gonic/gin"
)

type SectionController struct {
	sectionService *services.SectionService
}

func NewSectionController(sectionService *services.SectionService) *SectionController {
	return &SectionController{
		sectionService: sectionService,
	}
}

func (sectionController SectionController) GetListSection(ctx *gin.Context) {
	response, responseErr := sectionController.sectionService.GetListSection(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
