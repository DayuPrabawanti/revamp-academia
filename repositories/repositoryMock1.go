package repositories

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	dbcontext "codeid.revampacademy/repositories/dbContext"
	"github.com/gin-gonic/gin"
)

type RepositoryMock1 struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewRepositoryMock1(dbHandler *sql.DB) *RepositoryMock1 {
	return &RepositoryMock1{
		dbHandler: dbHandler,
	}
}

func (mr RepositoryMock1) GetListMock1(ctx *gin.Context) ([]*models.CurriculumProgramEntity, *models.ResponseError) {

	store := dbcontext.New(mr.dbHandler)
	special_offer, err := store.Listprogram_entity(ctx)

	listSalesOffer := make([]*models.CurriculumProgramEntity, 0)

	for _, v := range special_offer {
		sales := &models.CurriculumProgramEntity{
			ProgTitle:        v.ProgTitle,
			ProgHeadline:     v.ProgHeadline,
			ProgLearningType: v.ProgLearningType,
			ProgImage:        v.ProgImage,
			ProgPrice:        v.ProgPrice,
			ProgDuration:     v.ProgDuration,
		}
		listSalesOffer = append(listSalesOffer, sales)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listSalesOffer, nil
}

func (rm RepositoryMock1) GetMockup1(ctx *gin.Context, nama string) (*models.CurriculumProgramEntity, *models.ResponseError) {

	store := dbcontext.New(rm.dbHandler)
	mockup1, err := store.Getprogram_entity(ctx, string(nama))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &mockup1, nil
}
