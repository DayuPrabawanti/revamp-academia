package jobhireService

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/jobhireRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

func (jp JobService) UpdateJobPost(ctx *gin.Context, jobPostParams *dbContext.UpdateJobPostParams, id int64) *models.ResponseError {

	responseErr := ValidateParamsJobForUpdate(jobPostParams)

	if responseErr != nil {
		return responseErr
	}
	return jp.repositoryMgr.UpdateJobPosting(ctx, jobPostParams)
}

func ValidateParamsJobForUpdate(jobParams *dbContext.UpdateJobPostParams) *models.ResponseError {
	if jobParams.JobHirePost.JopoEntityID == 0 {
		return &models.ResponseError{
			Message: "Invalid job post id",
			Status:  http.StatusBadRequest,
		}
	}
	if jobParams.JobHirePost.JopoTitle == "" {
		return &models.ResponseError{
			Message: "Invalid Job Title",
			Status:  http.StatusBadRequest,
		}
	}
	return nil
}
