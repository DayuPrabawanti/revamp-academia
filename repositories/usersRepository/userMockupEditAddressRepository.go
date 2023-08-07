package usersRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/usersRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type EditAddressRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewEditAddressRepository(dbHandler *sql.DB) *EditAddressRepository {
	return &EditAddressRepository{
		dbHandler: dbHandler,
	}
}

func (cr EditAddressRepository) AddAddress(ctx *gin.Context, masterAddressParams *dbContext.AddAddressParams, userAddressParams *dbContext.AddressParams, id int32) *models.ResponseError {

	store := dbContext.New(cr.dbHandler)

	err := store.AddAddress(ctx, *masterAddressParams)
	err = store.AddAddress2(ctx, *userAddressParams, id)

	if err != nil {
		return &models.ResponseError{
			Message: "error when Insert",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been Inserted",
		Status:  http.StatusOK,
	}


	// err := store.AddAddress(ctx, *masterAddressParams, id)

	// if err != nil {
	// 	return &models.ResponseError{
	// 		Message: "error when Insert",
	// 		Status:  http.StatusInternalServerError,
	// 	}
	// }
	// userAddress, err := store.AddAddress(ctx, *masterAddressParams, id)
	// if err != nil {
	// 	return nil, &models.ResponseError{
	// 		Message: err.Message,
	// 		Status:  http.StatusInternalServerError,
	// 	}
	// }

	// userAddr, err := store.AddAddress2(ctx, *userAddressParams, id)
	// if err != nil {
	// 	return nil, &models.ResponseError{
	// 		Message: err.Message,
	// 		Status:  http.StatusInternalServerError,
	// 	}
	// }

	// Address := &dbContext.Address{
	// 	Address: userAddress,
	// 	Addr: userAddr,
	// }

	// return Address, nil


}