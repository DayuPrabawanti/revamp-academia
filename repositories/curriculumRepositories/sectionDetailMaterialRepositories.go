package curriculumRepositories

import (
	"net/http"

	models "codeid.revampacademy/models"
	dbcontext "codeid.revampacademy/repositories/curriculumRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

func (sdm ProgEntityRepository) GetListSectionDetailMaterial(ctx *gin.Context) ([]*models.CurriculumSectionDetailMaterial, *models.ResponseError) {

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

func (sdm ProgEntityRepository) GetSectionDetailMaterial(ctx *gin.Context, id int64) (*models.CurriculumSectionDetailMaterial, *models.ResponseError) {

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

func (sdm ProgEntityRepository) CreatesectiondetailMaterial(ctx *gin.Context, sectionDetailMaterialParams *dbcontext.CreatesectionDetailMaterialParams) (*models.CurriculumSectionDetailMaterial, *models.ResponseError) {

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

func (sdm ProgEntityRepository) UpdateSectionDetailMaterial(ctx *gin.Context, sectionDetailMaterialParams *dbcontext.CreatesectionDetailMaterialParams) *models.ResponseError {

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

func (sdm ProgEntityRepository) DeleteSectionDetailMaterial(ctx *gin.Context, id int64) *models.ResponseError {

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
