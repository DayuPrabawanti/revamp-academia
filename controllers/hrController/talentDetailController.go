package hrController

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"codeid.revampacademy/repositories/hrRepository/dbContext"
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

func (talentDetailController TalentsDetailMockupController) SearchTalentDetail(ctx *gin.Context) {
	clitName := ctx.DefaultQuery("name", "")

	talents, responseErr := talentDetailController.talentDetailService.SearchTalentDetail(ctx, clitName)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, talents)
}

func (talentDetailController TalentsDetailMockupController) UpdateSwitch(ctx *gin.Context) {

	id := ctx.Query("id") // Mengambil nilai query parameter id dari URL

	switchUp, err := strconv.Atoi(id)

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update department request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var switchIdle dbContext.UpdateSwitchParams
	err = json.Unmarshal(body, &switchIdle)
	if err != nil {
		log.Println("Error while unmarshaling update department request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := talentDetailController.talentDetailService.UpdateSwitch(ctx, &switchIdle, int64(switchUp))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}
