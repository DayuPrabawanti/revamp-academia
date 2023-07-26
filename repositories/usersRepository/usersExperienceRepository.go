package usersRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/usersRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type UserExperienceRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewUserExperienceRepository(dbHandler *sql.DB) *UserExperienceRepository {
	return &UserExperienceRepository{
		dbHandler: dbHandler,
	}
}

func (cr UserExperienceRepository) GetListUserExperience(ctx *gin.Context) ([]*models.UsersUsersExperience, *models.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	usersExperience, err := store.ListExperience(ctx)

	listUsersExperience := make([]*models.UsersUsersExperience, 0)

	for _, v := range usersExperience {
		userEx := &models.UsersUsersExperience{
			UsexID: v.UsexID,
			UsexEntityID: v.UsexEntityID,
			UsexTitle: v.UsexTitle,
			UsexProfileHeadline: v.UsexProfileHeadline,
			UsexEmploymentType: v.UsexEmploymentType,
			UsexCompanyName: v.UsexCompanyName,
			UsexIsCurrent: v.UsexIsCurrent,
			UsexStartDate: v.UsexStartDate,
			UsexEndDate: v.UsexEndDate,
			UsexIndustry: v.UsexIndustry,
			UsexDescription: v.UsexDescription,
			UsexExperienceType: v.UsexExperienceType,
			UsexCityID: v.UsexCityID,


		}
		listUsersExperience = append(listUsersExperience, userEx)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listUsersExperience, nil
}

func (cr UserExperienceRepository) GetExperience(ctx *gin.Context, id int32) (*models.UsersUsersExperience, *models.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	userEx, err := store.GetExperience(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &userEx, nil
}

func (cr UserExperienceRepository) CreateExperience(ctx *gin.Context, userExperienceParams *dbContext.CreateExperienceParams) (*models.UsersUsersExperience, *models.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	userEx, err := store.CreateExperience(ctx, *userExperienceParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return userEx, nil
}

func (cr UserExperienceRepository) UpdateUserExperience(ctx *gin.Context, experienceParams *dbContext.CreateExperienceParams) *models.ResponseError {

	store := dbContext.New(cr.dbHandler)
	err := store.UpdateExperience(ctx, *experienceParams)

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

func (cr UserExperienceRepository) DeleteExperience(ctx *gin.Context, id int32) *models.ResponseError {

	store := dbContext.New(cr.dbHandler)
	err := store.DeleteExperience(ctx, int32(id))

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