package masterRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/masterRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type MasterAddressRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewMasteraddressRepository(dbHandler *sql.DB) *MasterAddressRepository {
	return &MasterAddressRepository{
		dbHandler: dbHandler,
	}
}

func (cr MasterAddressRepository) GetListMasterAddress(ctx *gin.Context) ([]*models.MasterAddress, *models.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	masterAddr, err := store.ListMasterAddress(ctx)

	listMasterAddress := make([]*models.MasterAddress, 0)

	for _, v := range masterAddr {
		masterAddress := &models.MasterAddress{
			AddrID: v.AddrID,
			AddrLine1: v.AddrLine1,
			AddrLine2: v.AddrLine2,
			AddrPostalCode: v.AddrPostalCode,
			AddrSpatialLocation: v.AddrSpatialLocation,
			AddrModifiedDate: v.AddrModifiedDate,
			AddrCityID: v.AddrCityID,
		}
		listMasterAddress = append(listMasterAddress, masterAddress)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listMasterAddress, nil
}