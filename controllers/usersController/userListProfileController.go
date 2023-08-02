package usersController

import (
	"log"
	"net/http"
	"strconv"

	"codeid.revampacademy/services/usersService"
	"github.com/gin-gonic/gin"
)

type UserListProfileController struct {
	userListProfileService *usersService.UserListProfileService
}

// Declare constructor
func NewUserListProfileController(userListProfileService *usersService.UserListProfileService) *UserListProfileController {
	return &UserListProfileController{
		userListProfileService: userListProfileService,
	}
}

func (userListProfileController UserListProfileController) GetProfileUser(ctx *gin.Context) {

	userId, err := strconv.Atoi(ctx.Param("id")) // atoi mengubah string ke integer

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := userListProfileController.userListProfileService.GetProfileUsers(ctx, int32(userId))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (userListProfileController UserListProfileController) GetProfileEmail(ctx *gin.Context) {

	userId, err := strconv.Atoi(ctx.Param("id")) // atoi mengubah string ke integer

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := userListProfileController.userListProfileService.GetProfileEmail(ctx, int32(userId))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (userListProfileController UserListProfileController) GetProfile(ctx *gin.Context) {

	userId, err := strconv.Atoi(ctx.Param("id")) // atoi mengubah string ke integer

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := userListProfileController.userListProfileService.GetProfile(ctx, int32(userId))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}