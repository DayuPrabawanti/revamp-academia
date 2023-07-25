package usersRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/usersRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type UserLicenseRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewUserLicenseRepository(dbHandler *sql.DB) *UserLicenseRepository {
	return &UserLicenseRepository{
		dbHandler: dbHandler,
	}
}

// GetList Users License
func (cr UserLicenseRepository) GetListUsersLicense(ctx *gin.Context) ([]*models.UsersUsersLicense, *models.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	usersLicense, err := store.ListLicense(ctx)

	listUsersLicense := make([]*models.UsersUsersLicense, 0)

	for _, v := range usersLicense {
		usersLicense := &models.UsersUsersLicense{
			UsliID:           v.UsliID,
			UsliLicenseCode:  v.UsliLicenseCode,
			UsliModifiedDate: v.UsliModifiedDate,
			UsliStatus:       v.UsliStatus,
			UsliEntityID:     v.UsliEntityID,
		}
		listUsersLicense = append(listUsersLicense, usersLicense)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listUsersLicense, nil
}

// Get Users License
func (cr UserLicenseRepository) GetUserLicense(ctx *gin.Context, id int32) (*models.UsersUsersLicense, *models.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	userLicense, err := store.GetLicense(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &userLicense, nil
}

// Create User License
func (cr UserLicenseRepository) CreateUserLicense(ctx *gin.Context, LicenseParams *dbContext.CreateLicenseParams) (*models.UsersUsersLicense, *models.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	userLicense, err := store.CreateLicense(ctx, *LicenseParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return userLicense, nil
}

// Update Table
func (cr UserLicenseRepository) UpdateLicense(ctx *gin.Context, userLicenseParams *dbContext.CreateLicenseParams) *models.ResponseError {

	store := dbContext.New(cr.dbHandler)
	err := store.UpdateLicense(ctx, *userLicenseParams)

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
func (cr UserLicenseRepository) DeleteLicense(ctx *gin.Context, id int32) *models.ResponseError {

	store := dbContext.New(cr.dbHandler)
	err := store.DeleteLicense(ctx, int32(id))

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
