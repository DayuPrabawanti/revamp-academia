package users

import (
	"database/sql"
	"net/http"

	"codeid.revamptwo/models/usersm"
	"codeid.revamptwo/repositories/users/dbContext"
	"github.com/gin-gonic/gin"
)

type UserRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewUserRepository(dbHandler *sql.DB) *UserRepository {
	return &UserRepository{
		dbHandler: dbHandler,
	}
}

func (cr UserRepository) GetListUsers(ctx *gin.Context) ([]*usersm.UsersUser, *usersm.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	categories, err := store.ListUsers(ctx)

	listCategories := make([]*usersm.UsersUser, 0)

	for _, v := range categories {
		category := &usersm.UsersUser{
			UserEntityID:       v.UserEntityID,
			UserName:           v.UserName,
			UserPassword:       v.UserPassword,
			UserFirstName:      v.UserFirstName,
			UserLastName:       v.UserLastName,
			UserBirthDate:      v.UserBirthDate,
			UserEmailPromotion: v.UserEmailPromotion,
			UserDemographic:    v.UserDemographic,
			UserModifiedDate:   v.UserModifiedDate,
			UserPhoto:          v.UserPhoto,
			UserCurrentRole:    v.UserCurrentRole,
		}
		listCategories = append(listCategories, category)
	}

	if err != nil {
		return nil, &usersm.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listCategories, nil
}

func (cr UserRepository) GetUser(ctx *gin.Context, id int32) (*usersm.UsersUser, *usersm.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	category, err := store.GetUser(ctx, int32(id))

	if err != nil {
		return nil, &usersm.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &category, nil
}

func (cr UserRepository) CreateUser(ctx *gin.Context, userParams *dbContext.CreateUsersParams) (*usersm.UsersUser, *usersm.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	category, err := store.CreateUsers(ctx, *userParams)

	if err != nil {
		return nil, &usersm.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return category, nil
}

func (cr UserRepository) UpdateUser(ctx *gin.Context, userParams *dbContext.CreateUsersParams) *usersm.ResponseError {

	store := dbContext.New(cr.dbHandler)
	err := store.UpdateUser(ctx, *userParams)

	if err != nil {
		return &usersm.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &usersm.ResponseError{
		Message: "data has been update",
		Status:  http.StatusOK,
	}
}

func (cr UserRepository) DeleteCategory(ctx *gin.Context, id int32) *usersm.ResponseError {

	store := dbContext.New(cr.dbHandler)
	err := store.DeleteUsers(ctx, int32(id))

	if err != nil {
		return &usersm.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &usersm.ResponseError{
		Message: "data has been deleted",
		Status:  http.StatusOK,
	}
}