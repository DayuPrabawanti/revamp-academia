package controllers

import (
	"net/http"

	"codeid.revampacademy/services"
	"github.com/gin-gonic/gin"
)

type ProgramApplyProgressController struct {
	programApplyProgressService *services.ProgramApplyProgressService
}

// declare constructor
func NewProgramApplyProgressController(programApplyProgressService *services.ProgramApplyProgressService) *ProgramApplyProgressController {
	return &ProgramApplyProgressController{
		programApplyProgressService: programApplyProgressService,
	}
}

// method
func (programApplyProgressController ProgramApplyProgressController) GetListProgramApplyProgress(ctx *gin.Context) {
	response, responseErr := programApplyProgressController.programApplyProgressService.GetListProgramApplyProgress(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}
