package jobhireController

import (
	"log"
	"net/http"
	"strconv"

	// "codeid.revampacademy/service"
	"codeid.revampacademy/service/jobhireService"
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

func (jh JobHireController) GetJobPostDetailControl(ctx *gin.Context) {
	jobPostId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading parameter id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := jh.jobservice.GetJobDetailService(ctx, int32(jobPostId))
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (jh JobHireController) GetJobPostSearch(ctx *gin.Context) {
	cityName := ctx.Query("cityname")
	joroName := ctx.Query("joroName")
	wotyName := ctx.Query("wotyName")

	response, responseErr := jh.jobservice.GetListJobPostSearch(ctx, cityName, joroName, wotyName)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}
