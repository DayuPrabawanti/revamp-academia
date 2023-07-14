package controllers

import (
	"net/http"

	"codeid.revampacademy/services"
	"github.com/gin-gonic/gin"
)

type SectionDetailMaterialController struct {
	sectionDetailMaterialService *services.SectionDetailMaterialService
}

func NewSectionDetailMaterialController(sectionDetailMaterialService *services.SectionDetailMaterialService) *SectionDetailMaterialController {
	return &SectionDetailMaterialController{
		sectionDetailMaterialService: sectionDetailMaterialService,
	}
}

func (sectionDetailMaterialController SectionDetailMaterialController) GetListSectionDetailMaterial(ctx *gin.Context) {
	response, responseErr := sectionDetailMaterialController.sectionDetailMaterialService.GetListSectionDetailMaterial(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
