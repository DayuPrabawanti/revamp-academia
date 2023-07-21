package curriculumServices

import (
	"net/http"

	mod "codeid.revampacademy/models"
	repo "codeid.revampacademy/repositories/curriculumRepositories"
	db "codeid.revampacademy/repositories/curriculumRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type ProgEntityService struct {
	progentityRepository *repo.ProgEntityRepository
}

func NewProgEntityService(progentityRepository *repo.ProgEntityRepository) *ProgEntityService {
	return &ProgEntityService{
		progentityRepository: progentityRepository,
	}
}
func (cs ProgEntityService) GetListProgEntity(ctx *gin.Context) ([]*mod.CurriculumProgramEntity, *mod.ResponseError) {
	return cs.progentityRepository.GetListProgEntity(ctx)
}

func (cs ProgEntityService) GetListSection(ctx *gin.Context) ([]*mod.CurriculumSection, *mod.ResponseError) {
	return cs.progentityRepository.GetListSection(ctx)
}
func (cs ProgEntityService) GetListMasterCategory(ctx *gin.Context) ([]*mod.MasterCategory, *mod.ResponseError) {
	return cs.progentityRepository.GetListMasterCategory(ctx)
}
func (cs ProgEntityService) GetListSectionDetail(ctx *gin.Context) ([]*mod.CurriculumSectionDetail, *mod.ResponseError) {
	return cs.progentityRepository.GetListSectionDetail(ctx)
}

func (cs ProgEntityService) GetProgEntity(ctx *gin.Context, id int64) (*mod.CurriculumProgramEntity, *mod.ResponseError) {
	return cs.progentityRepository.GetProgEntity(ctx, id)
}
func (cs ProgEntityService) GetSection(ctx *gin.Context, id int64) (*[]mod.CurriculumSectionGet, *mod.ResponseError) {
	return cs.progentityRepository.GetSection(ctx, id)
}
func (cs ProgEntityService) GetGabung(ctx *gin.Context, id int64) (*[]mod.GetGabung, *mod.ResponseError) {
	return cs.progentityRepository.GetGabung(ctx, id)
}
func (cs ProgEntityService) CreateProgEntity(ctx *gin.Context, progentityParams *db.Createprogram_entityParams) (*mod.CurriculumProgramEntity, *mod.ResponseError) {
	responseErr := validateProgEntity(progentityParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return cs.progentityRepository.CreateProgEntity(ctx, progentityParams)
}

func (cs ProgEntityService) CreateSections(ctx *gin.Context, sectionsParams *db.CreatesectionsParams) (*mod.CurriculumSection, *mod.ResponseError) {
	responseErr := validateSection(sectionsParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return cs.progentityRepository.CreateSection(ctx, sectionsParams)
}

func (cs ProgEntityService) CreateGabung(ctx *gin.Context, gabungParams *db.CreateGabung) (*mod.CreateGabung, *mod.ResponseError) {

	return cs.progentityRepository.CreateGabung(ctx, gabungParams)
}
func (cs ProgEntityService) UpdateProgEntity(ctx *gin.Context, progentityParams *db.Createprogram_entityParams, id int64) *mod.ResponseError {
	responseErr := validateProgEntity(progentityParams)
	if responseErr != nil {
		return responseErr
	}

	return cs.progentityRepository.UpdateProgEntity(ctx, progentityParams)
}

// func (cs ProgEntityService) UpdateSection(ctx *gin.Context, sectionParams *db.CreatesectionsParams, id int64) *mod.ResponseError {
// 	responseErr := validateSection(sectionParams)
// 	if responseErr != nil {
// 		return responseErr
// 	}

//		return cs.progentityRepository.UpdateSection(ctx, sectionParams)
//	}
func (cs ProgEntityService) UpdateGabung(ctx *gin.Context, progentityParams *db.Createprogram_entityParams, id int64) *mod.ResponseError {
	responseErr := validateProgEntity(progentityParams)
	if responseErr != nil {
		return responseErr
	}

	return cs.progentityRepository.UpdateProgEntity(ctx, progentityParams)
}

func (cs ProgEntityService) DeleteProgEntity(ctx *gin.Context, id int64) *mod.ResponseError {
	return cs.progentityRepository.DeleteProgEntity(ctx, id)
}
func (cs ProgEntityService) Gabung(ctx *gin.Context) ([]*mod.Gabung, *mod.ResponseError) {
	return cs.progentityRepository.Gabung(ctx)

}

func validateProgEntity(progentityParams *db.Createprogram_entityParams) *mod.ResponseError {
	if progentityParams.ProgEntityID == 0 {
		return &mod.ResponseError{
			Message: "Invalid Program Entity id",
			Status:  http.StatusBadRequest,
		}
	}

	if progentityParams.ProgTitle == "" {
		return &mod.ResponseError{
			Message: "Invalid Program Entity Title",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}
func validateSection(sectionParams *db.CreatesectionsParams) *mod.ResponseError {
	if sectionParams.SectID == 0 {
		return &mod.ResponseError{
			Message: "Invalid Section id",
			Status:  http.StatusBadRequest,
		}
	}

	if sectionParams.SectTitle == "" {
		return &mod.ResponseError{
			Message: "Invalid Section Title",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}
