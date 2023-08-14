package jobhireRepositories

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/jobhireRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

func (jp JobHirePostRepo) UpdateJobPosting(ctx *gin.Context, jobPostParams *dbContext.UpdateJobPostParams) *models.ResponseError {

	market := dbContext.New(jp.dbHandler)
	err := market.UpdateJobPost(ctx, *jobPostParams)

	if err != nil {
		return &models.ResponseError{
			Message: "Error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "Data has been update",
		Status:  http.StatusOK,
	}
}
