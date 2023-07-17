package controllers

import (
	"net/http"

	"codeid.revampacademy/services"
	"github.com/gin-gonic/gin"
)

type ProgramApplyController struct {
	programApplyService *services.ProgramApplyService
}

// declare constructor
func NewProgramApplyController(programApplyService *services.ProgramApplyService) *ProgramApplyController {
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
