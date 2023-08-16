package jobhireController

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"codeid.revampacademy/repositories/jobhireRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

func (jc JobHireController) UpdateJobPostController(ctx *gin.Context) {
	id := ctx.Query("id")

	jobPostId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Error while reading parameter id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update job post request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var jobpost dbContext.UpdateJobPostParams

	err = json.Unmarshal(body, &jobpost)
	if err != nil {
		log.Println("Error while unmarshaling update job post request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := jc.jobservice.UpdateJobPost(ctx, &jobpost, int32(jobPostId))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
