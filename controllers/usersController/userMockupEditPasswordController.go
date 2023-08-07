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

type EditUserPasswordController struct {
	userPasswordService *usersService.EditPasswordService
}

// Declare constructor
func NewEditUserPasswordController(userpasswordService *usersService.EditPasswordService) *EditUserPasswordController {
	return &EditUserPasswordController{
		userPasswordService: userpasswordService,
	}
}

func (EditUserpasswordController EditUserPasswordController) GetPassword(ctx *gin.Context) {

	userId, err := strconv.Atoi(ctx.Param("id")) // atoi mengubah string ke integer

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := EditUserpasswordController.userPasswordService.GetUserPassword(ctx, int32(userId))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (userPasswordController EditUserPasswordController) UpdatePassword(ctx *gin.Context) {

	userpasswordId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var userpassword dbContext.EditUsersPasswordParams
	err = json.Unmarshal(body, &userpassword)
	if err != nil {
		log.Println("Error while unmarshaling update category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := userPasswordController.userPasswordService.EditUserPassword(ctx, int32(userpasswordId),&userpassword )
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}