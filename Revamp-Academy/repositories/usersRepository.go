package repositories

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/dbContext"
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

func (cr UserRepository) GetListUsers(ctx *gin.Context) ([]*models.UsersUser, *models.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	users, err := store.ListUsers(ctx)

	listUsers := make([]*models.UsersUser, 0)

	for _, v := range users {
		user := &models.UsersUser{
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
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listUsers, nil
}
