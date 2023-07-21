package jobhireService

import (
	"codeid.revampacademy/models"
	// "codeid.revampacademy/repositories"
	"codeid.revampacademy/repositories/jobhireRepositories"
	"github.com/gin-gonic/gin"
)

type JobService struct {
	jobService *jobhireRepositories.JobHirePostRepo
}

func NewJobService(jobService *jobhireRepositories.JobHirePostRepo) *JobService {
	return &JobService{
		jobService: jobService,
	}
}
func (js JobService) GetListJobPost(ctx *gin.Context) ([]*models.JobhireJobPost, *models.ResponseError) {
	return js.jobService.GetListJobPost(ctx)
}

func (js JobService) GetListJobMerge(ctx *gin.Context) ([]*models.MergeJobAndMaster, *models.ResponseError) {
	return js.jobService.GetListJobPostMerge(ctx)
}
