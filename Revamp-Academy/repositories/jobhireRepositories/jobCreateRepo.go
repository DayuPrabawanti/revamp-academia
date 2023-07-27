package jobhireRepositories

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/jobhireRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

func (jp JobHirePostRepo) CreateJobPostRepo(ctx *gin.Context, jobPostParams *dbContext.CreateJobPostParams) (*models.CreateJobPost, *models.ResponseError) {
	market := dbContext.New(jp.dbHandler)
	jobPost, err := market.CreateJobPost(ctx, *jobPostParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return jobPost, nil
}
