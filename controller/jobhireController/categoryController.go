package jobhireController

import (
	"net/http"

	// "codeid.revampacademy/service"
	"codeid.revampacademy/service/jobhireService"
	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	categoryService *jobhireService.CategoryService
}

func NewCategoryController(categoryService *jobhireService.CategoryService) *CategoryController {
	return &CategoryController{
		categoryService: categoryService,
	}
}

func (cc CategoryController) GetListCategoryControl(ctx *gin.Context) {
	response, responseErr := cc.categoryService.GetListJobHireCategory(ctx)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
	}
	ctx.JSON(http.StatusOK, response)
}
