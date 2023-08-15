package bootcampController

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"codeid.revampacademy/repositories/bootcampRepository/dbContext"
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

func (userEmailController UserEmailController) UpdateEmail(ctx *gin.Context) {

	emailId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update email request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var user dbContext.CreateEmailParams
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("Error while unmarshaling update email request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := userEmailController.userEmailService.UpdateEmail(ctx, &user, int64(emailId))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (UserEmailController UserEmailController) GetUSerEmail(ctx *gin.Context) {

	pmailEntityId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := UserEmailController.userEmailService.GetUSerEmail(ctx, int64(pmailEntityId))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
