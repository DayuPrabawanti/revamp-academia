package hrs

import (
	"database/sql"
	"net/http"

	"codeid.revamptwo/models/hrsMdl"
	"codeid.revamptwo/repositories/hrs/dbContext"
	"github.com/gin-gonic/gin"
)

type EmployeeDepartmentHistoryRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewEmployeeDepartmentHistoryRepository(dbHandler *sql.DB) *EmployeeDepartmentHistoryRepository {
	return &EmployeeDepartmentHistoryRepository{
		dbHandler: dbHandler,
	}
}

// method GetListDepartment
func (cr EmployeeDepartmentHistoryRepository) ListEmployeeDepartmentHistory(ctx *gin.Context) ([]*hrsMdl.HrEmployeeDepartmentHistory, *hrsMdl.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	departmenthistory, err := store.ListEmployeeDepartmentHistory(ctx)

	listDepartmentHistory := make([]*hrsMdl.HrEmployeeDepartmentHistory, 0)

	for _, v := range departmenthistory {
		departments := &hrsMdl.HrEmployeeDepartmentHistory{
			EdhiID:           v.EdhiID,
			EdhiEntityID:     v.EdhiEntityID,
			EdhiStartDate:    v.EdhiStartDate,
			EdhiEndDate:      v.EdhiEndDate,
			EdhiModifiedDate: v.EdhiModifiedDate,
			EdhiDeptID:       v.EdhiDeptID,
		}
		listDepartmentHistory = append(listDepartmentHistory, departments)
	}

	if err != nil {
		return nil, &hrsMdl.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listDepartmentHistory, nil
}

// method GetDepartment
func (cr EmployeeDepartmentHistoryRepository) GetEmployeeDepartmentHistory(ctx *gin.Context, id int64) (*hrsMdl.HrEmployeeDepartmentHistory, *hrsMdl.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	departmenthistory, err := store.GetEmployeeDepartmentHistory(ctx, int32(id))

	if err != nil {
		return nil, &hrsMdl.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &departmenthistory, nil
}

// CreateDepartment
func (cr EmployeeDepartmentHistoryRepository) CreateEmployeeDepartmentHistory(ctx *gin.Context, departmentHistoryParams *dbContext.CreateEmployeeDepartmentHistoryParams) (*hrsMdl.HrEmployeeDepartmentHistory, *hrsMdl.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	departmenthistory, err := store.CreateEmployeeDepartmentHistory(ctx, *departmentHistoryParams)

	if err != nil {
		return nil, &hrsMdl.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return departmenthistory, nil
}

// UpdateDepartment
func (cr EmployeeDepartmentHistoryRepository) UpdateEmployeeDepartmentHistory(ctx *gin.Context, departmentHistoryParams *dbContext.UpdateEmployeeDepartmentHistoryParams) *hrsMdl.ResponseError {

	store := dbContext.New(cr.dbHandler)
	err := store.UpdateEmployeeDepartmentHistory(ctx, *departmentHistoryParams)

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
func (cr EmployeeDepartmentHistoryRepository) DeleteEmployeeDepartmentHistory(ctx *gin.Context, id int64) *hrsMdl.ResponseError {

	store := dbContext.New(cr.dbHandler)
	err := store.DeleteEmployeeDepartmentHistory(ctx, int32(id))

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
