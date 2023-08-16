package jobhireService

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/jobhireRepositories"
	"codeid.revampacademy/repositories/jobhireRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

func (js JobService) CreateJobPostService(ctx *gin.Context, jobParams *dbContext.CreateJobPostParams) (*models.CreateJobPost, *models.ResponseError) {
	responseErr := ValidateParamsJob(jobParams)
	if responseErr != nil {
		return nil, responseErr
	}
	return js.repositoryMgr.CreateJobPostRepo(ctx, jobParams)
}

func ValidateParamsJob(jobParams *dbContext.CreateJobPostParams) *models.ResponseError {
	if jobParams.JobHirePost.JopoTitle == "" {
		return &models.ResponseError{
			Message: "Invalid Job Title",
			Status:  http.StatusInternalServerError,
		}
	}
	if jobParams.JobHirePost.JopoPrimarySkill == "" {
		return &models.ResponseError{
			Message: "No Job Post primary skill Available",
		}
	}
	return nil
}

func (js JobService) CreateJobPostWithDescription(ctx *gin.Context, createJobPostWithDescription *models.CreateJobPost) (*models.CreateJobPost, *models.ResponseError) {

	err := jobhireRepositories.BeginTransaction(js.repositoryMgr)
	if err != nil {
		return nil, &models.ResponseError{
			Message: "Failed to start transaction",
			Status:  http.StatusBadRequest,
		}
	}
	//first query statement
	_, responseErr := js.CreateJobPostService(ctx, (*dbContext.CreateJobPostParams)(createJobPostWithDescription))
	if responseErr != nil {
		jobhireRepositories.RollbackTransaction(js.repositoryMgr)
		return nil, responseErr
	}

	// //second query statement
	// responseErr = cs.DeleteCategory(ctx, int64(response.CategoryID))
	// if responseErr != nil {
	// 	//when delete not succeed, transaction will rollback
	// 	repositories.RollbackTransaction(cs.repositoryMgr)
	// 	return nil, responseErr
	// }

	// if all statement ok, transaction will commit/save to db
	jobhireRepositories.CommitTransaction(js.repositoryMgr)
	return nil, &models.ResponseError{
		Message: "Data has been created",
		Status:  http.StatusOK,
	}
}
