package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (programEntityController ProgramEntityController) GroupList(ctx *gin.Context) {
	response, responseErr := programEntityController.programEntityService.GroupList(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
