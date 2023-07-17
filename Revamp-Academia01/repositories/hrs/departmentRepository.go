package hrs

import (
	"database/sql"
	"net/http"

	"codeid.revamptwo/models/hrsMdl"
	"codeid.revamptwo/repositories/hrs/dbContext"
	"github.com/gin-gonic/gin"
)

type DepartmentRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewDepartmentRepository(dbHandler *sql.DB) *DepartmentRepository {
	return &DepartmentRepository{
		dbHandler: dbHandler,
	}
}

// method GetListDepartment
func (cr DepartmentRepository) GetListDepartment(ctx *gin.Context) ([]*hrsMdl.HrDepartment, *hrsMdl.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	departments, err := store.ListDepartment(ctx)

	listDepartment := make([]*hrsMdl.HrDepartment, 0)

	for _, v := range departments {
		department := &hrsMdl.HrDepartment{
			DeptID:           v.DeptID,
			DeptName:         v.DeptName,
			DeptModifiedDate: v.DeptModifiedDate,
		}
		listDepartment = append(listDepartment, department)
	}

	if err != nil {
		return nil, &hrsMdl.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listDepartment, nil
}

// method GetDepartment
func (cr DepartmentRepository) GetDepartment(ctx *gin.Context, id int64) (*hrsMdl.HrDepartment, *hrsMdl.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	department, err := store.GetDepartment(ctx, int32(id))

	if err != nil {
		return nil, &hrsMdl.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &department, nil
}

// CreateDepartment
func (cr DepartmentRepository) CreateDepartment(ctx *gin.Context, departmentParams *dbContext.CreateDepartmentParams) (*hrsMdl.HrDepartment, *hrsMdl.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	department, err := store.CreateDepartment(ctx, *departmentParams)

	if err != nil {
		return nil, &hrsMdl.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return department, nil
}

// UpdateDepartment
func (cr DepartmentRepository) UpdateDepartment(ctx *gin.Context, departmentParams *dbContext.CreateDepartmentParams) *hrsMdl.ResponseError {

	store := dbContext.New(cr.dbHandler)
	err := store.UpdateDepartment(ctx, *departmentParams)

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
func (cr DepartmentRepository) DeleteDepartment(ctx *gin.Context, id int64) *hrsMdl.ResponseError {

	store := dbContext.New(cr.dbHandler)
	err := store.DeleteDepartment(ctx, int32(id))

	if err != nil {
		return &hrsMdl.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &hrsMdl.ResponseError{
		Message: "data has been deleted",
		Status:  http.StatusOK,
	}
}
