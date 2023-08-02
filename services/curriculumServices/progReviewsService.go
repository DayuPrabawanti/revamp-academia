package curriculumServices

import (
	models "codeid.revampacademy/models"
	"codeid.revampacademy/models/features"
	repo "codeid.revampacademy/repositories/curriculumRepositories"
	"github.com/gin-gonic/gin"
)

type ProgReviewService struct {
	repositoryManager repo.RepositoryManager
}

func NewProgReviewsService(repoMgr *repo.RepositoryManager) *ProgReviewService {
	return &ProgReviewService{
		repositoryManager: *repoMgr,
	}
}

func (pr ProgReviewService) GetListProgReviews(ctx *gin.Context, metadata *features.Metadata) ([]*models.CurriculumProgramReview, *models.ResponseError) {
	return pr.repositoryManager.ProgEntityRepository.GetListProgReviews(ctx)
}

func (pr ProgReviewService) GetProgramReviews(ctx *gin.Context, id int64) (*models.CurriculumProgramReview, *models.ResponseError) {
	return pr.repositoryManager.ProgEntityRepository.GetProgramReviews(ctx, id)
}
