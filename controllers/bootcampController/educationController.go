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

type EducationController struct {
	educationService *bootcampService.UsersUsersEducation
}

// Declare constructor
func NewEducationController(educationService *bootcampService.UsersUsersEducation) *EducationController {
	return &EducationController{
		educationService: educationService,
	}
}

func (EducationController EducationController) GetListEducation(ctx *gin.Context) {

	response, responseErr := EducationController.educationService.GetListEducation(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

	//ctx.JSON(http.StatusOK, "Hello gin Framework")
}

func (EducationController EducationController) UpdateEducation(ctx *gin.Context) {

	UsduEntityID, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update Education request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var education dbContext.UsersUsersEducation
	err = json.Unmarshal(body, &education)
	if err != nil {
		log.Println("Error while unmarshaling update Education request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := EducationController.educationService.UpdateEducation(ctx, &education, int64(UsduEntityID))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (EducationController EducationController) GetEducation(ctx *gin.Context) {

	usduId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := EducationController.educationService.GetEducation(ctx, int64(usduId))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
