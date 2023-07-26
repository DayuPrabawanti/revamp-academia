package curriculumServices

import (
	models "codeid.revampacademy/models"
	repositories "codeid.revampacademy/repositories/curriculumRepositories"
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

func (pr ProgReviewService) GetProgramReviews(ctx *gin.Context, id int64) (*models.CurriculumProgramReview, *models.ResponseError) {
	return pr.progReviewsRepository.GetProgramReviews(ctx, id)
}
