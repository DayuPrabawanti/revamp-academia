package curriculumcontrollers

import (
	"log"
	"net/http"
	"strconv"

	"codeid.revampacademy/models/features"
	services "codeid.revampacademy/services/curriculumServices"
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

	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("pageNo", "0"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "3"))
	searchBy := ctx.DefaultQuery("searchBy", "")

	metadata := features.Metadata{
		PageNo:   pageNo,
		PageSize: pageSize,
		SearchBy: searchBy,
	}
	response, responseErr := progReviewsController.progReviewsService.GetListProgReviews(ctx, &metadata)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (progReviewsController ProgReviewsController) GetProgramReviews(ctx *gin.Context) {

	prowUserEntityId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Println("Error while reading paramater id", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, responseErr := progReviewsController.progReviewsService.GetProgramReviews(ctx, int64(prowUserEntityId))
	if responseErr != nil {

		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
