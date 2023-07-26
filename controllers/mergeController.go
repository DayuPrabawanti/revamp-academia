package controllers

import (
	"net/http"

	"codeid.revampacademy/services"
	"github.com/gin-gonic/gin"
)

type JobPostingController struct {
	jobPostingService *services.JobPostingService
}

func NewJobPostingController(jobPostingService *services.JobPostingService) *JobPostingController {
	return &JobPostingController{
		jobPostingService: jobPostingService,
	}
}

func (jobPostingController JobPostingController) GetJobPostingHttp(ctx *gin.Context) {

	jopoTitle := ctx.Query("title")

	response, responseErr := jobPostingController.jobPostingService.GetJobPostingService(ctx, string(jopoTitle))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (jobPostingController JobPostingController) ListJobPostingHttp(ctx *gin.Context) {
	coba := ctx.Query("nama")
	response, responseErr := jobPostingController.jobPostingService.ListJobPostRepo(ctx, string(coba))

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
