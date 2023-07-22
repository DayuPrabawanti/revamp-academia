package usersController

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"codeid.revampacademy/repositories/usersRepository/dbContext"
	userServ "codeid.revampacademy/services/usersService"
	"github.com/gin-gonic/gin"
)

type UserLicenseController struct {
	UserLicenseService *userServ.UserLicenseService
}

// Declare constructor
func NewUserLicenseController(UserLicenseService *userServ.UserLicenseService) *UserLicenseController {
	return &UserLicenseController{
		UserLicenseService: UserLicenseService,
	}
}

// GetList User License
func (UserLicenseController UserLicenseController) GetListUserLicense(ctx *gin.Context) {

	response, responseErr := UserLicenseController.UserLicenseService.GetListUserLicense(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

	//ctx.JSON(http.StatusOK, "Hello gin Framework")
}

// Get User License
func (UserLicenseController UserLicenseController) GetUsersLicense(ctx *gin.Context) {

	licenseId, err := strconv.Atoi(ctx.Param("id")) // atoi mengubah string ke integer

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := UserLicenseController.UserLicenseService.GetuserLicense(ctx, int32(licenseId))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// Create User License
func (UserLicenseController UserLicenseController) CreateUserLicense(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create experience request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var userLicense dbContext.CreateLicenseParams
	err = json.Unmarshal(body, &userLicense)
	if err != nil {
		log.Println("Error while unmarshaling create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := UserLicenseController.UserLicenseService.CreateUserLicense(ctx, &userLicense)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

// Update User License
func (UserLicenseController UserLicenseController) UpdateUserLicense(ctx *gin.Context) {

	licenseId, err := strconv.Atoi(ctx.Param("id"))

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

	var user dbContext.CreateLicenseParams
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("Error while unmarshaling update email request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := UserLicenseController.UserLicenseService.UpdateLicense(ctx, &user, int64(licenseId))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// Delete Table
func (UserLicenseController UserLicenseController) DeleteLicense(ctx *gin.Context) {

	licenseId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseErr := UserLicenseController.UserLicenseService.DeleteLicense(ctx, int32(licenseId))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}
