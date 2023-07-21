package usersRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/usersRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type UserAddressRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewUserAddressRepository(dbHandler *sql.DB) *UserRepository {
	return &UserRepository{
		dbHandler: dbHandler,
	}
}

func (cr UserAddressRepository) GetListUserAddress(ctx *gin.Context) ([]*models.UsersUsersAddress, *models.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	usersAddr, err := store.ListAddress(ctx)

	listUsersAddress := make([]*models.UsersUsersAddress, 0)

	for _, v := range usersAddr {
		userAddr := &models.UsersUsersAddress{
			EtadAddrID: v.EtadAddrID,
			EtadModifiedDate: v.EtadModifiedDate,
			EtadEntityID: v.EtadEntityID,
			EtadAdtyID: v.EtadAdtyID,
		}
		listUsersAddress = append(listUsersAddress, userAddr)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listUsersAddress, nil
}