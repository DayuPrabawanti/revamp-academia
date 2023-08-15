package bootcampController

import (
	"net/http"

	"codeid.revampacademy/services/bootcampService"
	"github.com/gin-gonic/gin"
)

type ProgramEntityController struct {
	programEntityService *bootcampService.ProgramEntityService
}

func NewProgramEntityController(programEntityService *bootcampService.ProgramEntityService) *ProgramEntityController {
	return &ProgramEntityController{
		programEntityService: programEntityService,
	}
}

func (programEntityController ProgramEntityController) GetListProgramEntity(ctx *gin.Context) {
	response, responseErr := programEntityController.programEntityService.GetListProgramEntity(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
