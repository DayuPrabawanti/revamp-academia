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
	transaction *sql.Tx
	MasterIndustryRepository
	JobPostRepository

}

func NewJobPostingRepository(dbHandler *sql.DB) *JobPostingRepository {
    return &JobPostingRepository{
        dbHandler: dbHandler,
    }
}


func (jpingr JobPostingRepository) ListJobPostingRepo (ctx *gin.Context) ([]*models.JobPosting, *models.ResponseError) {

	store := dbContext.New(jpingr.dbHandler)
	jobPostingGroup, err := store.ListJobPostingImpl(ctx)

	jobPostingMakeList := make([]*models.JobPosting, 0)

	for _, v := range jobPostingGroup {
		jobPostingData := &models.JobPosting{
			JobhireJobPost: v.JobhireJobPost,
			MasterIndustry:       v.MasterIndustry,
		}
		jobPostingMakeList = append(jobPostingMakeList, jobPostingData)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return jobPostingMakeList, nil
}