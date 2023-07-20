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

type SectionController struct {
	sectionService *services.SectionService
}

func NewSectionController(sectionService *services.SectionService) *SectionController {
	return &SectionController{
		sectionService: sectionService,
	}
}

func (sectionController SectionController) GetListSection(ctx *gin.Context) {
	response, responseErr := sectionController.sectionService.GetListSection(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (sectionController SectionController) GetSections(ctx *gin.Context) {

	sectId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := sectionController.sectionService.GetSections(ctx, int64(sectId))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (sectionController SectionController) Createsections(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create sections request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var section dbcontext.CreatesectionsParams
	err = json.Unmarshal(body, &section)
	if err != nil {
		log.Println("Error while unmarshaling create sections request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := sectionController.sectionService.Createsections(ctx, &section)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (sectionController SectionController) UpdateSections(ctx *gin.Context) {

	sectId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update sections request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var section dbcontext.CreatesectionsParams
	err = json.Unmarshal(body, &sectId)
	if err != nil {
		log.Println("Error while unmarshaling update sections request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := sectionController.sectionService.UpdateSections(ctx, &section, int64(sectId))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (sectionController SectionController) DeleteSections(ctx *gin.Context) {

	sectId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseErr := sectionController.sectionService.DeleteSections(ctx, int64(sectId))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}
