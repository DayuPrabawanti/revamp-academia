package usersController

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"codeid.revampacademy/repositories/usersRepository/dbContext"
	"codeid.revampacademy/services/usersService"
	"github.com/gin-gonic/gin"
)

type UserExperienceController struct {
	userExperienceService *usersService.UserExperienceService
}

// Declare constructor
func NewUserExperienceController(userExperienceService *usersService.UserExperienceService) *UserExperienceController {
	return &UserExperienceController{
		userExperienceService: userExperienceService,
	}
}

func (userExperienceController UserExperienceController) GetListUserExperience(ctx *gin.Context){

	response, responseErr := userExperienceController.userExperienceService.GetListUserExperience(ctx)

	if responseErr != nil{
		ctx.JSON(responseErr.Status,responseErr)
		return 
	}

	ctx.JSON(http.StatusOK,response)
	
	//ctx.JSON(http.StatusOK, "Hello gin Framework")
}

func (userExperienceController UserExperienceController) GetExperience(ctx *gin.Context) {

	userEx, err := strconv.Atoi(ctx.Param("id")) // atoi mengubah string ke integer

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := userExperienceController.userExperienceService.GetExperience(ctx, int32(userEx))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}


func (userExperienceController UserExperienceController) CreateExperience(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create experience request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var userEx dbContext.CreateExperienceParams
	err = json.Unmarshal(body, &userEx)
	if err != nil {
		log.Println("Error while unmarshaling create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := userExperienceController.userExperienceService.CreateExperience(ctx, &userEx)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (userExperienceController UserExperienceController) UpdateExperience(ctx *gin.Context) {

	userEx, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update experience request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var experience dbContext.CreateExperienceParams
	err = json.Unmarshal(body, &experience)
	if err != nil {
		log.Println("Error while unmarshaling update experience request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := userExperienceController.userExperienceService.UpdateUserExperience(ctx, &experience, int64(userEx))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (userExperienceController UserExperienceController) DeleteExperience(ctx *gin.Context) {

	userEx, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseErr := userExperienceController.userExperienceService.DeleteExperience(ctx, int32(userEx))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}