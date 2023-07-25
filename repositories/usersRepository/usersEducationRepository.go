package usersRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	dbcontext "codeid.revampacademy/repositories/usersRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type UserEducationRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewUserEducationRepository(dbHandler *sql.DB) *UserEducationRepository {
	return &UserEducationRepository{
		dbHandler: dbHandler,
	}
}

// GetList Users Education
func (cr UserEducationRepository) GetListUsersEducation(ctx *gin.Context) ([]*models.UsersUsersEducation, *models.ResponseError) {

	store := dbcontext.New(cr.dbHandler)
	userEducation, err := store.ListEducation(ctx)

	listUsersEducation := make([]*models.UsersUsersEducation, 0)

	for _, v := range userEducation {
		userEducation := &models.UsersUsersEducation{
			UsduID:           v.UsduID,
			UsduEntityID:     v.UsduEntityID,
			UsduSchool:       v.UsduSchool,
			UsduDegree:       v.UsduDegree,
			UsduFieldStudy:   v.UsduFieldStudy,
			UsduGraduateYear: v.UsduGraduateYear,
			UsduStartDate:    v.UsduStartDate,
			UsduEndDate:      v.UsduEndDate,
			UsduGrade:        v.UsduGrade,
			UsduActivities:   v.UsduActivities,
			UsduDescription:  v.UsduDescription,
			UsduModifiedDate: v.UsduModifiedDate,
		}
		listUsersEducation = append(listUsersEducation, userEducation)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listUsersEducation, nil
}

// Get User Education
func (cr UserEducationRepository) GetUserEducation(ctx *gin.Context, id int32) (*models.UsersUsersEducation, *models.ResponseError) {

	store := dbcontext.New(cr.dbHandler)
	userEducation, err := store.GetEducation(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &userEducation, nil
}

// Create User Education
func (cr UserEducationRepository) CreateUserEducation(ctx *gin.Context, educationPrams *dbcontext.CreateEducationParams) (*models.UsersUsersEducation, *models.ResponseError) {

	store := dbcontext.New(cr.dbHandler)
	education, err := store.CreateEducation(ctx, *educationPrams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return education, nil
}

// Update Table
func (cr UserEducationRepository) UpdateEducation(ctx *gin.Context, userEducationParams *dbcontext.CreateEducationParams) *models.ResponseError {

	store := dbcontext.New(cr.dbHandler)
	err := store.UpdateEducation(ctx, *userEducationParams)

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

// Delete Table
func (cr UserEducationRepository) DeleteEducation(ctx *gin.Context, id int32) *models.ResponseError {

	store := dbcontext.New(cr.dbHandler)
	err := store.DeleteEducation(ctx, int32(id))

	if err != nil {
		return &models.ResponseError{
			Message: "error when delete",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been deleted",
		Status:  http.StatusOK,
	}
}
