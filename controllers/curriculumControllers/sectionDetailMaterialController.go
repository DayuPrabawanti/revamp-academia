package curriculumcontrollers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	dbcontext "codeid.revampacademy/repositories/curriculumRepositories/dbContext"
	services "codeid.revampacademy/services/curriculumServices"
	"github.com/gin-gonic/gin"
)

type SectionDetailMaterialController struct {
	sectionDetailMaterialService *services.SectionDetailMaterialService
}

func NewSectionDetailMaterialController(sectionDetailMaterialService *services.SectionDetailMaterialService) *SectionDetailMaterialController {
	return &SectionDetailMaterialController{
		sectionDetailMaterialService: sectionDetailMaterialService,
	}
}

func (sectionDetailMaterialController SectionDetailMaterialController) GetListSectionDetailMaterial(ctx *gin.Context) {
	response, responseErr := sectionDetailMaterialController.sectionDetailMaterialService.GetListSectionDetailMaterial(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (sectionDetailMaterialController SectionDetailMaterialController) GetSectionDetailMaterial(ctx *gin.Context) {

	sedmId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := sectionDetailMaterialController.sectionDetailMaterialService.GetSectionDetailMaterial(ctx, int64(sedmId))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (sectionDetailMaterialController SectionDetailMaterialController) CreatesectiondetailMaterial(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create section detail request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var sectionDetailMaterial dbcontext.CreatesectionDetailMaterialParams
	err = json.Unmarshal(body, &sectionDetailMaterial)
	if err != nil {
		log.Println("Error while unmarshaling create section detail request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := sectionDetailMaterialController.sectionDetailMaterialService.CreatesectiondetailMaterial(ctx, &sectionDetailMaterial)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (sectionDetailMaterialController SectionDetailMaterialController) UpdateSectionDetailMaterial(ctx *gin.Context) {

	sedmId, err := strconv.Atoi(ctx.Param("id"))

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

	var sectionDetailMaterial dbcontext.CreatesectionDetailMaterialParams
	err = json.Unmarshal(body, &sectionDetailMaterial)
	if err != nil {
		log.Println("Error while unmarshaling update sections request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := sectionDetailMaterialController.sectionDetailMaterialService.UpdateSectionDetailMaterial(ctx, &sectionDetailMaterial, int64(sedmId))
	// if response != nil {
	// 	ctx.AbortWithStatusJSON(response.Status, response)
	// 	return
	// }

	ctx.JSON(http.StatusOK, response)

}

func (sectionDetailMaterialController SectionDetailMaterialController) DeleteSectionDetailMaterial(ctx *gin.Context) {

	sedmID, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseErr := sectionDetailMaterialController.sectionDetailMaterialService.DeleteSectionDetailMaterial(ctx, int64(sedmID))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}
