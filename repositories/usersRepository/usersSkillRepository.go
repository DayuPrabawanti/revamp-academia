package usersRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/usersRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type UserSkillRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewUserSkillRepository(dbHandler *sql.DB) *UserSkillRepository {
	return &UserSkillRepository{
		dbHandler: dbHandler,
	}
}

// GetList Users Skill
func (cr UserSkillRepository) GetListUsersSkill(ctx *gin.Context) ([]*models.UsersUsersSkill, *models.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	usersSkill, err := store.ListSkill(ctx)

	listUsersSkill := make([]*models.UsersUsersSkill, 0)

	for _, v := range usersSkill {
		usersSkill := &models.UsersUsersSkill{
			UskiID:           v.UskiID,
			UskiEntityID:     v.UskiEntityID,
			UskiModifiedDate: v.UskiModifiedDate,
			UskiSktyName:     v.UskiSktyName,
		}
		listUsersSkill = append(listUsersSkill, usersSkill)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listUsersSkill, nil
}

// Get Users Skill
func (cr UserSkillRepository) GetUserSkill(ctx *gin.Context, id int32) (*models.UsersUsersSkill, *models.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	userSkill, err := store.GetSkill(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &userSkill, nil
}

// Create User Skill
func (cr UserSkillRepository) CreateUserSkill(ctx *gin.Context, SkillParams *dbContext.CreateSkillParams) (*models.UsersUsersSkill, *models.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	userSkill, err := store.CreateSkill(ctx, *SkillParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return userSkill, nil
}

// Update Table
func (cr UserSkillRepository) UpdateSkill(ctx *gin.Context, userSkillParams *dbContext.CreateSkillParams) *models.ResponseError {

	store := dbContext.New(cr.dbHandler)
	err := store.UpdateSkill(ctx, *userSkillParams)

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
func (cr UserSkillRepository) DeleteSkill(ctx *gin.Context, id int32) *models.ResponseError {

	store := dbContext.New(cr.dbHandler)
	err := store.DeleteSkill(ctx, int32(id))

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
