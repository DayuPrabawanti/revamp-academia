package repositories

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/dbContext"
	"github.com/gin-gonic/gin"
)

type MasterIndustryRepository struct {
	dbHandler   *sql.DB
	Transaction *sql.Tx
}

func NewMasterIndustryRepository(dbHandler *sql.DB) *MasterIndustryRepository {
	return &MasterIndustryRepository{
		dbHandler: dbHandler,
	}
}

func (mir MasterIndustryRepository) GetMasterIndustryRepo(ctx *gin.Context, InduCodeID int32) (*models.MasterIndustry, *models.ResponseError) {

	store2 := dbContext.New(mir.dbHandler)
	masterIndustry, err := store2.GetMasterIndustryImpl(ctx, int32(InduCodeID))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &masterIndustry, nil
}

func (mir MasterIndustryRepository) ListMasterIndustryRepo(ctx *gin.Context) ([]*models.MasterIndustry, *models.ResponseError) {

	store := dbContext.New(mir.dbHandler)
	masterIndustry, err := store.ListMasterIndustryImpl(ctx)

	ListMasterIndustry := make([]*models.MasterIndustry, 0)

	for _, v := range masterIndustry {
		masterIndus := &models.MasterIndustry{
			InduCodeID: v.InduCodeID,
			InduName:   v.InduName,
		}
		ListMasterIndustry = append(ListMasterIndustry, masterIndus)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return ListMasterIndustry, nil
}
