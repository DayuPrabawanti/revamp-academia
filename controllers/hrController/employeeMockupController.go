package hrController

import (
	"net/http"

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
