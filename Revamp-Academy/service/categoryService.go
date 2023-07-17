package service

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories"
	"github.com/gin-gonic/gin"
)

type CategoryService struct {
	categoryRepo *repositories.CategoryRepo
}

func NewCategoryService(categoryRepo *repositories.CategoryRepo) *CategoryService {
	return &CategoryService{
		categoryRepo: categoryRepo,
	}
}

func (cs CategoryService) GetListJobHireCategory(ctx *gin.Context) ([]*models.JobhireJobCategory, *models.ResponseError) {
	return cs.categoryRepo.GetListCategoryJob(ctx)
}
