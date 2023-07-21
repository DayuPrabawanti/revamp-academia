package bootcampController

import (
	"net/http"

	"codeid.revampacademy/services/bootcampService"
	"github.com/gin-gonic/gin"
)

type UserEmailController struct {
	userEmailService *bootcampService.UserEmailService
}

// Declare constructor
func NewUserEmailController(userEmailService *bootcampService.UserEmailService) *UserEmailController {
	return &UserEmailController{
		userEmailService: userEmailService,
	}
}

func (userEmailController UserEmailController) GetListUsersEmail(ctx *gin.Context) {

	response, responseErr := userEmailController.userEmailService.GetListUsersEmail(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

	//ctx.JSON(http.StatusOK, "Hello gin Framework")
}
