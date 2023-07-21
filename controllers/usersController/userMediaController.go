package usersController

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	dbcontext "codeid.revampacademy/repositories/usersRepository/dbContext"
	userServ "codeid.revampacademy/services/usersService"
	"github.com/gin-gonic/gin"
)

type UserMediaController struct {
	userMediaService *userServ.UserMediaService
}

// Declare constructor
func NewUserMediaController(userMediaService *userServ.UserMediaService) *UserMediaController {
	return &UserMediaController{
		userMediaService: userMediaService,
	}
}

// GetList User Media
func (userMediaController UserMediaController) GetListUserMedia(ctx *gin.Context) {

	response, responseErr := userMediaController.userMediaService.GetListUserMedia(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

	//ctx.JSON(http.StatusOK, "Hello gin Framework")
}

// Get User Media
func (userMediaController UserMediaController) GetUserMedia(ctx *gin.Context) {

	mediaId, err := strconv.Atoi(ctx.Param("id")) // atoi mengubah string ke integer

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := userMediaController.userMediaService.GetUserMedia(ctx, int32(mediaId))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// Create User Media
func (usermediacontroller UserMediaController) CreateUserMedia(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create media request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var media dbcontext.CreateMediaParams
	err = json.Unmarshal(body, &media)
	if err != nil {
		log.Println("Error while unmarshaling create media request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := usermediacontroller.userMediaService.CreateUserMedia(ctx, &media)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

// Update Table
func (userMediaController UserMediaController) UpdateMedia(ctx *gin.Context) {

	mediaId, err := strconv.Atoi(ctx.Param("id"))

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

	var user dbcontext.CreateMediaParams
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("Error while unmarshaling update email request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := userMediaController.userMediaService.UpdateMedia(ctx, &user, int64(mediaId))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// Delete Table
func (userMediaController UserMediaController) DeleteMedia(ctx *gin.Context) {

	mediaId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseErr := userMediaController.userMediaService.DeleteMedia(ctx, int32(mediaId))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}
