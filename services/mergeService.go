package services

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories"
	"github.com/gin-gonic/gin"
)

type JobPostingService struct {
	jobPostingRepository *repositories.JobPostingRepository
}

func NewJobPostingService(jobPostingRepository *repositories.JobPostingRepository) *JobPostingService {
	return &JobPostingService{
		jobPostingRepository: jobPostingRepository,
	}
}

func (jpings JobPostingService) GetJobPostingService(ctx *gin.Context, title string) (*models.JobPosting, *models.ResponseError) {
	return jpings.jobPostingRepository.GetJobPostingRepo(ctx, title)
}

func (jpings JobPostingService) ListJobPostRepo(ctx *gin.Context, nama string) ([]*models.JobPosting, *models.ResponseError) {
	return jpings.jobPostingRepository.ListJobPostingRepo(ctx, nama)
}
