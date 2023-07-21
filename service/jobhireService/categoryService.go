package jobhireService

import (
	"codeid.revampacademy/models"
	// "codeid.revampacademy/repositories"
	"codeid.revampacademy/repositories/jobhireRepositories"
	"github.com/gin-gonic/gin"
)

type CategoryService struct {
	categoryRepo *jobhireRepositories.CategoryRepo
}

func NewCategoryService(categoryRepo *jobhireRepositories.CategoryRepo) *CategoryService {
	return &CategoryService{
		categoryRepo: categoryRepo,
	}
}

func (cs CategoryService) GetListJobHireCategory(ctx *gin.Context) ([]*models.JobhireJobCategory, *models.ResponseError) {
	return cs.categoryRepo.GetListCategoryJob(ctx)
}
