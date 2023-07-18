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

type EmployeeCLientContractController struct {
	employeeClientContractService *hrsSc.EmployeeClientContractService
}

// declare constructor
func NewEmployeeClientContractController(employeeClientContractService *hrsSc.EmployeeClientContractService) *EmployeeCLientContractController {
	return &EmployeeCLientContractController{
		// struct 			parameter
		employeeClientContractService: employeeClientContractService,
	}
}

// LIST
func (employeeClientContractController EmployeeCLientContractController) GetListEmployeeClientContract(ctx *gin.Context) {
	responses, responseErr := employeeClientContractController.employeeClientContractService.GetListEmployeeClientContract(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, responses)

	// ctx.JSON(http.StatusOK, "Hello gin framework")
}

// GET
func (employeeClientContractController EmployeeCLientContractController) GetEmployeeClientContract(ctx *gin.Context) {

	employeeClientConId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := employeeClientContractController.employeeClientContractService.GetEmployeeClientContract(ctx, int64(employeeClientConId))
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// CREATE
func (employeeClientContractController EmployeeCLientContractController) CreateEmployeeClientContract(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var employeeClientCon dbContext.CreateEmployeeClientContractParams
	err = json.Unmarshal(body, &employeeClientCon)
	if err != nil {
		log.Println("Error while unmarshaling create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := employeeClientContractController.employeeClientContractService.CreateEmployeeClientContract(ctx, &employeeClientCon)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

// UPDATE
func (employeeClientContractController EmployeeCLientContractController) UpdateClientContract(ctx *gin.Context) {

	clientContractId, err := strconv.Atoi(ctx.Param("id"))

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

	var clientcontract dbContext.CreateEmployeeClientContractParams
	err = json.Unmarshal(body, &clientcontract)
	if err != nil {
		log.Println("Error while unmarshaling update category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := employeeClientContractController.employeeClientContractService.UpdateClientContract(ctx, &clientcontract, int64(clientContractId))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

// DELETE
func (employeeClientContractController EmployeeCLientContractController) DeleteClientContract(ctx *gin.Context) {

	clientcontractId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseErr := employeeClientContractController.employeeClientContractService.DeleteClientContract(ctx, int64(clientcontractId))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}
