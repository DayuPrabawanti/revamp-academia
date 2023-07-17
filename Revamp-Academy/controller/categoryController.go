package controller

import (
	"net/http"

	"codeid.revampacademy/service"
	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	categoryService *service.CategoryService
}

func NewCategoryController(categoryService *service.CategoryService) *CategoryController {
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
