package repositories

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/dbContext"
	"github.com/gin-gonic/gin"
)

type JobCategoryRepository struct {
	dbHandler   *sql.DB
	Transaction *sql.Tx
}

type JobClientRepository struct {
	dbHandler   *sql.DB
	Transaction *sql.Tx
}

type JobPostRepository struct {
	dbHandler   *sql.DB
	Transaction *sql.Tx
}

// --------------------------- JOB CATEGORY ---------------------------

func NewJobCategoryRepository(dbHandler *sql.DB) *JobCategoryRepository {
	return &JobCategoryRepository{
		dbHandler: dbHandler,
	}
}

func (jcr JobCategoryRepository) GetJobCategoryRepo(ctx *gin.Context, id int64) (*models.JobhireJobCategory, *models.ResponseError) {
	store := dbContext.New(jcr.dbHandler)
	category, err := store.GetJobCategoryImpl(ctx, int16(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &category, nil
}

func (jcr JobCategoryRepository) ListJobCategoryRepo(ctx *gin.Context) ([]*models.JobhireJobCategory, *models.ResponseError) {

	store := dbContext.New(jcr.dbHandler)
	jobhireCategory, err := store.ListJobCategoryImpl(ctx)

	listCategories := make([]*models.JobhireJobCategory, 0)

	for _, v := range jobhireCategory {
		jobCategory := &models.JobhireJobCategory{
			JocaID:           v.JocaID,
			JocaName:         v.JocaName,
			JocaModifiedDate: v.JocaModifiedDate,
		}
		listCategories = append(listCategories, jobCategory)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return listCategories, nil
}

func (jcr JobCategoryRepository) CreateJobCategoryRepo(ctx *gin.Context, CreateJobCategoryParams *dbContext.CreateJobCategoryParams) (*models.JobhireJobCategory, *models.ResponseError) {

	store := dbContext.New(jcr.dbHandler)
	jobCategory, err := store.CreateJobCategoryImpl(ctx, *CreateJobCategoryParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return jobCategory, nil
}

func (jcr JobCategoryRepository) UpdateJobCategoryRepo(ctx *gin.Context, CreateJobCategoryParams *dbContext.CreateJobCategoryParams) *models.ResponseError {

	store := dbContext.New(jcr.dbHandler)
	err := store.UpdateJobCategoryImpl(ctx, *CreateJobCategoryParams)

	if err != nil {
		return &models.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been update",
		Status:  http.StatusOK,
	}
}

func (jcr JobCategoryRepository) DeleteJobCategoryRepo(ctx *gin.Context, id int64) *models.ResponseError {

	store := dbContext.New(jcr.dbHandler)
	err := store.DeleteJobCategoryImpl(ctx, int32(id))

	if err != nil {
		return &models.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been deleted",
		Status:  http.StatusOK,
	}
}

// --------------------------- JOB CLIENT ---------------------------

func NewJobClientRepository(dbHandler *sql.DB) *JobClientRepository {
	return &JobClientRepository{
		dbHandler: dbHandler,
	}
}

func (client JobClientRepository) GetJobClientRepo(ctx *gin.Context, id int64) (*models.JobhireClient, *models.ResponseError) {

	store2 := dbContext.New(client.dbHandler)
	clientCategory, err := store2.GetJobClientImpl(ctx, int64(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &clientCategory, nil
}

func (client JobClientRepository) ListJobClientRepo(ctx *gin.Context) ([]*models.JobhireClient, *models.ResponseError) {

	store := dbContext.New(client.dbHandler)
	jobhireClient, err := store.ListJobClientImpl(ctx)

	listClient := make([]*models.JobhireClient, 0)

	for _, v := range jobhireClient {
		jobClient := &models.JobhireClient{
			ClitID:     v.ClitID,
			ClitName:   v.ClitName,
			ClitAbout:  v.ClitAbout,
			ClitAddrID: v.ClitAddrID,
			ClitEmraID: v.ClitEmraID,
		}
		listClient = append(listClient, jobClient)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return listClient, nil
}

func (client JobClientRepository) CreateJobClientRepo(ctx *gin.Context, CreateJobClientParams *dbContext.CreateJobClientParams) (*models.JobhireClient, *models.ResponseError) {

	store := dbContext.New(client.dbHandler)
	jobClient, err2 := store.CreateJobClientImpl(ctx, *CreateJobClientParams)

	if err2 != nil {
		return nil, &models.ResponseError{
			Message: err2.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return jobClient, nil
}

func (client JobClientRepository) UpdateJobClientRepo(ctx *gin.Context, CreateJobClientParams *dbContext.CreateJobClientParams) *models.ResponseError {

	store := dbContext.New(client.dbHandler)
	err := store.UpdateJobClientImpl(ctx, *CreateJobClientParams)

	if err != nil {
		return &models.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been update",
		Status:  http.StatusOK,
	}
}

func (client JobClientRepository) DeleteJobClientRepo(ctx *gin.Context, id int64) *models.ResponseError {

	store := dbContext.New(client.dbHandler)
	err := store.DeleteJobClientImpl(ctx, int32(id))

	if err != nil {
		return &models.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been deleted",
		Status:  http.StatusOK,
	}
}

// --------------------------- JOB POST ---------------------------

func NewJobPostRepository(dbHandler *sql.DB) *JobPostRepository {
	return &JobPostRepository{
		dbHandler: dbHandler,
	}
}

func (jobPost JobPostRepository) GetJobPostRepo(ctx *gin.Context, id int32) (*models.JobhireJobPost, *models.ResponseError) {

	store2 := dbContext.New(jobPost.dbHandler)
	jobPostCategory, err := store2.GetJobPostImpl(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &jobPostCategory, nil
}

func (jobPost JobPostRepository) ListJobPostRepo(ctx *gin.Context) ([]*models.JobhireJobPost, *models.ResponseError) {

	store := dbContext.New(jobPost.dbHandler)
	jobhirePost, err := store.ListJobPostImpl(ctx)

	listJobPost := make([]*models.JobhireJobPost, 0)

	for _, v := range jobhirePost {
		jobPost := &models.JobhireJobPost{
			JopoEntityID:       v.JopoEntityID,
			JopoNumber:         v.JopoNumber,
			JopoTitle:          v.JopoTitle,
			JopoStartDate:      v.JopoStartDate,
			JopoEndDate:        v.JopoEndDate,
			JopoMinSalary:      v.JopoMinSalary,
			JopoMaxSalary:      v.JopoMaxSalary,
			JopoMinExperience:  v.JopoMinExperience,
			JopoMaxExperience:  v.JopoMaxExperience,
			JopoPrimarySkill:   v.JopoPrimarySkill,
			JopoSecondarySkill: v.JopoSecondarySkill,
			JopoPublishDate:    v.JopoPublishDate,
			JopoModifiedDate:   v.JopoModifiedDate,
			JopoEmpEntityID:    v.JopoEmpEntityID,
			JopoClitID:         v.JopoClitID,
			JopoJoroID:         v.JopoJoroID,
			JopoJotyID:         v.JopoJotyID,
			JopoJocaID:         v.JopoJocaID,
			JopoAddrID:         v.JopoAddrID,
			JopoWorkCode:       v.JopoWorkCode,
			JopoEduCode:        v.JopoEduCode,
			JopoInduCode:       v.JopoInduCode,
			JopoStatus:         v.JopoStatus,
		}
		listJobPost = append(listJobPost, jobPost)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return listJobPost, nil
}

func (jobPost JobPostRepository) CreateJobPostRepo(ctx *gin.Context, CreateJobPostParams *dbContext.CreateJobPostParams) (*models.JobhireJobPost, *models.ResponseError) {

	store := dbContext.New(jobPost.dbHandler)
	jobClient, err2 := store.CreateJobPostImpl(ctx, *CreateJobPostParams)

	if err2 != nil {
		return nil, &models.ResponseError{
			Message: err2.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return jobClient, nil
}

func (jobPost JobPostRepository) UpdateJobPostRepo(ctx *gin.Context, CreateJobPostParams *dbContext.CreateJobPostParams) *models.ResponseError {

	store := dbContext.New(jobPost.dbHandler)
	err := store.UpdateJobPostImpl(ctx, *CreateJobPostParams)

	if err != nil {
		return &models.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been update",
		Status:  http.StatusOK,
	}
}

func (jobPost JobPostRepository) DeleteJobPostRepo(ctx *gin.Context, id int64) *models.ResponseError {

	store := dbContext.New(jobPost.dbHandler)
	err := store.DeleteJobPostImpl(ctx, int32(id))

	if err != nil {
		return &models.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been deleted",
		Status:  http.StatusOK,
	}
}
