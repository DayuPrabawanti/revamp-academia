package usersRepository

import (
	"database/sql"
	"net/http"

	model "codeid.revampacademy/models"
	dbcontext "codeid.revampacademy/repositories/usersRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type UserMediaRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewUserMediaRepository(dbHandler *sql.DB) *UserMediaRepository {
	return &UserMediaRepository{
		dbHandler: dbHandler,
	}
}

// GetList User Media
func (cr UserMediaRepository) GetListUserMedia(ctx *gin.Context) ([]*model.UsersUsersMedia, *model.ResponseError) {

	store := dbcontext.New(cr.dbHandler)
	usersMedia, err := store.ListMedia(ctx)

	listUsersMedia := make([]*model.UsersUsersMedia, 0)

	for _, v := range usersMedia {
		userMedia := &model.UsersUsersMedia{
			UsmeID:           v.UsmeID,
			UsmeEntityID:     v.UsmeEntityID,
			UsmeFileLink:     v.UsmeFileLink,
			UsmeFilename:     v.UsmeFilename,
			UsmeFilesize:     v.UsmeFilesize,
			UsmeFiletype:     v.UsmeFiletype,
			UsmeNote:         v.UsmeNote,
			UsmeModifiedDate: v.UsmeModifiedDate,
		}
		listUsersMedia = append(listUsersMedia, userMedia)
	}

	if err != nil {
		return nil, &model.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listUsersMedia, nil
}

// Get User Media
func (cr UserMediaRepository) GetUserMedia(ctx *gin.Context, id int32) (*model.UsersUsersMedia, *model.ResponseError) {

	store := dbcontext.New(cr.dbHandler)
	userMedia, err := store.GetMedia(ctx, int32(id))

	if err != nil {
		return nil, &model.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &userMedia, nil
}

// Create User media
func (cr UserMediaRepository) CreateUserMedia(ctx *gin.Context, mediaParams *dbcontext.CreateMediaParams) (*model.UsersUsersMedia, *model.ResponseError) {

	store := dbcontext.New(cr.dbHandler)
	media, err := store.CreateMedia(ctx, *mediaParams)

	if err != nil {
		return nil, &model.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return media, nil
}

// Update Table
func (cr UserMediaRepository) UpdateMedia(ctx *gin.Context, userMediaParams *dbcontext.CreateMediaParams) *model.ResponseError {

	store := dbcontext.New(cr.dbHandler)
	err := store.UpdateMedia(ctx, *userMediaParams)

	if err != nil {
		return &model.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &model.ResponseError{
		Message: "data has been update",
		Status:  http.StatusOK,
	}
}

// Delete Table
func (cr UserMediaRepository) DeleteMedia(ctx *gin.Context, id int32) *model.ResponseError {

	store := dbcontext.New(cr.dbHandler)
	err := store.DeleteMedia(ctx, int32(id))

	if err != nil {
		return &model.ResponseError{
			Message: "error when delete",
			Status:  http.StatusInternalServerError,
		}
	}
	return &model.ResponseError{
		Message: "data has been deleted",
		Status:  http.StatusOK,
	}
}
