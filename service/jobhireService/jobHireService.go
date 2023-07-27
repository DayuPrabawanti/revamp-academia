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

func (js JobService) GetJobDetailService(ctx *gin.Context, id int32) (*models.MergeJobDetail, *models.ResponseError) {
	return js.jobService.GetJobRepoDetail(ctx, id)
}

func (js JobService) GetListJobPostSearch(ctx *gin.Context, cityName string, joroName string, wotyName string) ([]*models.MergeJobSearch, *models.ResponseError) {
	return js.jobService.GetListJobPostSearch(ctx, cityName, joroName, wotyName)
}
