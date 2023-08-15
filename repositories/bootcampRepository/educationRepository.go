package bootcampRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/bootcampRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type EducationRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewEducationRepository(dbHandler *sql.DB) *EducationRepository {
	return &EducationRepository{
		dbHandler: dbHandler,
	}
}

func (ed EducationRepository) GetListEducation(ctx *gin.Context) ([]*models.UsersUsersEducation, *models.ResponseError) {

	store := dbContext.New(ed.dbHandler)
	UsersUsersEducation, err := store.ListUsersEducation(ctx)

	listProgramEntity := make([]*models.UsersUsersEducation, 0)

	for _, v := range UsersUsersEducation {
		programEntity := &models.UsersUsersEducation{
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
		listProgramEntity = append(listProgramEntity, programEntity)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listProgramEntity, nil
}
func (ed EducationRepository) UpdateEducation(ctx *gin.Context, EducationParams *dbContext.UsersUsersEducation) *models.ResponseError {

	store := dbContext.New(ed.dbHandler)
	err := store.UpdateEducation(ctx, *EducationParams)

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
func (ed EducationRepository) GetEducation(ctx *gin.Context, id int64) (*models.UsersUsersEducation, *models.ResponseError) {

	store := dbContext.New(ed.dbHandler)
	education, err := store.GetEducation(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &education, nil
}
