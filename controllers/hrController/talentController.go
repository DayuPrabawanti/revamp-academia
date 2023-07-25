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
	skillName := ctx.DefaultQuery("tekno", "")
	batchStatus := ctx.DefaultQuery("status", "")

	talents, responseErr := talentController.talentService.SearchTalent(ctx, userName, skillName, batchStatus)
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

func (talentController TalentsMockupController) GetBatch(ctx *gin.Context) {
	id := ctx.Query("batchid") // Mengambil nilai query parameter id dari URL

	batchId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Error while parsing id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := talentController.talentService.GetBatch(ctx, int64(batchId))
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (talentController TalentsMockupController) UpdateBatch(ctx *gin.Context) {
	id := ctx.Query("id") // Mengambil nilai query parameter id dari URL

	batchId, err := strconv.Atoi(id)
	// batchId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update batch request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var batch dbContext.UpdateBatchParams
	err = json.Unmarshal(body, &batch)
	if err != nil {
		log.Println("Error while unmarshaling update batch request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := talentController.talentService.UpdateBatch(ctx, &batch, int64(batchId))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}
