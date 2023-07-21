package salesController

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"codeid.revampacademy/repositories/dbContext/salesContext"
	"codeid.revampacademy/services/salesServices"
	"github.com/gin-gonic/gin"
)

type EducationController struct {
	educationService *salesServices.EducationService
}

// declare Constructor
func NewEducationController(educationService *salesServices.EducationService) *EducationController {
	return &EducationController{
		educationService: educationService,
	}
}

func (educationController EducationController) CreateEducation(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var education salesContext.CreateEducationParams
	err = json.Unmarshal(body, &education)
	if err != nil {
		log.Println("Error while unmarshaling create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := educationController.educationService.CreateEducation(ctx, &education)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (educationController EducationController) GetEducation(ctx *gin.Context) {

	UsduID, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := educationController.educationService.GetEducation(ctx, int32(UsduID))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
