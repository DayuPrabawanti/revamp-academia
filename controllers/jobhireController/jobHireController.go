package jobhireController

import (
	"net/http"

	// "codeid.revampacademy/service"
	"codeid.revampacademy/services/jobhireService"
	"github.com/gin-gonic/gin"
)

type JobHireController struct {
	jobservice *jobhireService.JobService
}

func NewJobControll(jobService *jobhireService.JobService) *JobHireController {
	return &JobHireController{
		jobservice: jobService,
	}
}

func (jh JobHireController) GetJobPostControl(ctx *gin.Context) {
	response, responseErr := jh.jobservice.GetListJobPost(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
	}
	ctx.JSON(http.StatusOK, response)
}

func (jh JobHireController) GetJobPostMergeControl(ctx *gin.Context) {
	response, responseErr := jh.jobservice.GetListJobMerge(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
	}
	ctx.JSON(http.StatusOK, response)
}
