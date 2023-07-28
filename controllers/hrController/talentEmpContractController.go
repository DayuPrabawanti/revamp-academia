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

type TalentClientContractController struct {
	talentClientContractService *hrService.TalentClientContractService
}

// declare constructor
func NewTalentClientContractController(talentClientContractService *hrService.TalentClientContractService) *TalentClientContractController {
	return &TalentClientContractController{
		talentClientContractService: talentClientContractService,
	}
}

func (talentClientContractController TalentClientContractController) GetTalentClientContract(ctx *gin.Context) {

	ecco_id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := talentClientContractController.talentClientContractService.GetTalentClientContract(ctx, int64(ecco_id))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (talentClientContractController TalentClientContractController) UpdateTalentClientContract(ctx *gin.Context) {

	id := ctx.Query("id") // Mengambil nilai query parameter id dari URL

	ecco_id, err := strconv.Atoi(id)

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update client contract request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var clientContract dbContext.UpdateTalentClientContractParams
	err = json.Unmarshal(body, &clientContract)
	if err != nil {
		log.Println("Error while unmarshaling update client contract request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := talentClientContractController.talentClientContractService.UpdateTalentClientContract(ctx, &clientContract, int64(ecco_id))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}
