package bootcampRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/bootcampRepository/dbContext"
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

func (cr UserPhoneRepository) GetListUsersPhone(ctx *gin.Context) ([]*models.UsersUsersPhone, *models.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	usersPhone, err := store.ListPhones(ctx)

	listUsersPhone := make([]*models.UsersUsersPhone, 0)

	for _, v := range usersPhone {
		userPhone := &models.UsersUsersPhone{
			UspoEntityID:     v.UspoEntityID,
			UspoNumber:       v.UspoNumber,
			UspoModifiedDate: v.UspoModifiedDate,
			UspoPontyCode:    v.UspoPontyCode,
		}
		listUsersPhone = append(listUsersPhone, userPhone)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listUsersPhone, nil
}
