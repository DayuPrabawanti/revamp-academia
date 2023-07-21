package bootcampController

import (
	"net/http"

	"codeid.revampacademy/services/bootcampService"
	"github.com/gin-gonic/gin"
)

type ProgramApplyController struct {
	programApplyService *bootcampService.ProgramApplyService
}

// declare constructor
func NewProgramApplyController(programApplyService *bootcampService.ProgramApplyService) *ProgramApplyController {
	return &ProgramApplyController{
		programApplyService: programApplyService,
	}
}

// method
func (programApplyController ProgramApplyController) GetListProgramApply(ctx *gin.Context) {
	response, responseErr := programApplyController.programApplyService.GetListProgramApply(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}
