package users

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models/usersModel"
	"codeid.revampacademy/repositories/users/dbContext"
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

func (cr UserRepository) GetListUsers(ctx *gin.Context) ([]*usersModel.UsersUser, *usersModel.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	users, err := store.ListUsers(ctx)

	listUsers := make([]*usersModel.UsersUser, 0)

	for _, v := range users {
		user := &usersModel.UsersUser{
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
		listUsers = append(listUsers, user)
	}

	if err != nil {
		return nil, &usersModel.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listUsers, nil
}

func (cr UserRepository) GetUser(ctx *gin.Context, id int32) (*usersModel.UsersUser, *usersModel.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	category, err := store.GetUser(ctx, int32(id))

	if err != nil {
		return nil, &usersModel.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &category, nil
}

func (cr UserRepository) CreateUser(ctx *gin.Context, userParams *dbContext.CreateUsersParams) (*usersModel.UsersUser, *usersModel.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	user, err := store.CreateUsers(ctx, *userParams)

	if err != nil {
		return nil, &usersModel.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return user, nil
}

func (cr UserRepository) UpdateUser(ctx *gin.Context, userParams *dbContext.CreateUsersParams) *usersModel.ResponseError {

	store := dbContext.New(cr.dbHandler)
	err := store.UpdateUser(ctx, *userParams)

	if err != nil {
		return &usersModel.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &usersModel.ResponseError{
		Message: "data has been update",
		Status:  http.StatusOK,
	}
}

func (cr UserRepository) DeleteCategory(ctx *gin.Context, id int32) *usersModel.ResponseError {

	store := dbContext.New(cr.dbHandler)
	err := store.DeleteUsers(ctx, int32(id))

	if err != nil {
		return &usersModel.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &usersModel.ResponseError{
		Message: "data has been deleted",
		Status:  http.StatusOK,
	}
}