package jobhireController

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/jobhireRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

func (jh JobHireController) CreateJobPostController(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)

	if err != nil {
		log.Println("Error While Reading Create Address Request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var jobPost dbContext.CreateJobPostParams
	err = json.Unmarshal(body, &jobPost)

	if err != nil {
		log.Println("Error while unmarshalling create job post request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := jh.jobservice.CreateJobPostService(ctx, &jobPost)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (jh JobHireController) CreateJobPostWithDescription(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)

	if err != nil {
		log.Println("Error While Reading Create Address Request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var jobPost models.CreateJobPost
	err = json.Unmarshal(body, &jobPost)

	if err != nil {
		log.Println("Error while unmarshalling create job post request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := jh.jobservice.CreateJobPostWithDescription(ctx, &jobPost)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}
