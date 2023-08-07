package hrController

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"codeid.revampacademy/repositories/hrRepository/dbContext"
	"codeid.revampacademy/services/hrService"
	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	employeeService *hrService.EmployeeService
}

// declare constructor
func NewEmployeeController(employeeService *hrService.EmployeeService) *EmployeeController {
	return &EmployeeController{
		employeeService: employeeService,
	}
}

// create method
func (employeeController EmployeeController) GetListEmployee(ctx *gin.Context) {

	response, responseErr := employeeController.employeeService.GetListEmployee(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
	// ctx.JSON(http.StatusOK, "Hello gin framework")
}

func (employeeController EmployeeController) GetEmployee(ctx *gin.Context) {

	emp_entity_id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := employeeController.employeeService.GetEmployee(ctx, int64(emp_entity_id))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (employeeController EmployeeController) CreateEmployee(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create employee request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var employee dbContext.CreateEmployeeParams
	err = json.Unmarshal(body, &employee)
	if err != nil {
		log.Println("Error while unmarshaling create department request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := employeeController.employeeService.CreateEmployee(ctx, &employee)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (employeeController EmployeeController) UpdateEmployee(ctx *gin.Context) {

	emp_entity_id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update employee request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var employee dbContext.UpdateEmployeeParams
	err = json.Unmarshal(body, &employee)
	if err != nil {
		log.Println("Error while unmarshaling update employee request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := employeeController.employeeService.UpdateEmployee(ctx, &employee, int64(emp_entity_id))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (employeeController EmployeeController) DeleteEmployee(ctx *gin.Context) {

	emp_entity_id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseErr := employeeController.employeeService.DeleteEmployee(ctx, int64(emp_entity_id))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}

// USER

func (employeeController EmployeeController) CreateUser(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var user dbContext.CreateUsersParams
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("Error while unmarshaling create User request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := employeeController.employeeService.CreateUser(ctx, &user)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (employeeController EmployeeController) UpdateUser(ctx *gin.Context) {

	userId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var user dbContext.UpdateUsersParams
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("Error while unmarshaling update category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := employeeController.employeeService.UpdateUser(ctx, &user, int64(userId))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}
