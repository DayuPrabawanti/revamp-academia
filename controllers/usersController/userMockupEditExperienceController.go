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

type EditExperienceController struct {
	userExperienceService *usersService.EditExperienceService
}

// Declare constructor
func NewEditExperienceController(userExperienceService *usersService.EditExperienceService) *EditExperienceController {
	return &EditExperienceController{
		userExperienceService: userExperienceService,
	}
}

func (addExperienceController EditExperienceController) AddExperience(ctx *gin.Context) {

	experienceId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create experience request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var experienceParams *dbContext.AddExperienceParams
	err = json.Unmarshal(body, &experienceParams)
	if err != nil {
		log.Println("Error while unmarshaling insert User request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	responseErr := addExperienceController.userExperienceService.AddExperience(ctx, experienceParams, int32(experienceId))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}