package bootcampRepository

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/bootcampRepository/dbContext"
	"github.com/gin-gonic/gin"
)

type ProgramentityRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewProgramentityRepository(dbHandler *sql.DB) *ProgramentityRepository {
	return &ProgramentityRepository{
		dbHandler: dbHandler,
	}
}

func (cr ProgramentityRepository) GetListProgramEntity(ctx *gin.Context) ([]*models.CurriculumProgramEntity, *models.ResponseError) {

	store := dbContext.New(cr.dbHandler)
	program_entity, err := store.ListProgramEntity(ctx)

	listProgramEntity := make([]*models.CurriculumProgramEntity, 0)

	for _, v := range program_entity {
		programEntity := &models.CurriculumProgramEntity{
			ProgEntityID:     v.ProgEntityID,
			ProgTitle:        v.ProgTitle,
			ProgHeadline:     v.ProgHeadline,
			ProgType:         v.ProgType,
			ProgLearningType: v.ProgLearningType,
			ProgRating:       v.ProgRating,
			ProgBestSeller:   v.ProgBestSeller,
			ProgPrice:        v.ProgPrice,
			ProgLanguage:     v.ProgLanguage,
			ProgModifiedDate: v.ProgModifiedDate,
			ProgDuration:     v.ProgDuration,
			ProgDurationType: v.ProgDurationType,
			ProgTagSkill:     v.ProgTagSkill,
			ProgCityID:       v.ProgCityID,
			ProgCateID:       v.ProgCateID,
			ProgCreatedBy:    v.ProgCreatedBy,
			ProgStatus:       v.ProgStatus,
		}
		listProgramEntity = append(listProgramEntity, programEntity)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listProgramEntity, nil
}
