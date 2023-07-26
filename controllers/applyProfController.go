package controllers

import (
	"net/http"

	"codeid.revampacademy/services"
	"github.com/gin-gonic/gin"
)

type ApplyProfController struct {
	ApplyProfService *services.ApplyProfService
}

func NewApplyProfController(ApplyProfService *services.ApplyProfService) *ApplyProfController {
	return &ApplyProfController{
		ApplyProfService: ApplyProfService,
	}
}

func (ApplyProfController ApplyProfController) ListApplyProfHttp(ctx *gin.Context) {
	coba := ctx.Query("nama")
	response, responseErr := ApplyProfController.ApplyProfService.ListApplyProfService(ctx, string(coba))

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
