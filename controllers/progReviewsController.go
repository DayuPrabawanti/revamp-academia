package controllers

import (
	"net/http"

	"codeid.revampacademy/services"
	"github.com/gin-gonic/gin"
)

type ProgReviewsController struct {
	progReviewsService *services.ProgReviewService
}

func NewProgReviewsController(progReviewsService *services.ProgReviewService) *ProgReviewsController {
	return &ProgReviewsController{
		progReviewsService: progReviewsService,
	}
}

func (progReviewsController ProgReviewsController) GetListProgReviews(ctx *gin.Context) {
	response, responseErr := progReviewsController.progReviewsService.GetListProgReviews(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
