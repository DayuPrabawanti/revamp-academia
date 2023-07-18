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

type EmployeeDepartmentHistoryController struct {
	departmentHistoryService *hrsSc.EmployeeDepartmentHistoryService
}

// declare constructor
func NewEmployeeDepartmentHistoryController(departmentHistoryService *hrsSc.EmployeeDepartmentHistoryService) *EmployeeDepartmentHistoryController {
	return &EmployeeDepartmentHistoryController{
		// struct 			parameter
		departmentHistoryService: departmentHistoryService,
	}
}

// METHOD
// LIST
func (departmentHistoryController EmployeeDepartmentHistoryController) ListEmployeeDepartmentHistory(ctx *gin.Context) {
	responses, responseErr := departmentHistoryController.departmentHistoryService.ListEmployeeDepartmentHistory(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, responses)

	// ctx.JSON(http.StatusOK, "Hello gin framework")
}

// GET
func (departmentHistoryController EmployeeDepartmentHistoryController) GetEmployeeDepartmentHistory(ctx *gin.Context) {

	departmentHistoryId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := departmentHistoryController.departmentHistoryService.GetEmployeeDepartmentHistory(ctx, int64(departmentHistoryId))
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// CREATE
func (departmentHistoryController EmployeeDepartmentHistoryController) CreateEmployeeDepartmentHistory(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var departmentHistory dbContext.CreateEmployeeDepartmentHistoryParams
	err = json.Unmarshal(body, &departmentHistory)
	if err != nil {
		log.Println("Error while unmarshaling create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := departmentHistoryController.departmentHistoryService.CreateEmployeeDepartmentHistory(ctx, &departmentHistory)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

// UPDATE
func (departmentHistoryController EmployeeDepartmentHistoryController) UpdateEmployeeDepartmentHistory(ctx *gin.Context) {

	departmentHistoryId, err := strconv.Atoi(ctx.Param("id"))

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

	var departments dbContext.CreateEmployeeDepartmentHistoryParams
	err = json.Unmarshal(body, &departments)
	if err != nil {
		log.Println("Error while unmarshaling update category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := departmentHistoryController.departmentHistoryService.UpdateEmployeeDepartmentHistory(ctx, &departments, int64(departmentHistoryId))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

// DELETE
func (departmentHistoryController EmployeeDepartmentHistoryController) DeleteEmployeeDepartmentHistory(ctx *gin.Context) {

	departmentHistoryId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseErr := departmentHistoryController.departmentHistoryService.DeleteEmployeeDepartmentHistory(ctx, int64(departmentHistoryId))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}
