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

type DepartmentController struct {
	departmentService *hrsSc.DepartmentService
}

// declare constructor
func NewDepartmentController(departmentService *hrsSc.DepartmentService) *DepartmentController {
	return &DepartmentController{
		// struct 			parameter
		departmentService: departmentService,
	}
}

// METHOD
// LIST
func (departmentController DepartmentController) GetListDepartment(ctx *gin.Context) {
	responses, responseErr := departmentController.departmentService.GetListDepartment(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, responses)

	// ctx.JSON(http.StatusOK, "Hello gin framework")
}

// GET
func (departmentController DepartmentController) GetDepartment(ctx *gin.Context) {

	departmentId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := departmentController.departmentService.GetDepartment(ctx, int64(departmentId))
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// CREATE
func (departmentController DepartmentController) CreateDepartment(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var department dbContext.CreateDepartmentParams
	err = json.Unmarshal(body, &department)
	if err != nil {
		log.Println("Error while unmarshaling create category request body", err)
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

// UPDATE
func (departmentController DepartmentController) UpdateDepartment(ctx *gin.Context) {

	departmentId, err := strconv.Atoi(ctx.Param("id"))

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

	var department dbContext.CreateDepartmentParams
	err = json.Unmarshal(body, &department)
	if err != nil {
		log.Println("Error while unmarshaling update category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := departmentController.departmentService.UpdateDepartment(ctx, &department, int64(departmentId))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

// DELETE
func (departmentController DepartmentController) DeleteDepartment(ctx *gin.Context) {

	categoryId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseErr := departmentController.departmentService.DeleteDepartment(ctx, int64(categoryId))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}
