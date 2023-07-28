package curriculumRepositories

import (
	"database/sql"
	"net/http"

	mod "codeid.revampacademy/models"
	models "codeid.revampacademy/models"
	"codeid.revampacademy/repositories/curriculumRepositories/dbContext"
	db "codeid.revampacademy/repositories/curriculumRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type CurriculumRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
	dbQueries   dbContext.Queries
}

func NewCurriculumRepository(dbHandler *sql.DB) *CurriculumRepository {
	return &CurriculumRepository{
		dbHandler: dbHandler,
		dbQueries: *dbContext.New(dbHandler),
	}
}

func (per CurriculumRepository) GetProgramEntity(ctx *gin.Context, id int64) (*mod.CurriculumProgramEntityMockup, *mod.ResponseError) {
	progentity, err := per.dbQueries.Getprogramentity(ctx, int32(id))

	if err != nil {
		return nil, &mod.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &progentity, nil
}

func (per CurriculumRepository) GetSection(ctx *gin.Context, id int64) (*[]mod.CurriculumSectionGet, *mod.ResponseError) {

	section, err := per.dbQueries.Getsection(ctx, int64(id))

	if err != nil {
		return nil, &mod.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &section, nil
}
func (per CurriculumRepository) GetMsCategory(ctx *gin.Context, id int64) (*[]mod.MasterCategoryProgramEntity, *mod.ResponseError) {

	msCategory, err := per.dbQueries.GetMsCategory(ctx, int32(id))

	if err != nil {
		return nil, &mod.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &msCategory, nil
}

func (per CurriculumRepository) GetProgramEntityDesctiption(ctx *gin.Context, id int64) (*mod.CurriculumProgramEntityDescriptionMockup, *mod.ResponseError) {

	progentitydesc, err := per.dbQueries.Getprogramentitydescription(ctx, int32(id))

	if err != nil {
		return nil, &mod.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &progentitydesc, nil
}

func (per CurriculumRepository) GetSectionDetails(ctx *gin.Context, id int64) (*[]mod.CurriculumSectionDetailMockup, *models.ResponseError) {

	sectionDetail, err := per.dbQueries.Getsectiondetail(ctx, int64(id))

	if err != nil {
		return nil, &mod.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &sectionDetail, nil
}

func (per CurriculumRepository) GetCurriculum(ctx *gin.Context, id int64) (*[]mod.CurriculumAll, *mod.ResponseError) {
	// Ambil data CurriculumProgramEntity
	progEntity, err := per.GetProgramEntity(ctx, id)
	if err != nil {
		return nil, err
	}

	// Ambil data CurriculumSection
	sections, err := per.GetSection(ctx, id)
	if err != nil {
		return nil, err
	}
	// Ambil data CurriculumSectionDetailMockup berdasarkan sectID dari setiap CurriculumSectionGet
	for i := range *sections {
		sectionDetails, err := per.GetSectionDetails(ctx, int64((*sections)[i].SectID))
		if err != nil {
			return nil, err
		}
		(*sections)[i].CurriculumSectionDetailMockup = *sectionDetails
	}
	msCategory, err := per.GetMsCategory(ctx, id)
	if err != nil {
		return nil, err
	}
	// progEntityDesc, err := per.GetProgramEntityDesctiption(ctx, id)
	// if err != nil {
	// 	return nil, err
	// }
	// sectiondetails, err := per.GetSectionDetails(ctx, id)
	// if err != nil {
	// 	return nil, err
	// }

	// Buat objek GetGabung
	getCurriculum := &mod.CurriculumAll{
		CurriculumProgramEntityMockup: *progEntity,
		CurriculumSectionGet:          *sections,
		MasterCategory:                *msCategory,
	}

	// Masukkan objek GetGabung ke dalam slice, karena GetGabung berada dalam slice
	curriculumSlice := []mod.CurriculumAll{*getCurriculum}

	return &curriculumSlice, nil
}

//UPDATE

func (per CurriculumRepository) UpdateProgramEntity(ctx *gin.Context, progentityParams *db.UpdateprogramentityParams) *mod.ResponseError {

	err := per.dbQueries.Updateprogramentity(ctx, *progentityParams)

	if err != nil {
		return &mod.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &mod.ResponseError{
		Message: "data has been update",
		Status:  http.StatusOK,
	}
}
func (per CurriculumRepository) Updateprogramentitydescription(ctx *gin.Context, progentitydescParams *db.UpdateprogramentitydescriptionParams) *mod.ResponseError {

	err := per.dbQueries.Updateprogramentitydescription(ctx, *progentitydescParams)

	if err != nil {
		return &mod.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &mod.ResponseError{
		Message: "data has been update",
		Status:  http.StatusOK,
	}
}

func (per CurriculumRepository) UpdateScore(ctx *gin.Context, scoreParams *db.UpdateScoreParams) *mod.ResponseError {

	err := per.dbQueries.UpdateScore(ctx, *scoreParams)

	if err != nil {
		return &mod.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &mod.ResponseError{
		Message: "data has been update",
		Status:  http.StatusOK,
	}
}

func (per CurriculumRepository) UpdateCurriculum(ctx *gin.Context, gabungParams *db.UpdateCurriculum) *mod.ResponseError {

	err := per.dbQueries.Updateprogramentity(ctx, *&gabungParams.UpdateprogramentityParams)

	err2 := per.dbQueries.Updateprogramentitydescription(ctx, *&gabungParams.UpdateprogramentitydescriptionParams)

	err3 := per.dbQueries.UpdateScore(ctx, *&gabungParams.UpdateScoreParams)

	if err != nil {
		return &mod.ResponseError{
			Message: "error when update Progentity",
			Status:  http.StatusInternalServerError,
		}
	}
	if err2 != nil {
		return &mod.ResponseError{
			Message: "error when update Prog Description",
			Status:  http.StatusInternalServerError,
		}
	}
	if err3 != nil {
		return &mod.ResponseError{
			Message: "error when update Score",
			Status:  http.StatusInternalServerError,
		}
	}

	return &mod.ResponseError{
		Message: "data has been update",
		Status:  http.StatusOK,
	}
}
