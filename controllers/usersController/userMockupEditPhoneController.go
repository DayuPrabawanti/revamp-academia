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

type EditPhoneController struct {
	userPhoneService *usersService.EditPhoneService
}

// Declare constructor
func NewEditPhoneController(userPhoneService *usersService.EditPhoneService) *EditPhoneController {
	return &EditPhoneController{
		userPhoneService: userPhoneService,
	}
}


func (addPhoneController EditPhoneController) AddPhone(ctx *gin.Context) {

	phoneId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var phoneParams *dbContext.AddPhonesParams
	err = json.Unmarshal(body, &phoneParams)
	if err != nil {
		log.Println("Error while unmarshaling create User request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	responseErr := addPhoneController.userPhoneService.AddPhone(ctx, phoneParams, int32(phoneId))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}