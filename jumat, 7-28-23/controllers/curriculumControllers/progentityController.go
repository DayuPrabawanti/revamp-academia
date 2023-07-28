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

type ProgEntityController struct {
	progentityService *services.ProgEntityService
}

func NewProgEntityController(progentityService *services.ProgEntityService) *ProgEntityController {
	return &ProgEntityController{
		progentityService: progentityService,
	}
}

func (progentityController ProgEntityController) GetListProgEntity(ctx *gin.Context) {
	response, responerr := progentityController.progentityService.GetListProgEntity(ctx)
	if responerr != nil {
		ctx.JSON(responerr.Status, responerr)
	}
	ctx.JSON(http.StatusOK, response)
}
func (progentityController ProgEntityController) GetListMasterCategory(ctx *gin.Context) {
	response, responerr := progentityController.progentityService.GetListMasterCategory(ctx)
	if responerr != nil {
		ctx.JSON(responerr.Status, responerr)
	}
	ctx.JSON(http.StatusOK, response)
	//ctx.JSON(http.StatusOK, "bisa")
}
func (progentityController ProgEntityController) GetListSection(ctx *gin.Context) {
	response, responerr := progentityController.progentityService.GetListSection(ctx)
	if responerr != nil {
		ctx.JSON(responerr.Status, responerr)
	}
	ctx.JSON(http.StatusOK, response)
	//ctx.JSON(http.StatusOK, "bisa")
}
func (progentityController ProgEntityController) GetListSectionDetail(ctx *gin.Context) {
	response, responerr := progentityController.progentityService.GetListSectionDetail(ctx)
	if responerr != nil {
		ctx.JSON(responerr.Status, responerr)
	}
	ctx.JSON(http.StatusOK, response)
	//ctx.JSON(http.StatusOK, "bisa")
}
func (progentityController ProgEntityController) GetListGabung(ctx *gin.Context) {
	response, responerr := progentityController.progentityService.Gabung(ctx)
	if responerr != nil {
		ctx.JSON(responerr.Status, responerr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (progentityController ProgEntityController) GetProgEntity(ctx *gin.Context) {
	progEntityID, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := progentityController.progentityService.GetProgEntity(ctx, int64(progEntityID))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
func (progentityController ProgEntityController) GetSection(ctx *gin.Context) {
	sectionId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := progentityController.progentityService.GetSection(ctx, int64(sectionId))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
func (progentityController ProgEntityController) GetGabung(ctx *gin.Context) {
	sectionId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := progentityController.progentityService.GetGabung(ctx, int64(sectionId))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
func (progentityController ProgEntityController) CreateProgEntity(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var progentity db.Createprogram_entityParams
	err = json.Unmarshal(body, &progentity)
	if err != nil {
		log.Println("Error while unmarshaling create category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := progentityController.progentityService.CreateProgEntity(ctx, &progentity)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}
func (progentityController ProgEntityController) CreateSection(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create Section request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var sections db.CreatesectionsParams
	err = json.Unmarshal(body, &sections)
	if err != nil {
		log.Println("Error while unmarshaling create Section request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := progentityController.progentityService.CreateSections(ctx, &sections)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (progentityController ProgEntityController) CreateGabung(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create Gabung request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var gabungParams db.CreateGabung
	err = json.Unmarshal(body, &gabungParams)
	if err != nil {
		log.Println("Error while unmarshaling create Gabung request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := progentityController.progentityService.CreateGabung(ctx, &gabungParams)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (progentityController ProgEntityController) UpdateProgEntity(ctx *gin.Context) {

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

	var progentity db.Createprogram_entityParams
	err = json.Unmarshal(body, &progentity)
	if err != nil {
		log.Println("Error while unmarshaling update category request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response := progentityController.progentityService.UpdateProgEntity(ctx, &progentity, int64(progentityId))
	if response != nil {
		ctx.AbortWithStatusJSON(response.Status, response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (progentityController ProgEntityController) DeleteProgEntity(ctx *gin.Context) {

	progentityId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseErr := progentityController.progentityService.DeleteProgEntity(ctx, int64(progentityId))
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.Status(http.StatusNoContent)
}
