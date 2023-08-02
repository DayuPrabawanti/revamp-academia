package curriculumRepositories

import (
	"net/http"

	models "codeid.revampacademy/models"
	dbcontext "codeid.revampacademy/repositories/curriculumRepositories/dbContext"

	"github.com/gin-gonic/gin"
)

func (pr ProgEntityRepository) GetListProgReviews(ctx *gin.Context) ([]*models.CurriculumProgramReview, *models.ResponseError) {

	store := dbcontext.New(pr.dbHandler)
	progReview, err := store.ListProgReviews(ctx)

	listProgReviews := make([]*models.CurriculumProgramReview, 0)

	for _, v := range progReview {
		progReviews := &models.CurriculumProgramReview{
			ProwUserEntityID: v.ProwUserEntityID,
			ProwProgEntityID: v.ProwProgEntityID,
			ProwReview:       v.ProwReview,
			ProwRating:       v.ProwRating,
			ProwModifiedDate: v.ProwModifiedDate,
		}
		listProgReviews = append(listProgReviews, progReviews)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listProgReviews, nil
}

func (pr ProgEntityRepository) GetProgramReviews(ctx *gin.Context, id int64) (*models.CurriculumProgramReview, *models.ResponseError) {

	store := dbcontext.New(pr.dbHandler)
	programReviews, err := store.GetProgramReviews(ctx, int16(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &programReviews, nil
}
