package curriculumRepositories

import (
	"database/sql"
	"net/http"

	models "codeid.revampacademy/models"
	dbcontext "codeid.revampacademy/repositories/curriculumRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type SectionDetailRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewSectionDetailRepository(dbHandler *sql.DB) *SectionDetailRepository {
	return &SectionDetailRepository{
		dbHandler: dbHandler,
	}
}

func (sdm SectionDetailRepository) GetSectionDetail(ctx *gin.Context, id int64) (*models.CurriculumSectionDetail, *models.ResponseError) {

	store := dbcontext.New(sdm.dbHandler)
	sectionDetail, err := store.Getsection_detail(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &sectionDetail, nil
}
