package hrController

import (
	"log"
	"net/http"
	"strconv"

	"codeid.revampacademy/services/hrService"
	"github.com/gin-gonic/gin"
)

type TalentsDetailMockupController struct {
	talentDetailService *hrService.TalentsDetailMockupService
}

// declare constructor
func NewTalentDetailMockupController(talentDetailService *hrService.TalentsDetailMockupService) *TalentsDetailMockupController {
	return &TalentsDetailMockupController{
		// struct 				parameter
		talentDetailService: talentDetailService,
	}
}

func (talentDetailController TalentsDetailMockupController) GetListTalentDetailMockupDetail(ctx *gin.Context) {
	responses, responseErr := talentDetailController.talentDetailService.GetListTalentDetailMockup(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, responses)

	// ctx.JSON(http.StatusOK, "Hello gin framework")
}

func (talentDetailController TalentsDetailMockupController) GetTalentDetail(ctx *gin.Context) {

	user_entity_id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := talentDetailController.talentDetailService.GetTalentDetail(ctx, int64(user_entity_id))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
