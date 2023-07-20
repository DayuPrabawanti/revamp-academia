package controllers

import (
	"net/http"

	"codeid.revampacademy/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

// Declare constructor
func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (userController UserController) GetListUser(ctx *gin.Context) {

	response, responseErr := userController.userService.GetListUser(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

	//ctx.JSON(http.StatusOK, "Hello gin Framework")
}
