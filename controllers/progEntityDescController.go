package controllers

import (
	"net/http"

	"codeid.revampacademy/services"
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
