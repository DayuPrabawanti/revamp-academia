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

type DepartmentHistoryController struct {
	departmentHistoryService *hrService.DepartmentHistoryService
}

// declare constructor
func NewDepartmentHistoryController(departmentHistoryService *hrService.DepartmentHistoryService) *DepartmentHistoryController {
	return &DepartmentHistoryController{
		departmentHistoryService: departmentHistoryService,
	}
}

// create method
func (departmentHistoryController DepartmentHistoryController) GetListDepartmentHistory(ctx *gin.Context) {

	response, responseErr := departmentHistoryController.departmentHistoryService.GetListDepartmentHistory(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
	// ctx.JSON(http.StatusOK, "Hello gin framework")
}

func (departmentHistoryController DepartmentHistoryController) GetDepartmentHistory(ctx *gin.Context) {

	edhi_id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := departmentHistoryController.departmentHistoryService.GetDepartmentHistory(ctx, int64(edhi_id))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (departmentHistoryController DepartmentHistoryController) CreateDepartmentHistory(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create department history request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var departmentHistory dbContext.CreateEmployeeDepartmentHistoryParams
	err = json.Unmarshal(body, &departmentHistory)
	if err != nil {
		log.Println("Error while unmarshaling create department history request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := departmentHistoryController.departmentHistoryService.CreateDepartmentHistory(ctx, &departmentHistory)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (departmentHistoryController DepartmentHistoryController) UpdateDepartmentHistory(ctx *gin.Context) {

	edhi_id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update department history request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var departmentHistory dbContext.UpdateEmployeeDepartmentHistoryParams
	err = json.Unmarshal(body, &departmentHistory)
	if err != nil {
		log.Println("Error while unmarshaling update department history request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := departmentHistoryController.departmentHistoryService.UpdateDepartmentHistory(ctx, &departmentHistory, int64(edhi_id))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (departmentHistoryController DepartmentHistoryController) DeleteDepartmenHistory(ctx *gin.Context) {

	edhi_id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseErr := departmentHistoryController.departmentHistoryService.DeleteDepartmentHistory(ctx, int64(edhi_id))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}
