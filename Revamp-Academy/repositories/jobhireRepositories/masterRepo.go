package jobhireRepositories

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/jobhireRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type MasterRepo struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
	dbQueries   dbContext.Queries
}

func NewMasterRepo(dbHandler *sql.DB) *MasterRepo {
	return &MasterRepo{
		dbHandler: dbHandler,
		dbQueries: *dbContext.New(dbHandler),
	}
}

func (mr MasterRepo) GetListMasterAddress(ctx *gin.Context) ([]*models.MasterAddress, *models.ResponseError) {
	market := dbContext.New(mr.dbHandler)
	address, err := market.ListMasterAddress(ctx)

	listAddress := make([]*models.MasterAddress, 0)

	for _, v := range address {
		master := models.MasterAddress{
			AddrID:              v.AddrID,
			AddrLine1:           v.AddrLine1,
			AddrLine2:           v.AddrLine2,
			AddrPostalCode:      v.AddrPostalCode,
			AddrSpatialLocation: v.AddrSpatialLocation,
			AddrModifiedDate:    v.AddrModifiedDate,
			AddrCityID:          v.AddrCityID,
		}
		listAddress = append(listAddress, &master)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listAddress, nil
}

func (mr MasterRepo) GetListMasterCity(ctx *gin.Context) ([]*models.MasterCity, *models.ResponseError) {
	market := dbContext.New(mr.dbHandler)
	city, err := market.ListMasterCity(ctx)

	listCity := make([]*models.MasterCity, 0)

	for _, v := range city {
		master := models.MasterCity{
			CityID:           v.CityID,
			CityName:         v.CityName,
			CityModifiedDate: v.CityModifiedDate,
			CityProvID:       v.CityProvID,
		}
		listCity = append(listCity, &master)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listCity, nil
}

// Repo for create master.address
func (mr MasterRepo) CreateMasterAddressRepo(ctx *gin.Context, addressParams *dbContext.CreateMasterAddressParams) (*models.MasterAddress, *models.ResponseError) {
	market := dbContext.New(mr.dbHandler)
	address, err := market.CreateMasterAddress(ctx, *addressParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return address, nil
}
