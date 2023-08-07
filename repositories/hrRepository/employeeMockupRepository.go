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
	dbQueries   dbContext.Queries
}

func NewEmployeeMockupRepository(dbHandler *sql.DB) *EmployeeMockupRepository {
	return &EmployeeMockupRepository{
		dbHandler: dbHandler,
		dbQueries: *dbContext.New(dbHandler),
	}
}

func (er EmployeeMockupRepository) ListEmployeeMockup(ctx *gin.Context) ([]*models.EmployeeMockupList, *models.ResponseError) {

	store := dbContext.New(er.dbHandler)
	employees, err := store.ListEmployeeMockup(ctx)

	listEmployees := make([]*models.EmployeeMockupList, 0)

	for _, v := range employees {
		employee := &models.EmployeeMockupList{
			HrEmployee: v.HrEmployee,
			UsersUser:  v.UsersUser,
		}
		listEmployees = append(listEmployees, employee)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listEmployees, nil
}

func (cr EmployeeMockupRepository) CreateEmployeeMockup(ctx *gin.Context, userParam *dbContext.CreateUsersParams, empParam *dbContext.CreateEmployeeParams, salaryParam *dbContext.CreatePayHistoryParams, assigmentParam *dbContext.CreateEmployeeDepartmentHistoryParams, clientParam *dbContext.CreateClientContractParams) (*models.EmployeeMockupModel, *models.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	user, err := store.CreateUsers(ctx, *userParam)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}

	general, err := store.CreateEmployee(ctx, *empParam)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}

	salary, err := store.CreatePayHistory(ctx, *salaryParam)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}

	assigment, err := store.CreateEmployeeDepartmentHistory(ctx, *assigmentParam)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}

	client, err := store.CreateClientContract(ctx, *clientParam)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}

	employeeMockup := &models.EmployeeMockupModel{
		User:      *user,
		General:   *general,
		Salary:    *salary,
		Assigment: *assigment,
		Client:    *client,
	}

	return employeeMockup, nil
}

func (tdmr EmployeeMockupRepository) SearchEmployee(ctx *gin.Context, userName string, status string) ([]models.EmployeeMockupList, *models.ResponseError) {
	store := dbContext.New(tdmr.dbHandler)
	employee, err := store.SearchEmployee(ctx, userName, status)
	if err != nil {
		return nil, &models.ResponseError{
			Message: "Failed to search employee",
			Status:  http.StatusInternalServerError,
		}
	}

	return employee, nil
}

func (cr EmployeeMockupRepository) UpdateEmployeeMockup(ctx *gin.Context, usersParam *dbContext.UpdateUsersParams, empParam *dbContext.UpdateEmployeeParams, salaryParam *dbContext.UpdatePayHistoryParams, assigmentParam *dbContext.UpdateEmployeeDepartmentHistoryParams, id int64) *models.ResponseError {

	store := dbContext.New(cr.dbHandler)
	err := store.UpdateUser(ctx, *usersParam)
	if err != nil {
		return &models.ResponseError{
			Message: "error when update user",
			Status:  http.StatusInternalServerError,
		}
	}

	store2 := dbContext.New(cr.dbHandler)
	err2 := store2.UpdateEmployee(ctx, *empParam)
	if err2 != nil {
		return &models.ResponseError{
			Message: "error when update employee",
			Status:  http.StatusInternalServerError,
		}
	}

	store3 := dbContext.New(cr.dbHandler)
	err3 := store3.UpdatePayHistory(ctx, *salaryParam)
	if err3 != nil {
		return &models.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}

	store4 := dbContext.New(cr.dbHandler)
	err4 := store4.UpdateEmployeeDepartmentHistory(ctx, *assigmentParam)
	if err4 != nil {
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
