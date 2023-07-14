package services

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories"
	"github.com/gin-gonic/gin"
)

type ProgReviewService struct {
	progReviewsRepository *repositories.ProgReviewsRepository
}

func NewProgReviewsService(progReviewsRepository *repositories.ProgReviewsRepository) *ProgReviewService {
	return &ProgReviewService{
		progReviewsRepository: progReviewsRepository,
	}
}

func (pr ProgReviewService) GetListProgReviews(ctx *gin.Context) ([]*models.CurriculumProgramReview, *models.ResponseError) {
	return pr.progReviewsRepository.GetListProgReviews(ctx)
}
