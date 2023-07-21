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
}

func NewProgramEntityRepository(dbHandler *sql.DB) *ProgramEntityRepository {
	return &ProgramEntityRepository{
		dbHandler:               dbHandler,
		SectionRepository:       SectionRepository{},
		SectionDetailRepository: SectionDetailRepository{},
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

func (se SectionRepository) GetSections(ctx *gin.Context, id int64) (*models.CurriculumSection, *models.ResponseError) {

	store := dbcontext.New(se.dbHandler)
	section, err := store.GetSections(ctx, int16(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &section, nil
}

func (se SectionRepository) Createsections(ctx *gin.Context, sectionsParams *dbcontext.CreatesectionsParams) (*models.CurriculumSection, *models.ResponseError) {

	store := dbcontext.New(se.dbHandler)
	section, err := store.Createsections(ctx, *sectionsParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return section, nil
}

func (se SectionRepository) UpdateSections(ctx *gin.Context, sectionParams *dbcontext.CreatesectionsParams) *models.ResponseError {

	store := dbcontext.New(se.dbHandler)
	err := store.UpdateSections(ctx, *sectionParams)

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

func (se SectionRepository) DeleteSections(ctx *gin.Context, id int64) *models.ResponseError {

	store := dbcontext.New(se.dbHandler)
	err := store.DeleteSections(ctx, int16(id))

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

func (sd SectionDetailRepository) GetSectionDetail(ctx *gin.Context, id int64) (*models.CurriculumSectionDetail, *models.ResponseError) {

	store := dbcontext.New(sd.dbHandler)
	sectionDetail, err := store.GetSectionDetail(ctx, int16(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &sectionDetail, nil
}

func (sd SectionDetailRepository) CreateSectionDetail(ctx *gin.Context, sectionDetailParams *dbcontext.CreatesectionDetailParams) (*models.CurriculumSectionDetail, *models.ResponseError) {

	store := dbcontext.New(sd.dbHandler)
	sectionDetail, err := store.CreateSectionDetail(ctx, *sectionDetailParams)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return sectionDetail, nil
}

func (sd SectionDetailRepository) UpdateSectionDetail(ctx *gin.Context, sectionDetailParams *dbcontext.CreatesectionDetailParams) *models.ResponseError {

	store := dbcontext.New(sd.dbHandler)
	err := store.UpdateSectionDetail(ctx, *sectionDetailParams)

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

func (sd SectionDetailRepository) DeleteSectionDetail(ctx *gin.Context, id int64) *models.ResponseError {

	store := dbcontext.New(sd.dbHandler)
	err := store.DeleteSectionDetail(ctx, int16(id))

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

func (ped ProgEntityDescRepository) GetProgEntityDesc(ctx *gin.Context, id int64) (*models.CurriculumProgramEntityDescription, *models.ResponseError) {

	store := dbcontext.New(ped.dbHandler)
	programEntityDescription, err := store.GetProgEntityDesc(ctx, int16(id))

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return (*models.CurriculumProgramEntityDescription)(&programEntityDescription), nil
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

func (pr ProgReviewsRepository) GetProgramReviews(ctx *gin.Context, id int64) (*models.CurriculumProgramReview, *models.ResponseError) {

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

// GROUP

func (pr ProgramEntityRepository) GroupList(ctx *gin.Context) ([]*models.Group, *models.ResponseError) {

	store := dbcontext.New(pr.dbHandler)
	group, err := store.Group(ctx)

	groupList := make([]*models.Group, 0)

	for _, v := range group {
		group := &models.Group{
			CurriculumProgramEntity: v.CurriculumProgramEntity,
			CurriculumSection:       v.CurriculumSection,
			CurriculumSectionDetail: v.CurriculumSectionDetail,
			MasterCategory:          v.MasterCategory,
		}
		groupList = append(groupList, group)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return groupList, nil
}

// func (pe ProgramEntityRepository) Group(ctx *gin.Context) ([]*models.Group, *models.ResponseError) {

// 	// Ambil daftar CurriculumProgramEntity
// 	programEntity, err := pe.GetListProgramEntity(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Ambil daftar CurriculumSection
// 	sections, err := pe.GetListSection(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Ambil daftar CurriculumSectionDetail
// 	sectionDetails, err := pe.GetListSectionDetail(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Ambil daftar MasterCategory
// 	masterCategories, err := pe.GetListMasterCategory(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Buat map untuk menyimpan CurriculumSection berdasarkan SectProgEntityID
// 	sectionMap := make(map[int32]*models.CurriculumSection)
// 	for _, section := range sections {
// 		sectionMap[section.SectProgEntityID] = section
// 	}
// 	sectionDetailMap := make(map[int32]*models.CurriculumSectionDetail)
// 	for _, sectionDetail := range sectionDetails {
// 		sectionDetailMap[sectionDetail.SecdID] = sectionDetail
// 	}
// 	masterCategoryMap := make(map[int32]*models.MasterCategory)
// 	for _, masterCategory := range masterCategories {
// 		masterCategoryMap[masterCategory.CateID] = masterCategory
// 	}

// 	// Gabungkan data dari CurriculumProgramEntity dan Curriculum Section
// 	groupList := make([]*models.Group, 0)
// 	for _, progEntity := range programEntity {
// 		section, ok := sectionMap[progEntity.ProgEntityID]
// 		if !ok {
// 			// Jika tidak ada CurriculumSection yang cocok, lanjutkan ke iterasi selanjutnya
// 			continue
// 		}
// 		sectionDetail, ok := sectionDetailMap[section.SectID]
// 		if !ok {
// 			// Jika tidak ada CurriculumSectionDetail yang cocok, lanjutkan ke iterasi selanjutnya
// 			continue
// 		}
// 		masterCategory, ok := masterCategoryMap[progEntity.ProgCateID]
// 		if !ok {
// 			// Jika tidak ada MasterCategory yang cocok, lanjutkan ke iterasi selanjutnya
// 			continue
// 		}

// 		group := &models.Group{
// 			CurriculumProgramEntity: *progEntity,
// 			CurriculumSection:       *section,
// 			CurriculumSectionDetail: *sectionDetail,
// 			MasterCategory:          *masterCategory,
// 		}

// 		groupList = append(groupList, group)
// 	}

// 	return groupList, nil
// }

func (per ProgramEntityRepository) GetListMasterCategory(ctx *gin.Context) ([]*models.MasterCategory, *models.ResponseError) {

	store := dbcontext.New(per.dbHandler)
	section, err := store.ListMasterCategories(ctx)

	listSections := make([]*models.MasterCategory, 0)

	for _, v := range section {
		sections := &models.MasterCategory{
			CateID:           v.CateID,
			CateName:         v.CateName,
			CateCateID:       v.CateCateID,
			CateModifiedDate: v.CateModifiedDate,
		}
		listSections = append(listSections, sections)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listSections, nil
}

func (pe ProgramEntityRepository) CreateGroup(ctx *gin.Context, groupParams *dbcontext.CreateGroupParams) (*models.Group, *models.ResponseError) {
	// Buat CurriculumProgramEntity
	programEntity, err := pe.CreateProgramEntity(ctx, &groupParams.CreateProgramEntityParams)
	if err != nil {
		return nil, err
	}

	// Buat CurriculumSection
	//gabungParams.CreatesectionsParams.SectID = gabungParams.Createprogram_entityParams.ProgEntityID
	section, err := pe.Createsections(ctx, &groupParams.CreatesectionsParams)
	if err != nil {
		return nil, err
	}

	// Buat Gabung
	group := &models.Group{
		CurriculumProgramEntity: *programEntity,
		CurriculumSection:       *section,
	}

	return group, nil
}
