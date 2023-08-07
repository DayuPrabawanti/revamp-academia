package usersController

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"codeid.revampacademy/repositories/usersRepository/dbContext"
	"codeid.revampacademy/services/usersService"
	"github.com/gin-gonic/gin"
)

type EditSkillController struct {
	userSkillService *usersService.EditSkillService
}

// Declare constructor
func NewEditSkillController(EditSkillService *usersService.EditSkillService) *EditSkillController {
	return &EditSkillController{
		userSkillService: EditSkillService,
	}
}

func (addSkillController EditSkillController) AddSkill(ctx *gin.Context) {

	skillId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading skill request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var skillParams *dbContext.AddSkillParams
	err = json.Unmarshal(body, &skillParams)
	if err != nil {
		log.Println("Error while unmarshaling skill request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	responseErr := addSkillController.userSkillService.AddSkill(ctx, skillParams, int32(skillId))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}