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

type EditEmailController struct {
	userEmailService *usersService.EditEmailService
}

// Declare constructor
func NewEditEmailController(userEmailService *usersService.EditEmailService) *EditEmailController {
	return &EditEmailController{
		userEmailService: userEmailService,
	}
}

// func (userEmailcontroller EditEmailController) AddEmail(ctx *gin.Context) {

// 	body, err := io.ReadAll(ctx.Request.Body)
// 	if err != nil {
// 		log.Println("Error while reading create category request body", err)
// 		ctx.AbortWithError(http.StatusInternalServerError, err)
// 		return
// 	}

// 	var email dbContext.AddEmailParams
// 	err = json.Unmarshal(body, &email)
// 	if err != nil {
// 		log.Println("Error while unmarshaling create User request body", err)
// 		ctx.AbortWithError(http.StatusInternalServerError, err)
// 		return
// 	}

// 	response, responseErr := userEmailcontroller.userEmailService.AddEmail(ctx, &email)
// 	if responseErr != nil {
// 		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, response)

// }

func (addEmailController EditEmailController) AddEmail(ctx *gin.Context) {

	emailId, err := strconv.Atoi(ctx.Param("id"))

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
	var emailParams *dbContext.AddEmailParams
	err = json.Unmarshal(body, &emailParams)
	if err != nil {
		log.Println("Error while unmarshaling create User request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	responseErr := addEmailController.userEmailService.AddEmail(ctx, emailParams, int32(emailId))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}