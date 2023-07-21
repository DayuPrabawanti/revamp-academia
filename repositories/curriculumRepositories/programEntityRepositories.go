package curriculumRepositories

import (
	"database/sql"
	"net/http"

	mod "codeid.revampacademy/models"
	"codeid.revampacademy/repositories/curriculumRepositories/dbContext"
	db "codeid.revampacademy/repositories/curriculumRepositories/dbContext"

	"github.com/gin-gonic/gin"
)

type ProgEntityRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewProgEntityRepository(dbHandler *sql.DB) *ProgEntityRepository {
	return &ProgEntityRepository{
		dbHandler: dbHandler,
	}
}

func (per ProgEntityRepository) GetListProgEntity(ctx *gin.Context) ([]*mod.CurriculumProgramEntity, *mod.ResponseError) {

	store := dbContext.New(per.dbHandler)
	progentity, err := store.Listprogram_entity(ctx)

	listProgEntity := make([]*mod.CurriculumProgramEntity, 0)

	for _, v := range progentity {
		proentity := &mod.CurriculumProgramEntity{
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
		listProgEntity = append(listProgEntity, proentity)
	}

	if err != nil {
		return nil, &mod.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listProgEntity, nil
}
func (per ProgEntityRepository) GetListMasterCategory(ctx *gin.Context) ([]*mod.MasterCategory, *mod.ResponseError) {

	store := dbContext.New(per.dbHandler)
	section, err := store.ListMasterCategories(ctx)

	listSections := make([]*mod.MasterCategory, 0)

	for _, v := range section {
		sections := &mod.MasterCategory{
			CateID:           v.CateID,
			CateName:         v.CateName,
			CateCateID:       v.CateCateID,
			CateModifiedDate: v.CateModifiedDate,
		}
		listSections = append(listSections, sections)
	}

	if err != nil {
		return nil, &mod.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listSections, nil
}

func (per ProgEntityRepository) GetListSection(ctx *gin.Context) ([]*mod.CurriculumSection, *mod.ResponseError) {

	store := dbContext.New(per.dbHandler)
	section, err := store.Listsections(ctx)

	listSections := make([]*mod.CurriculumSection, 0)

	for _, v := range section {
		sections := &mod.CurriculumSection{
			SectID:           v.SectID,
			SectProgEntityID: v.SectProgEntityID,
			SectTitle:        v.SectTitle,
			SectDescription:  v.SectDescription,
			SectTotalSection: v.SectTotalSection,
			SectTotalLecture: v.SectTotalLecture,
			SectTotalMinute:  v.SectTotalMinute,
			SectModifiedDate: v.SectModifiedDate,
		}
		listSections = append(listSections, sections)
	}

	if err != nil {
		return nil, &mod.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listSections, nil
}
func (per ProgEntityRepository) GetListSectionDetail(ctx *gin.Context) ([]*mod.CurriculumSectionDetail, *mod.ResponseError) {

	store := dbContext.New(per.dbHandler)
	sectiondetail, err := store.Listsection_detail(ctx)

	listSectionsDetail := make([]*mod.CurriculumSectionDetail, 0)

	for _, v := range sectiondetail {
		sectionsDetail := &mod.CurriculumSectionDetail{
			SecdID:           v.SecdID,
			SecdTitle:        v.SecdTitle,
			SecdPreview:      v.SecdPreview,
			SecdScore:        v.SecdScore,
			SecdNote:         v.SecdNote,
			SecdMinute:       v.SecdMinute,
			SecdModifiedDate: v.SecdModifiedDate,
			SecdSectID:       v.SecdSectID,
		}
		listSectionsDetail = append(listSectionsDetail, sectionsDetail)
	}

	if err != nil {
		return nil, &mod.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listSectionsDetail, nil
}

func (per ProgEntityRepository) Gabung(ctx *gin.Context) ([]*mod.Gabung, *mod.ResponseError) {
	// Ambil daftar CurriculumProgramEntity
	progEntities, err := per.GetListProgEntity(ctx)
	if err != nil {
		return nil, err
	}

	// Ambil daftar CurriculumSection
	sections, err := per.GetListSection(ctx)
	if err != nil {
		return nil, err
	}

	// sectionDetails, err := per.GetListSectionDetail(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	// masterCategories, err := per.GetListMasterCategory(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	// Buat map untuk menyimpan CurriculumSection berdasarkan SectProgEntityID
	sectionMap := make(map[int32]*mod.CurriculumSection)
	for _, section := range sections {
		sectionMap[section.SectProgEntityID] = section
	}

	// sectionDetailMap := make(map[int32]*mod.CurriculumSectionDetail)
	// for _, sectionDetail := range sectionDetails {
	// 	sectionDetailMap[sectionDetail.SecdSectID] = sectionDetail
	// }
	// masterCategoryMap := make(map[int32]*mod.MasterCategory)
	// for _, masterCategory := range masterCategories {
	// 	masterCategoryMap[masterCategory.CateID] = masterCategory
	// }

	// Gabungkan data dari CurriculumProgramEntity dan CurriculumSection
	gabungList := make([]*mod.Gabung, 0)
	for _, progEntity := range progEntities {
		section, ok := sectionMap[progEntity.ProgEntityID]
		if !ok {
			// Jika tidak ada CurriculumSection yang cocok, lanjutkan ke iterasi selanjutnya
			continue
		}
		// sectionDetail, ok := sectionDetailMap[section.SectID]

		// if !ok {
		// 	// Jika tidak ada CurriculumSectionDetail yang cocok, lanjutkan ke iterasi selanjutnya
		// 	continue
		// }
		// masterCategory, ok := masterCategoryMap[progEntity.ProgCateID]
		// if !ok {
		// 	// Jika tidak ada MasterCategory yang cocok, lanjutkan ke iterasi selanjutnya
		// 	continue
		// }

		gabung := &mod.Gabung{
			CurriculumProgramEntity: *progEntity,
			CurriculumSection:       *section,
			// CurriculumSectionDetail: *sectionDetail,
			// MasterCategory:          *masterCategory,
		}

		gabungList = append(gabungList, gabung)
	}

	return gabungList, nil
}

func (per ProgEntityRepository) GetProgEntity(ctx *gin.Context, id int64) (*mod.CurriculumProgramEntity, *mod.ResponseError) {

	store := dbContext.New(per.dbHandler)
	progentity, err := store.Getprogram_entity(ctx, int32(id))

	if err != nil {
		return nil, &mod.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &progentity, nil
}

func (per ProgEntityRepository) GetSection(ctx *gin.Context, id int64) (*[]mod.CurriculumSectionGet, *mod.ResponseError) {

	store := dbContext.New(per.dbHandler)
	section, err := store.Getsections(ctx, int32(id))

	if err != nil {
		return nil, &mod.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &section, nil
}

func (per ProgEntityRepository) GetGabung(ctx *gin.Context, id int64) (*[]mod.GetGabung, *mod.ResponseError) {
	// Ambil data CurriculumProgramEntity
	progEntity, err := per.GetProgEntity(ctx, id)
	if err != nil {
		return nil, err
	}

	// Ambil data CurriculumSection
	sections, err := per.GetSection(ctx, id)
	if err != nil {
		return nil, err
	}

	// Buat objek GetGabung
	getGabung := &mod.GetGabung{
		CurriculumProgramEntity: *progEntity,
		CurriculumSectionGet:    *sections,
	}

	// Masukkan objek GetGabung ke dalam slice, karena GetGabung berada dalam slice
	gabungSlice := []mod.GetGabung{*getGabung}

	return &gabungSlice, nil
}

func (per ProgEntityRepository) CreateProgEntity(ctx *gin.Context, progentityParams *db.Createprogram_entityParams) (*mod.CurriculumProgramEntity, *mod.ResponseError) {

	store := dbContext.New(per.dbHandler)
	category, err := store.Createprogram_entity(ctx, *progentityParams)

	if err != nil {
		return nil, &mod.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return category, nil
}
func (per ProgEntityRepository) CreateSection(ctx *gin.Context, sectionsParams *db.CreatesectionsParams) (*mod.CurriculumSection, *mod.ResponseError) {
	store := dbContext.New(per.dbHandler)
	section, err := store.Createsections(ctx, *sectionsParams)
	if err != nil {
		return nil, &mod.ResponseError{
			Message: err.Message,
			Status:  http.StatusInternalServerError,
		}
	}
	return section, nil
}

func (per ProgEntityRepository) CreateGabung(ctx *gin.Context, gabungParams *db.CreateGabung) (*mod.CreateGabung, *mod.ResponseError) {
	// Buat CurriculumProgramEntity
	progEntity, err := per.CreateProgEntity(ctx, &gabungParams.Createprogram_entityParams)
	if err != nil {
		return nil, err
	}

	section, err := per.CreateSection(ctx, &gabungParams.CreatesectionsParams)
	if err != nil {
		return nil, err
	}

	// Buat Gabung
	gabung := &mod.CreateGabung{
		CurriculumProgramEntity: *progEntity,
		CurriculumSection:       *section,
	}

	return gabung, nil
}

func (per ProgEntityRepository) UpdateProgEntity(ctx *gin.Context, progentityParams *db.Createprogram_entityParams) *mod.ResponseError {

	store := dbContext.New(per.dbHandler)
	err := store.Updateprogram_entity(ctx, *progentityParams)

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

func (per ProgEntityRepository) DeleteProgEntity(ctx *gin.Context, id int64) *mod.ResponseError {

	store := dbContext.New(per.dbHandler)
	err := store.Deleteprogram_entity(ctx, int32(id))

	if err != nil {
		return &mod.ResponseError{
			Message: "error when update",
			Status:  http.StatusInternalServerError,
		}
	}
	return &mod.ResponseError{
		Message: "data has been deleted",
		Status:  http.StatusOK,
	}
}
