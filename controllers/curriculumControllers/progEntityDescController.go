package curriculumcontrollers

import (
	"log"
	"net/http"
	"strconv"

	services "codeid.revampacademy/services/curriculumServices"
	"github.com/gin-gonic/gin"
)

type ProgEntityDescController struct {
	progEntityDescService *services.ProgEntityDescService
}

func NewProgEntityDescController(progEntityDescService *services.ProgEntityDescService) *ProgEntityDescController {
	return &ProgEntityDescController{
		progEntityDescService: progEntityDescService,
	}
}

func (progEntityDescController ProgEntityDescController) GetListProgEntityDesc(ctx *gin.Context) {
	response, responseErr := progEntityDescController.progEntityDescService.GetListProgEntityDesc(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (progEntityDescController ProgEntityDescController) GetProgEntityDesc(ctx *gin.Context) {

	predProgEntityID, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := progEntityDescController.progEntityDescService.GetProgEntityDesc(ctx, int64(predProgEntityID))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
