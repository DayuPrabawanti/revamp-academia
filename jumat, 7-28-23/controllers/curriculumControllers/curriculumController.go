package curriculumcontrollers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	db "codeid.revampacademy/repositories/curriculumRepositories/dbContext"
	services "codeid.revampacademy/services/curriculumServices"
	"github.com/gin-gonic/gin"
)

type CurriculumController struct {
	curriculumService *services.CurriculumService
}

func NewCurriculumController(curriculumService *services.CurriculumService) *CurriculumController {
	return &CurriculumController{
		curriculumService: curriculumService,
	}
}
func (curriculumService CurriculumController) GetCurriculum(ctx *gin.Context) {
	sectionId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := curriculumService.curriculumService.GetCurriculum(ctx, int64(sectionId))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (curriculumService CurriculumController) UpdateProgramEntity(ctx *gin.Context) {

	progentityId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var progentity db.UpdateprogramentityParams
	err = json.Unmarshal(body, &progentity)
	if err != nil {
		log.Println("Error while unmarshaling update category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := curriculumService.curriculumService.UpdateProgramEntity(ctx, &progentity, int64(progentityId))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (curriculumService CurriculumController) Updateprogramentitydescription(ctx *gin.Context) {

	progentityId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var progentity db.UpdateprogramentitydescriptionParams
	err = json.Unmarshal(body, &progentity)
	if err != nil {
		log.Println("Error while unmarshaling update category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := curriculumService.curriculumService.Updateprogramentitydescription(ctx, &progentity, int64(progentityId))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (curriculumCtr CurriculumController) UpdateScore(ctx *gin.Context) {

	progentityId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var progentity db.UpdateScoreParams
	err = json.Unmarshal(body, &progentity)
	if err != nil {
		log.Println("Error while unmarshaling update category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := curriculumCtr.curriculumService.UpdateScore(ctx, &progentity, int64(progentityId))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (curriculumService CurriculumController) UpdateCurriculum(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create Gabung request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var updategabungParams db.UpdateCurriculum
	err = json.Unmarshal(body, &updategabungParams)
	if err != nil {
		log.Println("Error while unmarshaling create Gabung request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	responseErr := curriculumService.curriculumService.UpdateCurriculum(ctx, &updategabungParams)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
