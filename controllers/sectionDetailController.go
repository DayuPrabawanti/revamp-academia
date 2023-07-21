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

func (sectionDetailController SectionDetailController) UpdateSectionDetail(ctx *gin.Context) {

	secdId, err := strconv.Atoi(ctx.Param("id"))

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

	var sectionDetail dbcontext.CreatesectionDetailParams
	err = json.Unmarshal(body, &sectionDetail)
	if err != nil {
		log.Println("Error while unmarshaling update sections request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := sectionDetailController.sectionDetailService.UpdateSectionDetail(ctx, &sectionDetail, int64(secdId))
	// if response != nil {
	// 	ctx.AbortWithStatusJSON(response.Status, response)
	// 	return
	// }

	ctx.JSON(http.StatusOK, response)

}

func (sectionDetailController SectionDetailController) DeleteSectionDetail(ctx *gin.Context) {

	secdId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseErr := sectionDetailController.sectionDetailService.DeleteSectionDetail(ctx, int64(secdId))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}
