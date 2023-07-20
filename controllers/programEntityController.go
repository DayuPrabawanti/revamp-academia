package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"codeid.revampacademy/repositories/dbcontext"
	"codeid.revampacademy/services"
	"github.com/gin-gonic/gin"
)

type ProgramEntityController struct {
	programEntityService *services.ProgramEntityService
}

func NewProgramEntityController(programEntityService *services.ProgramEntityService) *ProgramEntityController {
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

func (programEntityController ProgramEntityController) GetListMasterCategory(ctx *gin.Context) {
	response, responerr := programEntityController.programEntityService.GetListMasterCategory(ctx)
	if responerr != nil {
		ctx.JSON(responerr.Status, responerr)
	}
	ctx.JSON(http.StatusOK, response)
}

func (ProgramEntityController ProgramEntityController) Group(ctx *gin.Context) {
	response, responseErr := ProgramEntityController.programEntityService.Group(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (programEntityController ProgramEntityController) GetProgramEntity(ctx *gin.Context) {

	prog_entity_id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := programEntityController.programEntityService.GetProgramEntity(ctx, int64(prog_entity_id))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (programEntityController ProgramEntityController) CreateProgramEntity(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create program entity request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var program_entity dbcontext.CreateProgramEntityParams
	err = json.Unmarshal(body, &program_entity)
	if err != nil {
		log.Println("Error while unmarshaling create program entity request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := programEntityController.programEntityService.CreateProgramEntity(ctx, &program_entity)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (programEntityController ProgramEntityController) CreateGroup(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create Gabung request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var groupParams dbcontext.CreateGroup
	err = json.Unmarshal(body, &groupParams)
	if err != nil {
		log.Println("Error while unmarshaling create Gabung request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := programEntityController.programEntityService.CreateGroup(ctx, &groupParams)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
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

	var programEntity dbcontext.CreateProgramEntityParams
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

func (programEntityController ProgramEntityController) DeleteProgramEntity(ctx *gin.Context) {

	progEntityID, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseErr := programEntityController.programEntityService.DeleteProgramEntity(ctx, int64(progEntityID))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}
