package users

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models/usersModel"
	"codeid.revampacademy/repositories/users/dbContext"
	"github.com/gin-gonic/gin"
)

type UserPhoneRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewUserPhoneRepository(dbHandler *sql.DB) *UserPhoneRepository {
	return &UserPhoneRepository{
		dbHandler: dbHandler,
	}
}

func (cr UserPhoneRepository) GetListUsersPhone(ctx *gin.Context) ([]*usersModel.UsersUsersPhone, *usersModel.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	usersPhone, err := store.ListPhones(ctx)

	listUsersPhone := make([]*usersModel.UsersUsersPhone, 0)

	for _, v := range usersPhone {
		userPhone := &usersModel.UsersUsersPhone{
			UspoEntityID: v.UspoEntityID,
			UspoNumber: v.UspoNumber,
			UspoModifiedDate: v.UspoModifiedDate,
			UspoPontyCode: v.UspoPontyCode,
		}
		listUsersPhone = append(listUsersPhone, userPhone)
	}

	if err != nil {
		return nil, &usersModel.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listUsersPhone, nil
}

func (cr UserPhoneRepository) GetPhone(ctx *gin.Context, id int32) (*usersModel.UsersUsersPhone, *usersModel.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	userPhone, err := store.GetPhones(ctx, int32(id))

	if err != nil {
		return nil, &usersModel.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &userPhone, nil
}

func (cr UserPhoneRepository) CreatePhones(ctx *gin.Context, phoneParams *dbContext.CreatePhonesParams) (*usersModel.UsersUsersPhone, *usersModel.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	phone, err := store.CreatePhones(ctx, *phoneParams)

	if err != nil {
		return nil, &usersModel.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return phone, nil
}

func (cr UserPhoneRepository) UpdatePhone(ctx *gin.Context, userPhoneParams *dbContext.CreatePhonesParams) *usersModel.ResponseError {

	store := dbContext.New(cr.dbHandler)
	err := store.UpdatePhones(ctx, *userPhoneParams)

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

func (cr UserPhoneRepository) DeletePhones(ctx *gin.Context, id int32) *usersModel.ResponseError {

	store := dbContext.New(cr.dbHandler)
	err := store.DeletePhones(ctx, int32(id))

	if err != nil {
		return &usersModel.ResponseError{
			Message: "error when delete",
			Status:  http.StatusInternalServerError,
		}
	}
	return &usersModel.ResponseError{
		Message: "data has been deleted",
		Status:  http.StatusOK,
	}
}