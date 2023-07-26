package salesRepositories

import (
	"net/http"

	"codeid.revampacademy/models"
	dbcontext "codeid.revampacademy/repositories/salesRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

func (rm RepositoryMock) GetMockupId(ctx *gin.Context, id int64) (*models.CurriculumProgramEntity, *models.ResponseError) {

	store := dbcontext.New(rm.dbHandler)
	mockup, err := store.GetProgramEntityId(ctx, int32(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &mockup, nil
}

func (rm RepositoryMock) GetListProgramMock(ctx *gin.Context, orderby string) ([]*models.CurriculumProgramEntity, *models.ResponseError) {

	store := dbcontext.New(rm.dbHandler)
	program_entity, err := store.ListProgram(ctx, orderby)

	listProgramEntity := make([]*models.CurriculumProgramEntity, 0)

	for _, v := range program_entity {
		sales := &models.CurriculumProgramEntity{
			ProgEntityID:     v.ProgEntityID,
			ProgTitle:        v.ProgTitle,
			ProgHeadline:     v.ProgHeadline,
			ProgType:         v.ProgType,
			ProgLearningType: v.ProgLearningType,
			ProgRating:       v.ProgRating,
			ProgTotalTraniee: v.ProgTotalTraniee,
			ProgModifiedDate: v.ProgModifiedDate,
			ProgImage:        v.ProgImage,
			ProgBestSeller:   v.ProgBestSeller,
			ProgPrice:        v.ProgPrice,
			ProgLanguage:     v.ProgLanguage,
			ProgDuration:     v.ProgDuration,
			ProgDurationType: v.ProgDurationType,
			ProgTagSkill:     v.ProgTagSkill,
			ProgCityID:       v.ProgCityID,
			ProgCateID:       v.ProgCateID,
			ProgCreatedBy:    v.ProgCreatedBy,
			ProgStatus:       v.ProgStatus,
		}
		listProgramEntity = append(listProgramEntity, sales)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listProgramEntity, nil
}
