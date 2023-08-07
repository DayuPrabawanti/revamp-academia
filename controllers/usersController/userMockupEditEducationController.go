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

type EditEducationController struct {
	userEducationService *usersService.EditEducationService
}

// Declare constructor
func NewEditEducationController(userEducationService *usersService.EditEducationService) *EditEducationController {
	return &EditEducationController{
		userEducationService: userEducationService,
	}
}

func (addEducationController EditEducationController) AddEducation(ctx *gin.Context) {

	educationId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create education request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var educationParams *dbContext.AddEducationParams
	err = json.Unmarshal(body, &educationParams)
	if err != nil {
		log.Println("Error while unmarshaling insert User request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	responseErr := addEducationController.userEducationService.AddEducation(ctx, educationParams, int32(educationId))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}