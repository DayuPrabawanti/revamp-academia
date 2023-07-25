package hrRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/hrRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type EmployeeMockupRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewEmployeeMockupRepository(dbHandler *sql.DB) *EmployeeMockupRepository {
	return &EmployeeMockupRepository{
		dbHandler: dbHandler,
	}
}

func (cr EmployeeMockupRepository) CreateEmployeeMockup(ctx *gin.Context, employeemockupParams *dbContext.EmployeeMockupParams) (*models.EmployeeMockupModel, *models.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	general, err := store.CreateEmployee(ctx, employeemockupParams.General)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}

	salary, err := store.CreatePayHistory(ctx, employeemockupParams.Salary)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}

	assigment, err := store.CreateEmployeeDepartmentHistory(ctx, employeemockupParams.Assigment)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	employeeMockup := &models.EmployeeMockupModel{
		General:    *general,
		Department: *salary,
		Assigment:  *assigment,
	}

	return employeeMockup, nil
}
