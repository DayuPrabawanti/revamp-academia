package hrRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/hrRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type DepartmentHistoryRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewDepartmentHistoryRepository(dbHandler *sql.DB) *DepartmentHistoryRepository {
	return &DepartmentHistoryRepository{
		dbHandler: dbHandler,
	}
}

func (dhr DepartmentHistoryRepository) GetListDepartmentHistory(ctx *gin.Context) ([]*models.HrEmployeeDepartmentHistory, *models.ResponseError) {

	store := dbContext.New(dhr.dbHandler)
	departmenthistories, err := store.ListEmployeeDepartmentHistory(ctx)

	listDepartmentHistory := make([]*models.HrEmployeeDepartmentHistory, 0)

	for _, v := range departmenthistories {
		departmentHistory := &models.HrEmployeeDepartmentHistory{
			EdhiID:           v.EdhiID,
			EdhiEntityID:     v.EdhiEntityID,
			EdhiStartDate:    v.EdhiStartDate,
			EdhiEndDate:      v.EdhiEndDate,
			EdhiModifiedDate: v.EdhiModifiedDate,
			EdhiDeptID:       v.EdhiDeptID,
		}
		listDepartmentHistory = append(listDepartmentHistory, departmentHistory)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listDepartmentHistory, nil
}

func (dhr DepartmentHistoryRepository) GetDepartmentHistory(ctx *gin.Context, id int64) (*models.HrEmployeeDepartmentHistory, *models.ResponseError) {

	store := dbContext.New(dhr.dbHandler)
	departmentHistory, err := store.GetEmployeeDepartmentHistory(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &departmentHistory, nil
}

func (dhr DepartmentHistoryRepository) CreateDepartmentHistory(ctx *gin.Context, departmentHistoryParams *dbContext.CreateEmployeeDepartmentHistoryParams) (*models.HrEmployeeDepartmentHistory, *models.ResponseError) {

	store := dbContext.New(dhr.dbHandler)
	departmentHistory, err := store.CreateEmployeeDepartmentHistory(ctx, *departmentHistoryParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return departmentHistory, nil
}

func (dhr DepartmentHistoryRepository) UpdateDepartmentHistory(ctx *gin.Context, departmentHistoryParams *dbContext.UpdateEmployeeDepartmentHistoryParams) *models.ResponseError {

	store := dbContext.New(dhr.dbHandler)
	err := store.UpdateEmployeeDepartmentHistory(ctx, *departmentHistoryParams)

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

func (dhr DepartmentHistoryRepository) DeleteDepartmentHistory(ctx *gin.Context, id int64) *models.ResponseError {

	store := dbContext.New(dhr.dbHandler)
	err := store.DeleteEmployeeDepartmentHistory(ctx, int32(id))

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
