package bootcampController

import (
	"net/http"

	"codeid.revampacademy/services/bootcampService"
	"github.com/gin-gonic/gin"
)

type InstructorProgramController struct {
	instructorProgramService *bootcampService.InstructorProgramService
}

// declare constructor
func NewInstructorProgramController(instructorProgramService *bootcampService.InstructorProgramService) *InstructorProgramController {
	return &InstructorProgramController{
		instructorProgramService: instructorProgramService,
	}
}

// method
func (instructorProgramController InstructorProgramController) GetListInstructorProgram(ctx *gin.Context) {
	response, responseErr := instructorProgramController.instructorProgramService.GetListInstructorProgram(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}
