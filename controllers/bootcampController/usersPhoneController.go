package bootcampController

import (
	"net/http"

	"codeid.revampacademy/services/bootcampService"
	"github.com/gin-gonic/gin"
)

type UserPhoneController struct {
	userPhoneService *bootcampService.UserPhoneService
}

// Declare constructor
func NewUserPhoneController(userPhoneService *bootcampService.UserPhoneService) *UserPhoneController {
	return &UserPhoneController{
		userPhoneService: userPhoneService,
	}
}

func (userPhoneController UserPhoneController) GetListUsersPhone(ctx *gin.Context) {

	response, responseErr := userPhoneController.userPhoneService.GetListUsersPhone(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

	//ctx.JSON(http.StatusOK, "Hello gin Framework")
}
