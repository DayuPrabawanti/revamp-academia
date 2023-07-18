package hrs

import (
	"database/sql"
	"net/http"

	"codeid.revamptwo/models/hrsMdl"
	"codeid.revamptwo/repositories/hrs/dbContext"
	"github.com/gin-gonic/gin"
)

type EmployeeClientContractRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewEmployeeCLientContractRepository(dbHandler *sql.DB) *EmployeeClientContractRepository {
	return &EmployeeClientContractRepository{
		dbHandler: dbHandler,
	}
}

func (cr EmployeeClientContractRepository) GetListEmployeeClientContract(ctx *gin.Context) ([]*hrsMdl.HrEmployeeClientContract, *hrsMdl.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	empClientContract, err := store.GetListEmployeeClientContract(ctx)

	listClientCon := make([]*hrsMdl.HrEmployeeClientContract, 0)

	for _, v := range empClientContract {
		empclientcontract := &hrsMdl.HrEmployeeClientContract{
			EccoID:             v.EccoID,
			EccoEntityID:       v.EccoEntityID,
			EccoContractNo:     v.EccoContractNo,
			EccoContractDate:   v.EccoContractDate,
			EccoStartDate:      v.EccoStartDate,
			EccoEndDate:        v.EccoEndDate,
			EccoNotes:          v.EccoNotes,
			EccoModifiedDate:   v.EccoModifiedDate,
			EccoMediaLink:      v.EccoMediaLink,
			EccoJotyID:         v.EccoJotyID,
			EccoAccountManager: v.EccoAccountManager,
			EccoClitID:         v.EccoClitID,
			EccoStatus:         v.EccoStatus,
		}
		listClientCon = append(listClientCon, empclientcontract)
	}

	if err != nil {
		return nil, &hrsMdl.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listClientCon, nil
}

func (cr EmployeeClientContractRepository) GetEmployeeClientContract(ctx *gin.Context, id int64) (*hrsMdl.HrEmployeeClientContract, *hrsMdl.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	employeecliencon, err := store.GetEmployeeClientContract(ctx, int32(id))

	if err != nil {
		return nil, &hrsMdl.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &employeecliencon, nil
}

func (cr EmployeeClientContractRepository) CreateEmployeeClientContract(ctx *gin.Context, employeeClientContractParams *dbContext.CreateEmployeeClientContractParams) (*hrsMdl.HrEmployeeClientContract, *hrsMdl.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	employeeClientCon, err := store.CreateEmployeeClientContract(ctx, *employeeClientContractParams)

	if err != nil {
		return nil, &hrsMdl.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return employeeClientCon, nil
}

// UpdateDepartment
func (cr EmployeeClientContractRepository) UpdateClientContract(ctx *gin.Context, employeeClientContractParams *dbContext.CreateEmployeeClientContractParams) *hrsMdl.ResponseError {

	store := dbContext.New(cr.dbHandler)
	err := store.UpdateEmployeeClientContract(ctx, *employeeClientContractParams)

	if err != nil {
		return &hrsMdl.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &hrsMdl.ResponseError{
		Message: "data has been update",
		Status:  http.StatusOK,
	}
}

// method DeleteCategory
func (cr EmployeeClientContractRepository) DeleteClientContract(ctx *gin.Context, id int64) *hrsMdl.ResponseError {

	store := dbContext.New(cr.dbHandler)
	err := store.DeleteEmployeeClientContract(ctx, int32(id))

	if err != nil {
		return &hrsMdl.ResponseError{
			Message: "error when deleted",
			Status:  http.StatusInternalServerError,
		}
	}
	return &hrsMdl.ResponseError{
		Message: "data has been deleted",
		Status:  http.StatusOK,
	}
}
