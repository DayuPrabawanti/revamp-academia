package curriculumRepositories

import (
	"database/sql"
	"net/http"

	models "codeid.revampacademy/models"
	dbcontext "codeid.revampacademy/repositories/curriculumRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type SectionDetailMaterialRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewSectionDetailMaterialRepository(dbHandler *sql.DB) *SectionDetailMaterialRepository {
	return &SectionDetailMaterialRepository{
		dbHandler: dbHandler,
	}
}

func (sdm SectionDetailMaterialRepository) GetListSectionDetailMaterial(ctx *gin.Context) ([]*models.CurriculumSectionDetailMaterial, *models.ResponseError) {

	store := dbcontext.New(sdm.dbHandler)
	sectionDetailMaterials, err := store.ListSectionDetailMaterial(ctx)

	listSectionDetailMaterial := make([]*models.CurriculumSectionDetailMaterial, 0)

	for _, v := range sectionDetailMaterials {
		sectionDetailMaterial := &models.CurriculumSectionDetailMaterial{
			SedmID:           v.SedmID,
			SedmFilename:     v.SedmFilename,
			SedmFilesize:     v.SedmFilesize,
			SedmFiletype:     v.SedmFiletype,
			SedmFilelink:     v.SedmFilelink,
			SedmModifiedDate: v.SedmModifiedDate,
			SedmSecdID:       v.SedmSecdID,
		}
		listSectionDetailMaterial = append(listSectionDetailMaterial, sectionDetailMaterial)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listSectionDetailMaterial, nil
}

func (sdm SectionDetailMaterialRepository) GetSectionDetailMaterial(ctx *gin.Context, id int64) (*models.CurriculumSectionDetailMaterial, *models.ResponseError) {

	store := dbcontext.New(sdm.dbHandler)
	sectionDetailMaterial, err := store.GetSectionDetailMaterial(ctx, int16(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &sectionDetailMaterial, nil
}

func (sdm SectionDetailMaterialRepository) CreatesectiondetailMaterial(ctx *gin.Context, sectionDetailMaterialParams *dbcontext.CreatesectionDetailMaterialParams) (*models.CurriculumSectionDetailMaterial, *models.ResponseError) {

	store := dbcontext.New(sdm.dbHandler)
	sectionDetailMaterial, err := store.CreatesectiondetailMaterial(ctx, *sectionDetailMaterialParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return sectionDetailMaterial, nil
}

func (sdm SectionDetailMaterialRepository) UpdateSectionDetailMaterial(ctx *gin.Context, sectionDetailMaterialParams *dbcontext.CreatesectionDetailMaterialParams) *models.ResponseError {

	store := dbcontext.New(sdm.dbHandler)
	err := store.UpdateSectionDetailMaterial(ctx, *sectionDetailMaterialParams)

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

func (sdm SectionDetailMaterialRepository) DeleteSectionDetailMaterial(ctx *gin.Context, id int64) *models.ResponseError {

	store := dbcontext.New(sdm.dbHandler)
	err := store.DeleteSectionDetailMaterial(ctx, int16(id))

	if err != nil {
		return &models.ResponseError{
			Message: "error when delete",
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.ResponseError{
		Message: "data has been deleted",
		Status:  http.StatusOK,
	}
}
