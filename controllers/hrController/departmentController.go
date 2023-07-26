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

type DepartmentController struct {
	departmentService *hrService.DepartmentService
}

// declare constructor
func NewDepartmentController(departmentService *hrService.DepartmentService) *DepartmentController {
	return &DepartmentController{
		departmentService: departmentService,
	}
}

// create method
func (departmentController DepartmentController) GetListDepartment(ctx *gin.Context) {

	response, responseErr := departmentController.departmentService.GetListDepartment(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
	// ctx.JSON(http.StatusOK, "Hello gin framework")
}

func (departmentController DepartmentController) GetDepartment(ctx *gin.Context) {

	dept_id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := departmentController.departmentService.GetDepartment(ctx, int64(dept_id))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (departmentController DepartmentController) CreateDepartment(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create department request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var department dbContext.CreateDepartmentParams
	err = json.Unmarshal(body, &department)
	if err != nil {
		log.Println("Error while unmarshaling create department request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := departmentController.departmentService.CreateDepartment(ctx, &department)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (departmentController DepartmentController) UpdateDepartment(ctx *gin.Context) {

	dept_id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update department request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var department dbContext.CreateDepartmentParams
	err = json.Unmarshal(body, &department)
	if err != nil {
		log.Println("Error while unmarshaling update department request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := departmentController.departmentService.UpdateDepartment(ctx, &department, int64(dept_id))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (departmentController DepartmentController) DeleteDepartment(ctx *gin.Context) {

	dept_id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseErr := departmentController.departmentService.DeleteDepartment(ctx, int64(dept_id))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}
