package jobhireService

import (
	"codeid.revampacademy/models"
	feature "codeid.revampacademy/models/features"

	// "codeid.revampacademy/repositories"
	"codeid.revampacademy/repositories/jobhireRepositories"
	"github.com/gin-gonic/gin"
)

type JobService struct {
	// jobService *jobhireRepositories.JobHirePostRepo

	repositoryMgr *jobhireRepositories.RepositoryManager
}

func NewJobService(repoMgr *jobhireRepositories.RepositoryManager) *JobService {
	return &JobService{
		repositoryMgr: repoMgr,
	}
}
func (js JobService) GetListJobPost(ctx *gin.Context, metadata *feature.Metadata) ([]*models.JobhireJobPost, *models.ResponseError) {
	return js.repositoryMgr.GetListJobPost(ctx, metadata)
}

func (js JobService) GetListJobMerge(ctx *gin.Context, metadata *feature.Metadata) ([]*models.MergeJobAndMaster, *models.ResponseError) {
	return js.repositoryMgr.GetListJobPostMerge(ctx, metadata)
}

func (js JobService) GetJobDetailService(ctx *gin.Context, id int32) (*models.MergeJobDetail, *models.ResponseError) {
	return js.repositoryMgr.GetJobRepoDetail(ctx, id)
}

func (js JobService) GetListJobPostSearch(ctx *gin.Context, metadata *feature.Metadata) ([]*models.MergeJobSearch, *models.ResponseError) {
	return js.repositoryMgr.GetListJobPostSearch(ctx, metadata)
}
