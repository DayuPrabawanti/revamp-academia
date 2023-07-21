package services

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories"
	"github.com/gin-gonic/gin"
)


type JobPostingService struct {
	masterIndustryServiceRepo *repositories.MasterIndustryRepository
	jobPostServiceRepo *repositories.JobPostRepository
	
}

func NewJobPostingService(jobPostingServiceRepo *repositories.JobPostingRepository) *JobPostingService {
	return &JobPostingService{
		jobPostServiceRepo: jobPostingServiceRepo,
	}
}


func (jpings JobPostingService) ListJobPostRepo(ctx *gin.Context) ([]*models.JobPosting, *models.ResponseError) {
	return jpings.JobPostingServiceRepo.ListJobPostingRepo(ctx)
}