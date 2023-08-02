package hrRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/hrRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type ClientContractRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewClientContractRepository(dbHandler *sql.DB) *ClientContractRepository {
	return &ClientContractRepository{
		dbHandler: dbHandler,
	}
}

func (ccr ClientContractRepository) GetListClientContract(ctx *gin.Context) ([]*models.HrEmployeeClientContract, *models.ResponseError) {

	store := dbContext.New(ccr.dbHandler)
	clientContracts, err := store.ListClientContract(ctx)

	listClientContract := make([]*models.HrEmployeeClientContract, 0)

	for _, v := range clientContracts {
		clientContract := &models.HrEmployeeClientContract{
			EccoID:         v.EccoID,
			EccoEntityID:   v.EccoEntityID,
			EccoContractNo: v.EccoContractNo,
			EccoStartDate:  v.EccoStartDate,
			EccoEndDate:    v.EccoEndDate,
			EccoNotes:      v.EccoNotes,
		}
		listClientContract = append(listClientContract, clientContract)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listClientContract, nil
}

func (ccr ClientContractRepository) GetClientContract(ctx *gin.Context, id int64) (*models.HrEmployeeClientContract, *models.ResponseError) {

	store := dbContext.New(ccr.dbHandler)
	clientContract, err := store.GetClientContract(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &clientContract, nil
}

func (ccr ClientContractRepository) CreateClientContract(ctx *gin.Context, clientContractParams *dbContext.CreateClientContractParams) (*models.HrEmployeeClientContract, *models.ResponseError) {

	store := dbContext.New(ccr.dbHandler)
	clientContract, err := store.CreateClientContract(ctx, *clientContractParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return clientContract, nil
}

func (ccr ClientContractRepository) UpdateClientContract(ctx *gin.Context, clientContractParams *dbContext.CreateClientContractParams) *models.ResponseError {

	store := dbContext.New(ccr.dbHandler)
	err := store.UpdateClientContract(ctx, *clientContractParams)

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

func (ccr ClientContractRepository) DeleteClientContract(ctx *gin.Context, id int64) *models.ResponseError {

	store := dbContext.New(ccr.dbHandler)
	err := store.DeleteClientContract(ctx, int32(id))

	if err != nil {
		return &models.ResponseError{
			Message: "error when delete",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been deleted",
		Status:  http.StatusOK,
	}
}
