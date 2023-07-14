package repositories

import (
	"database/sql"
	"net/http"

	models "codeid.revampacademy/models"
	"codeid.revampacademy/repositories/dbcontext"
	"github.com/gin-gonic/gin"
)

// PROGRAM ENTITY

type ProgramEntityRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
	SectionRepository
	SectionDetailRepository
	SectionDetailMaterialRepository
}

func NewProgramEntityRepository(dbHandler *sql.DB) *ProgramEntityRepository {
	return &ProgramEntityRepository{
		dbHandler: dbHandler,
	}
}

func (pe ProgramEntityRepository) GetListProgramEntity(ctx *gin.Context) ([]*models.CurriculumProgramEntity, *models.ResponseError) {

	store := dbcontext.New(pe.dbHandler)
	programEntity, err := store.ListProgramEntity(ctx)

	listProgramEntity := make([]*models.CurriculumProgramEntity, 0)

	for _, v := range programEntity {
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

func (pe ProgramEntityRepository) GetProgramEntity(ctx *gin.Context, id int64) (*models.CurriculumProgramEntity, *models.ResponseError) {

	store := dbcontext.New(pe.dbHandler)
	programEntity, err := store.GetProgramEntity(ctx, int16(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &programEntity, nil
}

func (pe ProgramEntityRepository) CreateProgramEntity(ctx *gin.Context, programEntityParams *dbcontext.CreateProgramEntityParams) (*models.CurriculumProgramEntity, *models.ResponseError) {

	store := dbcontext.New(pe.dbHandler)
	programEntity, err := store.CreateProgramEntity(ctx, *programEntityParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return programEntity, nil
}

func (pe ProgramEntityRepository) UpdateProgramEntity(ctx *gin.Context, programEntityParams *dbcontext.CreateProgramEntityParams) *models.ResponseError {

	store := dbcontext.New(pe.dbHandler)
	err := store.UpdateProgramEntity(ctx, *programEntityParams)

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

func (pe ProgramEntityRepository) DeleteProgramEntity(ctx *gin.Context, id int64) *models.ResponseError {

	store := dbcontext.New(pe.dbHandler)
	err := store.DeleteProgramEntity(ctx, int16(id))

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

// SECTIONS

type SectionRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewSectionRepository(dbHandler *sql.DB) *SectionRepository {
	return &SectionRepository{
		dbHandler: dbHandler,
	}
}

func (se SectionRepository) GetListSection(ctx *gin.Context) ([]*models.CurriculumSection, *models.ResponseError) {

	store := dbcontext.New(se.dbHandler)
	sections, err := store.ListSections(ctx)

	listSections := make([]*models.CurriculumSection, 0)

	for _, v := range sections {
		section := &models.CurriculumSection{
			SectTitle:        v.SectTitle,
			SectDescription:  v.SectDescription,
			SectTotalSection: v.SectTotalSection,
			SectTotalLecture: v.SectTotalLecture,
			SectTotalMinute:  v.SectTotalMinute,
			SectModifiedDate: v.SectModifiedDate,
			SectProgEntityID: v.SectProgEntityID,
			SectID:           v.SectID,
		}
		listSections = append(listSections, section)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listSections, nil
}

// SECTION DETAIL

type SectionDetailRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewSectionDetailRepository(dbHandler *sql.DB) *SectionDetailRepository {
	return &SectionDetailRepository{
		dbHandler: dbHandler,
	}
}

func (sd SectionDetailRepository) GetListSectionDetail(ctx *gin.Context) ([]*models.CurriculumSectionDetail, *models.ResponseError) {

	store := dbcontext.New(sd.dbHandler)
	sectionDetails, err := store.ListSectionDetail(ctx)

	listSectionDetail := make([]*models.CurriculumSectionDetail, 0)

	for _, v := range sectionDetails {
		sectionDetail := &models.CurriculumSectionDetail{
			SecdID:           v.SecdID,
			SecdTitle:        v.SecdTitle,
			SecdPreview:      v.SecdPreview,
			SecdScore:        v.SecdScore,
			SecdNote:         v.SecdNote,
			SecdMinute:       v.SecdMinute,
			SecdModifiedDate: v.SecdModifiedDate,
			SecdSectID:       v.SecdSectID,
		}
		listSectionDetail = append(listSectionDetail, sectionDetail)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listSectionDetail, nil
}

// SECTION DETAIL MATERIAL

type SectionDetailMaterialRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewSectionDetailMaterialRepository(dbHandler *sql.DB) *SectionDetailMaterialRepository {
	return &SectionDetailMaterialRepository{
		dbHandler: dbHandler,
	}
}

func (sd SectionDetailMaterialRepository) GetListSectionDetailMaterial(ctx *gin.Context) ([]*models.CurriculumSectionDetailMaterial, *models.ResponseError) {

	store := dbcontext.New(sd.dbHandler)
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

// PROGRAM ENTITY DESCRIPTION

type ProgEntityDescRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewProgEntityDescRepository(dbHandler *sql.DB) *ProgEntityDescRepository {
	return &ProgEntityDescRepository{
		dbHandler: dbHandler,
	}
}

func (ped ProgEntityDescRepository) GetListProgEntityDesc(ctx *gin.Context) ([]*models.CurriculumProgramEntityDescription, *models.ResponseError) {

	store := dbcontext.New(ped.dbHandler)
	progEntityDesc, err := store.ListProgEntityDesc(ctx)

	listProgEntityDesc := make([]*models.CurriculumProgramEntityDescription, 0)

	for _, v := range progEntityDesc {
		progEntityDesc := &models.CurriculumProgramEntityDescription{
			PredProgEntityID: v.PredProgEntityID,
			PredItemLearning: v.PredItemLearning,
			PredItemInclude:  v.PredItemInclude,
			PredRequirement:  v.PredRequirement,
			PredDescription:  v.PredDescription,
			PredTargetLevel:  v.PredTargetLevel,
		}
		listProgEntityDesc = append(listProgEntityDesc, progEntityDesc)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listProgEntityDesc, nil
}

// PROGRAM REVIEWS

type ProgReviewsRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewProgReviewsRepository(dbHandler *sql.DB) *ProgReviewsRepository {
	return &ProgReviewsRepository{
		dbHandler: dbHandler,
	}
}

func (pr ProgReviewsRepository) GetListProgReviews(ctx *gin.Context) ([]*models.CurriculumProgramReview, *models.ResponseError) {

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

// GROUP

func (per ProgramEntityRepository) Group(ctx *gin.Context) ([]*models.Group, *models.ResponseError) {

	// Ambil daftar CurriculumProgramEntity
	programEntity, err := per.GetListProgramEntity(ctx)
	if err != nil {
		return nil, err
	}

	// Ambil daftar CurriculumSection
	sections, err := per.GetListSection(ctx)
	if err != nil {
		return nil, err
	}

	// Buat map untuk menyimpan CurriculumSection berdasarkan SectProgEntityID
	sectionMap := make(map[int32]*models.CurriculumSection)
	for _, section := range sections {
		sectionMap[section.SectProgEntityID] = section
	}

	// Gabungkan data dari CurriculumProgramEntity dan Curriculum Section
	groupList := make([]*models.Group, 0)
	for _, progEntity := range programEntity {
		section, ok := sectionMap[progEntity.ProgEntityID]
		if !ok {
			// Jika tidak ada CurriculumSection yang cocok, lanjutkan ke iterasi selanjutnya
			continue
		}

		group := models.Group{
			CurriculumProgramEntity: *progEntity,
			CurriculumSection:       *section,
		}

		groupList = append(groupList, &group)
	}

	return groupList, nil
}
