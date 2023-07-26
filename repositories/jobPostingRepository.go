package repositories

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/dbContext"
	"github.com/gin-gonic/gin"
)

type JobPostingRepository struct {
	dbHandler   *sql.DB
	Transaction *sql.Tx
	MasterIndustryRepository
	JobPostRepository
}

func NewJobPostingRepository(dbHandler *sql.DB) *JobPostingRepository {
	return &JobPostingRepository{
		dbHandler: dbHandler,
	}
}

func (jpingr JobPostingRepository) GetJobPostingRepo(ctx *gin.Context, title string) (*models.JobPosting, *models.ResponseError) {

	store2 := dbContext.New(jpingr.dbHandler)
	jobPosting, err := store2.GetJobPostingImpl(ctx, title)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &jobPosting, nil
}

func (jpingr JobPostingRepository) ListJobPostingRepo(ctx *gin.Context, nama string) ([]*models.JobPosting, *models.ResponseError) {

	store := dbContext.New(jpingr.dbHandler)
	jobPostingGroup, err := store.ListJobPostingImpl(ctx, string(nama))

	jobPostingMakeList := make([]*models.JobPosting, 0)

	for _, v := range jobPostingGroup {
		jobPostingGroup := &models.JobPosting{
			JobhireJobPost: v.JobhireJobPost,
			MasterIndustry: v.MasterIndustry,
		}
		jobPostingMakeList = append(jobPostingMakeList, jobPostingGroup)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return jobPostingMakeList, nil
}
