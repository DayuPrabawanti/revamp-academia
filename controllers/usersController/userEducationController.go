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

type UserEducationController struct {
	UserEducationService *userServ.UserEducationService
}

// Declare constructor
func NewUserEducationController(UserEducationService *userServ.UserEducationService) *UserEducationController {
	return &UserEducationController{
		UserEducationService: UserEducationService,
	}
}

// GetList User Media
func (UserEducationController UserEducationController) GetListUsersEducation(ctx *gin.Context) {

	response, responseErr := UserEducationController.UserEducationService.GetListUsersEducation(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

	//ctx.JSON(http.StatusOK, "Hello gin Framework")
}

// Get User Media
func (UserEducationController UserEducationController) GetUserEducation(ctx *gin.Context) {

	educationId, err := strconv.Atoi(ctx.Param("id")) // atoi mengubah string ke integer

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := UserEducationController.UserEducationService.GetUserEducation(ctx, int32(educationId))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// Create User Media
func (UserEducationController UserEducationController) CreateUserEducation(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create media request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var media dbcontext.CreateEducationParams
	err = json.Unmarshal(body, &media)
	if err != nil {
		log.Println("Error while unmarshaling create media request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := UserEducationController.UserEducationService.CreateUserEducation(ctx, &media)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

// Update Table
func (UserEducationController UserEducationController) UpdateEducation(ctx *gin.Context) {

	educationId, err := strconv.Atoi(ctx.Param("id"))

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

	var user dbcontext.CreateEducationParams
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("Error while unmarshaling update email request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := UserEducationController.UserEducationService.UpdateEducation(ctx, &user, int64(educationId))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// Delete Table
func (UserEducationController UserEducationController) DeleteEducation(ctx *gin.Context) {

	educationId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseErr := UserEducationController.UserEducationService.DeleteEducation(ctx, int32(educationId))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}
