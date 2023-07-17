package hrsCr

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"codeid.revamptwo/repositories/hrs/dbContext"
	"codeid.revamptwo/services/hrsSc"
	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	employeeService *hrsSc.EmployeeService
}

// declare constructor
func NewEmployeeController(employeeService *hrsSc.EmployeeService) *EmployeeController {
	return &EmployeeController{
		// struct 			parameter
		employeeService: employeeService,
	}
}

// LIST
func (employeeController EmployeeController) GetListEmployee(ctx *gin.Context) {
	responses, responseErr := employeeController.employeeService.GetListEmployee(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, responses)

	// ctx.JSON(http.StatusOK, "Hello gin framework")
}

// GET
func (employeeController EmployeeController) GetEmployee(ctx *gin.Context) {

	employeeId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := employeeController.employeeService.GetEmployee(ctx, int64(employeeId))
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// CREATE
func (employeeController EmployeeController) CreateEmployee(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var employee dbContext.CreateEmployeeParams
	err = json.Unmarshal(body, &employee)
	if err != nil {
		log.Println("Error while unmarshaling create category request body", err)
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

// UPDATE
func (employeeController EmployeeController) UpdateEmployee(ctx *gin.Context) {

	employeeId, err := strconv.Atoi(ctx.Param("id"))

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

	var employee dbContext.CreateEmployeeParams
	err = json.Unmarshal(body, &employee)
	if err != nil {
		log.Println("Error while unmarshaling update category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := employeeController.employeeService.UpdateEmployee(ctx, &employee, int64(employeeId))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

// DELETE
func (employeeController EmployeeController) DeleteEmployee(ctx *gin.Context) {

	employeeId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseErr := employeeController.employeeService.DeleteEmployee(ctx, int64(employeeId))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}
