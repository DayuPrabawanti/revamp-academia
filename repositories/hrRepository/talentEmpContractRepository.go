package hrRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/hrRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type TalentClientContractRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewTalentClientContractRepository(dbHandler *sql.DB) *TalentClientContractRepository {
	return &TalentClientContractRepository{
		dbHandler: dbHandler,
	}
}

func (ccr TalentClientContractRepository) GetTalentClientContract(ctx *gin.Context, id int64) (*models.TalentClientContractGetUpdate, *models.ResponseError) {

	store := dbContext.New(ccr.dbHandler)
	clientContract, err := store.GetTalentClientContract(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &clientContract, nil
}

func (ccr TalentClientContractRepository) UpdateTalentClientContract(ctx *gin.Context, clientContractParams *dbContext.UpdateTalentClientContractParams) *models.ResponseError {

	store := dbContext.New(ccr.dbHandler)
	err := store.UpdateTalentClientContract(ctx, *clientContractParams)

	if err != nil {
		return &models.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been update",
		Status:  http.StatusOK,
	}
}
