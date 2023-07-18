package controllers

import (
	"net/http"

	"codeid.revampacademy/services"
	"github.com/gin-gonic/gin"
)

type TalentsDetailMockupController struct {
	talentDetailService *services.TalentsDetailMockupService
}

// declare constructor
func NewTalentDetailMockupController(talentDetailService *services.TalentsDetailMockupService) *TalentsDetailMockupController {
	return &TalentsDetailMockupController{
		// struct 				parameter
		talentDetailService: talentDetailService,
	}
}

func (talentDetailController TalentsDetailMockupController) GetListTalentDetailMockuptDetail(ctx *gin.Context) {
	responses, responseErr := talentDetailController.talentDetailService.GetListTalentDetailMockup(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, responses)

	// ctx.JSON(http.StatusOK, "Hello gin framework")
}
