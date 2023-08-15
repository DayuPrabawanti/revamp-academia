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
func (userPhoneController UserPhoneController) UpdatePhone(ctx *gin.Context) {

	phoneId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update Phone Number request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var phone dbContext.CreatePhonesParams
	err = json.Unmarshal(body, &phone)
	if err != nil {
		log.Println("Error while unmarshaling update phone number request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := userPhoneController.userPhoneService.UpdatePhone(ctx, &phone, int64(phoneId))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (UserPhoneController UserPhoneController) GetUserPhone(ctx *gin.Context) {

	UspoEntityID, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := UserPhoneController.userPhoneService.GetUserPhone(ctx, int64(UspoEntityID))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
