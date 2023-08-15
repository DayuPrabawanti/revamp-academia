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

func (pe ProgramentityRepository) GetListProgramEntity(ctx *gin.Context) ([]*models.CurriculumProgramEntity, *models.ResponseError) {

	store := dbContext.New(pe.dbHandler)
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

func (pe ProgramentityRepository) UpdateProgramEntity(ctx *gin.Context, curriParams *dbContext.CurriculumProgramEntityParams) *models.ResponseError {

	store := dbContext.New(pe.dbHandler)
	err := store.UpdateProgramEntity(ctx, *curriParams)

	if err != nil {
		return &models.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been update",
		Status:  http.StatusOK,
	}
}

func (pe ProgramentityRepository) GetProgEntity(ctx *gin.Context, id int64) (*models.CurriculumProgramEntity, *models.ResponseError) {

	store := dbContext.New(pe.dbHandler)
	progentityid, err := store.GetProgEntity(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &progentityid, nil
}
