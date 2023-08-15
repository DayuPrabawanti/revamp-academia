package bootcampController

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"codeid.revampacademy/repositories/bootcampRepository/dbContext"
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
func (programEntityController ProgramEntityController) UpdateProgramEntity(ctx *gin.Context) {

	progEntityID, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update program entity request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var programEntity dbContext.CurriculumProgramEntityParams
	err = json.Unmarshal(body, &programEntity)
	if err != nil {
		log.Println("Error while unmarshaling update program entity request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := programEntityController.programEntityService.UpdateProgramEntity(ctx, &programEntity, int64(progEntityID))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (programEntityController ProgramEntityController) GetProgEntity(ctx *gin.Context) {

	progEntityId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := programEntityController.programEntityService.GetProgEntity(ctx, int64(progEntityId))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
