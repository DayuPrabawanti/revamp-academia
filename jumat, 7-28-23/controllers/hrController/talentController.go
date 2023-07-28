package hrController

import (
	"net/http"
	"strconv"

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

func (talentController TalentsMockupController) SearchTalent(ctx *gin.Context) {
	userName := ctx.DefaultQuery("name", "")
	userSkill := ctx.DefaultQuery("skill", "")
	batchName := ctx.DefaultQuery("batch", "")
	status := ctx.DefaultQuery("status", "")

	talents, responseErr := talentController.talentService.SearchTalent(ctx, userName, userSkill, batchName, status)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, talents)
}

func (talentController TalentsMockupController) PagingTalent(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	pageSize, err := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	if err != nil || pageSize <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size"})
		return
	}

	offset := (page - 1) * pageSize

	talents, responseErr := talentController.talentService.PagingTalent(ctx, offset, pageSize)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, talents)
}
