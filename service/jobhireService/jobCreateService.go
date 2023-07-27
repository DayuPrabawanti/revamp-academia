package jobhireService

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/jobhireRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

func (js JobService) CreateJobPostService(ctx *gin.Context, jobParams *dbContext.CreateJobPostParams) (*models.CreateJobPost, *models.ResponseError) {
	responseErr := ValidateParamsJob(jobParams)
	if responseErr != nil {
		return nil, responseErr
	}
	return js.jobService.CreateJobPostRepo(ctx, jobParams)
}

func ValidateParamsJob(jobParams *dbContext.CreateJobPostParams) *models.ResponseError {
	if jobParams.JobHirePost.JopoEntityID == 0 {
		return &models.ResponseError{
			Message: "Invalid Job Id",
			Status:  http.StatusInternalServerError,
		}
	}
	if jobParams.JobHirePost.JopoNumber == "" {
		return &models.ResponseError{
			Message: "No Job Post Number Available",
		}
	}
	return nil
}
