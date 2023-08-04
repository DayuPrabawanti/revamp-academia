package hrController

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"codeid.revampacademy/repositories/hrRepository/dbContext"
	"codeid.revampacademy/services/hrService"
	"github.com/gin-gonic/gin"
)

type EmployeeMockupController struct {
	employeeMockupService *hrService.EmployeeMockupService
}

// declare constructor
func NewEmployeeMockupController(employeeMockupService *hrService.EmployeeMockupService) *EmployeeMockupController {
	return &EmployeeMockupController{
		employeeMockupService: employeeMockupService,
	}
}

func (employeeMockupController EmployeeMockupController) ListEmployeeMockup(ctx *gin.Context) {

	response, responseErr := employeeMockupController.employeeMockupService.ListEmployeeMockup(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (employeeMockupController EmployeeMockupController) SearchEmployee(ctx *gin.Context) {
	userName := ctx.DefaultQuery("name", "")
	status := ctx.DefaultQuery("status", "")

	employee, responseErr := employeeMockupController.employeeMockupService.SearchEmployee(ctx, userName, status)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, employee)
}

func (empMockupController EmployeeMockupController) CreateEmployeeMockup(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create employee request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var user dbContext.CreateUsersParams
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("Error while unmarshaling User request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var general dbContext.CreateEmployeeParams
	err = json.Unmarshal(body, &general)
	if err != nil {
		log.Println("Error while unmarshaling Employee request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var salary dbContext.CreatePayHistoryParams
	err = json.Unmarshal(body, &salary)
	if err != nil {
		log.Println("Error while unmarshaling Salary request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var assigment dbContext.CreateEmployeeDepartmentHistoryParams
	err = json.Unmarshal(body, &assigment)
	if err != nil {
		log.Println("Error while unmarshaling Department request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := empMockupController.employeeMockupService.EmployeeMockup(ctx, &user, &general, &salary, &assigment)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}
