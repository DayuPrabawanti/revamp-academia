package controller

import (
	"net/http"

	"codeid.revampacademy/service"
	"github.com/gin-gonic/gin"
)

type JobHire struct {
	jobservice *service.JobService
}

func NewJobControll(jobService *service.JobService) *JobHire {
	return &JobHire{
		jobservice: jobService,
	}
}

func (jh JobHire) GetJobPostControl(ctx *gin.Context) {
	response, responseErr := jh.jobservice.GetListJobPost(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
	}
	ctx.JSON(http.StatusOK, response)
}

func (jh JobHire) GetJobPostMergeControl(ctx *gin.Context) {
	response, responseErr := jh.jobservice.GetListJobMerge(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
	}
	ctx.JSON(http.StatusOK, response)
}
