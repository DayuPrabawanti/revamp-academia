package services

import (
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories"
	"codeid.revampacademy/repositories/dbContext"
	"github.com/gin-gonic/gin"
)

type JobCategoryService struct {
	JobCategoryServiceRepo *repositories.JobCategoryRepository
}

type JobClientService struct {
	JobClientServiceRepo *repositories.JobClientRepository
}

type JobPostService struct {
	JobPostServiceRepo *repositories.JobPostRepository
}


// --------------------------- JOB CATEGORY --------------------------- 


func NewJobCategoryService(jobhireCategory *repositories.JobCategoryRepository) *JobCategoryService {
	return &JobCategoryService{
		JobCategoryServiceRepo: jobhireCategory,
	}
}

func (jcs JobCategoryService) GetJobCategoryService(ctx *gin.Context, id int64) (*models.JobhireJobCategory, *models.ResponseError) {
			return jcs.JobCategoryServiceRepo.GetJobCategoryRepo(ctx, id)
}

	func (jcs JobCategoryService) ListJobCategoryService(ctx *gin.Context) ([]*models.JobhireJobCategory, *models.ResponseError) {
		return jcs.JobCategoryServiceRepo.ListJobCategoryRepo(ctx)
	}

		func (jcs JobCategoryService) CreateJobCategoryService(ctx *gin.Context, CreateJobCategoryParams *dbContext.CreateJobCategoryParams) (*models.JobhireJobCategory, *models.ResponseError) {
				responseErr := validateJobCategory(CreateJobCategoryParams)
				if responseErr != nil {
					return nil, responseErr
				}

				return jcs.JobCategoryServiceRepo.CreateJobCategoryRepo(ctx, CreateJobCategoryParams)
		}

			func (jcs JobCategoryService) UpdateCategoryService(ctx *gin.Context, CreateJobCategoryParams *dbContext.CreateJobCategoryParams, id int64) *models.ResponseError {
					responseErr := validateJobCategory(CreateJobCategoryParams)
					if responseErr != nil {
						return responseErr
					}

					return jcs.JobCategoryServiceRepo.UpdateJobCategoryRepo(ctx, CreateJobCategoryParams)
			}

				func (jcs JobCategoryService) DeleteCategoryService(ctx *gin.Context, id int64) *models.ResponseError {
									return jcs.JobCategoryServiceRepo.DeleteJobCategoryRepo(ctx, id)
				}


func validateJobCategory(CreateJobCategoryParams *dbContext.CreateJobCategoryParams) *models.ResponseError {
	if CreateJobCategoryParams.JocaID == 0 {
		return &models.ResponseError{
			Message: "Invalid category id",
			Status:  http.StatusBadRequest,
		}
	}

	if CreateJobCategoryParams.JocaName == "" {
		return &models.ResponseError{
			Message: "Invalid category name",
			Status:  http.StatusBadRequest,
		}
	}

	return nil
}

// --------------------------- JOB CLIENT --------------------------- 

func NewJobClientService(JobClientServiceRepo *repositories.JobClientRepository) *JobClientService {
		return &JobClientService{
			JobClientServiceRepo: JobClientServiceRepo,
		}
}

func (jclient JobClientService) GetJobClientService(ctx *gin.Context, id int64) (*models.JobhireClient, *models.ResponseError) {
	return jclient.JobClientServiceRepo.GetJobClientRepo(ctx, id)
}

	func (jclient JobClientService) ListJobClientService(ctx *gin.Context) ([]*models.JobhireClient, *models.ResponseError) {
			return jclient.JobClientServiceRepo.ListJobClientRepo(ctx)
	}

		func (jclient JobClientService) CreateJobClientService(ctx *gin.Context, CreateJobClientParams *dbContext.CreateJobClientParams) (*models.JobhireClient, *models.ResponseError) {
				responseErr := validateJobClient(CreateJobClientParams)
				if responseErr != nil {
					return nil, responseErr
				}
				return jclient.JobClientServiceRepo.CreateJobClientRepo(ctx,CreateJobClientParams)
		}

			func (jclient JobClientService) UpdateJobClientService(ctx *gin.Context, CreateJobClientParams *dbContext.CreateJobClientParams, id int64) *models.ResponseError {
								responseErr := validateJobClient(CreateJobClientParams)
								if responseErr != nil {
									return responseErr
								}

								return jclient.JobClientServiceRepo.UpdateJobClientRepo(ctx, CreateJobClientParams)
			}

				func (jclient JobClientService) DeleteJobClientService(ctx *gin.Context, id int64) *models.ResponseError {
							return jclient.JobClientServiceRepo.DeleteJobClientRepo(ctx, id)
				}

func validateJobClient(clientParams *dbContext.CreateJobClientParams) *models.ResponseError {
	if clientParams.ClitID == 0 {
		return &models.ResponseError{
			Message: "Invalid category id",
			Status:  http.StatusBadRequest,
		}
	}

	if clientParams.ClitName == "" {
		return &models.ResponseError{
			Message: "Invalid category name",
			Status:  http.StatusBadRequest,
		}
	}

	return nil
}

// --------------------------- JOB POST --------------------------- 

func NewJobPostService(JobPostServiceRepo *repositories.JobPostRepository) *JobPostService {
		return &JobPostService{
			JobPostServiceRepo: JobPostServiceRepo,
		}
}

func (jpost JobPostService) GetJobPostService(ctx *gin.Context, id int32) (*models.JobhireJobPost, *models.ResponseError) {
		return jpost.JobPostServiceRepo.GetJobPostRepo(ctx, id)
}

	func (jpost JobPostService) ListJobPostService(ctx *gin.Context) ([]*models.JobhireJobPost, *models.ResponseError) {
				return jpost.JobPostServiceRepo.ListJobPostRepo(ctx)
	}

		func (jpost JobPostService) CreateJobPostService(ctx *gin.Context, CreateJobPostParams *dbContext.CreateJobPostParams) (*models.JobhireJobPost, *models.ResponseError) {
					responseErr := validateJobPost(CreateJobPostParams)
					if responseErr != nil {
						return nil, responseErr
					}
					return jpost.JobPostServiceRepo.CreateJobPostRepo(ctx,CreateJobPostParams)
		}
		

func validateJobPost(jobPostParams *dbContext.CreateJobPostParams) *models.ResponseError {
	if jobPostParams.JopoEntityID == 0 {
		return &models.ResponseError{
			Message: "Invalid category id",
			Status:  http.StatusBadRequest,
		}
	}

	if jobPostParams.JopoTitle == "" {
		return &models.ResponseError{
			Message: "Invalid category name",
			Status:  http.StatusBadRequest,
		}
	}

	return nil
}