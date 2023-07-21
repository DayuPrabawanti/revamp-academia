package hrController

import (
	"net/http"

	"codeid.revampacademy/services/hrService"
	"github.com/gin-gonic/gin"
)

type TalentsMockupController struct {
	talentService *hrService.TalentsMockupService
}

// declare constructor
func NewTalentMockupController(talentService *hrService.TalentsMockupService) *TalentsMockupController {
	return &TalentsMockupController{
		// struct 				parameter
		talentService: talentService,
	}
}

func (talentController TalentsMockupController) GetListTalentMockup(ctx *gin.Context) {
	responses, responseErr := talentController.talentService.GetListTalentMockup(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, responses)

	// ctx.JSON(http.StatusOK, "Hello gin framework")
}
