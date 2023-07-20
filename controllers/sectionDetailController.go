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

type SectionDetailController struct {
	sectionDetailService *services.SectionDetailService
}

func NewSectionDetailController(sectionDetailService *services.SectionDetailService) *SectionDetailController {
	return &SectionDetailController{
		sectionDetailService: sectionDetailService,
	}
}

func (sectionDetailController SectionDetailController) GetListSectionDetail(ctx *gin.Context) {
	response, responseErr := sectionDetailController.sectionDetailService.GetListSectionDetail(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (sectionDetailController SectionDetailController) GetSectionDetail(ctx *gin.Context) {

	secdId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := sectionDetailController.sectionDetailService.GetSectionDetail(ctx, int64(secdId))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (sectionDetailController SectionDetailController) CreateSectionDetail(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create section detail request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var sectionDetail dbcontext.CreatesectionDetailParams
	err = json.Unmarshal(body, &sectionDetail)
	if err != nil {
		log.Println("Error while unmarshaling create section detail request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := sectionDetailController.sectionDetailService.CreateSectionDetail(ctx, &sectionDetail)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

// func (sectionController SectionController) UpdateSections(ctx *gin.Context) {

// 	sectId, err := strconv.Atoi(ctx.Param("id"))

// 	if err != nil {
// 		log.Println("Error while reading paramater id", err)
// 		ctx.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}

// 	body, err := io.ReadAll(ctx.Request.Body)
// 	if err != nil {
// 		log.Println("Error while reading update sections request body", err)
// 		ctx.AbortWithError(http.StatusInternalServerError, err)
// 		return
// 	}

// 	var section dbcontext.CreatesectionsParams
// 	err = json.Unmarshal(body, &sectId)
// 	if err != nil {
// 		log.Println("Error while unmarshaling update sections request body", err)
// 		ctx.AbortWithError(http.StatusInternalServerError, err)
// 		return
// 	}

// 	response := sectionController.sectionService.UpdateSections(ctx, &section, int64(sectId))
// 	if response != nil {
// 		ctx.AbortWithStatusJSON(response.Status, response)
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, response)

// }

// func (sectionController SectionController) DeleteSections(ctx *gin.Context) {

// 	sectId, err := strconv.Atoi(ctx.Param("id"))

// 	if err != nil {
// 		log.Println("Error while reading paramater id", err)
// 		ctx.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}

// 	responseErr := sectionController.sectionService.DeleteSections(ctx, int64(sectId))
// 	if responseErr != nil {
// 		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
// 		return
// 	}

// 	ctx.Status(http.StatusNoContent)
// }
