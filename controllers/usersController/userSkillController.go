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

type UserSkillController struct {
	UserSkillService *userServ.UserSkillService
}

// Declare constructor
func NewUserSkillController(UserSkillService *userServ.UserSkillService) *UserSkillController {
	return &UserSkillController{
		UserSkillService: UserSkillService,
	}
}

// GetList User License
func (UserSkillController UserSkillController) GetListUserSkill(ctx *gin.Context) {

	response, responseErr := UserSkillController.UserSkillService.GetListUserSkill(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

	//ctx.JSON(http.StatusOK, "Hello gin Framework")
}

// Get User License
func (UserSkillController UserSkillController) GetUsersSkill(ctx *gin.Context) {

	skillId, err := strconv.Atoi(ctx.Param("id")) // atoi mengubah string ke integer

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := UserSkillController.UserSkillService.GetUserSkill(ctx, int32(skillId))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// Create User License
func (UserSkillController UserSkillController) CreateUserSkill(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create experience request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var userSkill dbContext.CreateSkillParams
	err = json.Unmarshal(body, &userSkill)
	if err != nil {
		log.Println("Error while unmarshaling create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := UserSkillController.UserSkillService.CreateUserSkill(ctx, &userSkill)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

// Update User License
func (UserSkillController UserSkillController) UpdateUserSkill(ctx *gin.Context) {

	skillId, err := strconv.Atoi(ctx.Param("id"))

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

	var userSkill dbContext.CreateSkillParams
	err = json.Unmarshal(body, &userSkill)
	if err != nil {
		log.Println("Error while unmarshaling update email request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := UserSkillController.UserSkillService.UpdateSkill(ctx, &userSkill, int64(skillId))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// Delete Table
func (UserSkillController UserSkillController) DeleteSkill(ctx *gin.Context) {

	skillId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseErr := UserSkillController.UserSkillService.DeleteSkill(ctx, int32(skillId))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}
