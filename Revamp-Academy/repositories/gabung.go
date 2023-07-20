package repositories

import (
	"codeid.revampacademy/models"
	"github.com/gin-gonic/gin"
)

type GabungRepository struct {
}

func (per ProgEntityRepository) Gabung(ctx *gin.Context) ([]*models.Gabung, *models.ResponseError) {
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

	sectionDetails, err := per.GetListSectionDetail(ctx)
	if err != nil {
		return nil, err
	}

	masterCategories, err := per.GetListMasterCategory(ctx)
	if err != nil {
		return nil, err
	}

	// Buat map untuk menyimpan CurriculumSection berdasarkan SectProgEntityID
	sectionMap := make(map[int32]*mod.CurriculumSection)
	for _, section := range sections {
		sectionMap[section.SectProgEntityID] = section
	}

	sectionDetailMap := make(map[int32]*models.CurriculumSectionDetail)
	for _, sectionDetail := range sectionDetails {
		sectionDetailMap[sectionDetail.SecdSectID] = sectionDetail
	}
	masterCategoryMap := make(map[int32]*models.MasterCategory)
	for _, masterCategory := range masterCategories {
		masterCategoryMap[masterCategory.CateID] = masterCategory
	}

	// Gabungkan data dari CurriculumProgramEntity dan CurriculumSection
	gabungList := make([]*models.Gabung, 0)
	for _, progEntity := range progEntities {
		section, ok := sectionMap[progEntity.ProgEntityID]
		if !ok {
			// Jika tidak ada CurriculumSection yang cocok, lanjutkan ke iterasi selanjutnya
			continue
		}
		sectionDetail, ok := sectionDetailMap[section.SectID]

		if !ok {
			// Jika tidak ada CurriculumSectionDetail yang cocok, lanjutkan ke iterasi selanjutnya
			continue
		}
		masterCategory, ok := masterCategoryMap[progEntity.ProgCateID]
		if !ok {
			// Jika tidak ada MasterCategory yang cocok, lanjutkan ke iterasi selanjutnya
			continue
		}

		gabung := &mod.Gabung{
			CurriculumProgramEntity: *progEntity,
			CurriculumSection:       *section,
			CurriculumSectionDetail: *sectionDetail,
			MasterCategory:          *masterCategory,
		}

		gabungList = append(gabungList, gabung)
	}

	return gabungList, nil
}
