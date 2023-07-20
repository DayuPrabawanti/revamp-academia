package hrRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/hrRepository/dbContext"
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

func (dr DepartmentRepository) GetListDepartment(ctx *gin.Context) ([]*models.HrDepartment, *models.ResponseError) {

	store := dbContext.New(dr.dbHandler)
	departments, err := store.ListDepartment(ctx)

	listDepartment := make([]*models.HrDepartment, 0)

	for _, v := range departments {
		department := &models.HrDepartment{
			DeptID:           v.DeptID,
			DeptName:         v.DeptName,
			DeptModifiedDate: v.DeptModifiedDate,
		}
		listDepartment = append(listDepartment, department)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listDepartment, nil
}

func (dr DepartmentRepository) GetDepartment(ctx *gin.Context, id int64) (*models.HrDepartment, *models.ResponseError) {

	store := dbContext.New(dr.dbHandler)
	department, err := store.GetDepartment(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &department, nil
}

func (dr DepartmentRepository) CreateDepartment(ctx *gin.Context, departmentParams *dbContext.CreateDepartmentParams) (*models.HrDepartment, *models.ResponseError) {

	store := dbContext.New(dr.dbHandler)
	department, err := store.CreateDepartment(ctx, *departmentParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return department, nil
}

func (dr DepartmentRepository) UpdateDepartment(ctx *gin.Context, departmentParams *dbContext.CreateDepartmentParams) *models.ResponseError {

	store := dbContext.New(dr.dbHandler)
	err := store.UpdateDepartment(ctx, *departmentParams)

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

func (dr DepartmentRepository) DeleteDepartment(ctx *gin.Context, id int64) *models.ResponseError {

	store := dbContext.New(dr.dbHandler)
	err := store.DeleteDepartment(ctx, int32(id))

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
