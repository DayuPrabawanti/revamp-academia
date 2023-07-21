package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

)

func (jobPostingController JobPostController) ListJobPostingHttp(ctx *gin.Context) {
	response, responseErr := jobPostingController.JobPostService.ListJobPostingRepo(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}