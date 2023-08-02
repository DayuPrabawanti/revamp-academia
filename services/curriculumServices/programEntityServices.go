package curriculumServices

import (
	"net/http"

	mod "codeid.revampacademy/models"
	"codeid.revampacademy/models/features"
	repo "codeid.revampacademy/repositories/curriculumRepositories"
	db "codeid.revampacademy/repositories/curriculumRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type ProgEntityService struct {
	repositoryManager repo.RepositoryManager
}

func NewProgEntityService(repoMgr *repo.RepositoryManager) *ProgEntityService {
	return &ProgEntityService{
		repositoryManager: *repoMgr,
	}
}

func (cs ProgEntityService) GetListProgEntity(ctx *gin.Context, metadata *features.Metadata) ([]*mod.CurriculumProgramEntity, *mod.ResponseError) {
	return cs.repositoryManager.ProgEntityRepository.GetListProgEntity(ctx, metadata)
}

func (cs ProgEntityService) GetListSection(ctx *gin.Context, metadata *features.Metadata) ([]*mod.CurriculumSection, *mod.ResponseError) {
	return cs.repositoryManager.ProgEntityRepository.GetListSection(ctx, metadata)
}
func (cs ProgEntityService) GetListMasterCategory(ctx *gin.Context, metadata *features.Metadata) ([]*mod.MasterCategory, *mod.ResponseError) {
	return cs.repositoryManager.ProgEntityRepository.GetListMasterCategory(ctx, metadata)
}
func (cs ProgEntityService) GetListSectionDetail(ctx *gin.Context, metadata *features.Metadata) ([]*mod.CurriculumSectionDetail, *mod.ResponseError) {
	return cs.repositoryManager.ProgEntityRepository.GetListSectionDetail(ctx, metadata)
}

func (cs ProgEntityService) GetProgEntity(ctx *gin.Context, id int64) (*mod.CurriculumProgramEntity, *mod.ResponseError) {
	return cs.repositoryManager.ProgEntityRepository.GetProgEntity(ctx, id)
}

func (cs ProgEntityService) GetSection(ctx *gin.Context, id int64) (*[]mod.CurriculumSectionGet, *mod.ResponseError) {
	return cs.repositoryManager.ProgEntityRepository.GetSection(ctx, id)
}

func (cs ProgEntityService) GetGabung(ctx *gin.Context, id int64) (*[]mod.GetGabung, *mod.ResponseError) {
	return cs.repositoryManager.ProgEntityRepository.GetGabung(ctx, id)
}

func (cs ProgEntityService) GetCategory(ctx *gin.Context, id int64) (*[]mod.MasterCategory, *mod.ResponseError) {
	return cs.repositoryManager.ProgEntityRepository.GetCategory(ctx, id)
}

func (cs ProgEntityService) CreateProgEntity(ctx *gin.Context, progentityParams *db.Createprogram_entityParams) (*mod.CurriculumProgramEntity, *mod.ResponseError) {
	responseErr := validateProgEntity(progentityParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return cs.repositoryManager.ProgEntityRepository.CreateProgEntity(ctx, progentityParams)
}

func (cs ProgEntityService) CreateSections(ctx *gin.Context, sectionsParams *db.CreatesectionsParams) (*mod.CurriculumSection, *mod.ResponseError) {
	responseErr := validateSection(sectionsParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return cs.repositoryManager.ProgEntityRepository.CreateSection(ctx, sectionsParams)
}

func (cs ProgEntityService) CreateGabung(ctx *gin.Context, gabungParams *db.CreateGabung) (*mod.Gabung, *mod.ResponseError) {

	return cs.repositoryManager.ProgEntityRepository.CreateGabung(ctx, gabungParams)
}

func (cs ProgEntityService) UpdateProgEntity(ctx *gin.Context, progentityParams *db.Createprogram_entityParams, id int64) *mod.ResponseError {
	responseErr := validateProgEntity(progentityParams)
	if responseErr != nil {
		return responseErr
	}

	return cs.repositoryManager.UpdateProgEntity(ctx, progentityParams)
}

func (cs ProgEntityService) UpdateGabung(ctx *gin.Context, progentityParams *db.Createprogram_entityParams, id int64) *mod.ResponseError {
	responseErr := validateProgEntity(progentityParams)
	if responseErr != nil {
		return responseErr
	}

	return cs.repositoryManager.UpdateProgEntity(ctx, progentityParams)
}

func (cs ProgEntityService) DeleteProgEntity(ctx *gin.Context, id int64) *mod.ResponseError {
	return cs.repositoryManager.DeleteProgEntity(ctx, id)
}

func (cs ProgEntityService) Gabung(ctx *gin.Context, metadata *features.Metadata) ([]*mod.Gabung, *mod.ResponseError) {
	return cs.repositoryManager.ProgEntityRepository.Gabung(ctx, metadata)

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

// func (Pe ProgEntityService) CreateGroupDto(ctx *gin.Context, progEntityWithSectionDto *mod.CreateGroupDto) (*models.CurriculumProgramEntity, *models.ResponseError) {

// 	err := repo.BeginTransaction(&Pe.repositoryManager)
// 	if err != nil {
// 		return nil, &models.ResponseError{
// 			Message: "Failed to start transaction",
// 			Status:  http.StatusBadRequest,
// 		}
// 	}
// 	//first query statement
// 	response, responseErr := Pe.CreateProgEntity(ctx, (*dbContext.Createprogram_entityParams)(&progEntityWithSectionDto.CreateProgEntityDto))
// 	if responseErr != nil {
// 		repo.RollbackTransaction(&Pe.repositoryManager)
// 		return nil, responseErr
// 	}
// 	//second query statement
// 	responseErr = Pe.DeleteProgEntity(ctx, int64(response.ProgEntityID))
// 	if responseErr != nil {
// 		//when delete not succeed, transaction will rollback
// 		repo.RollbackTransaction(&Pe.repositoryManager)
// 		return nil, responseErr
// 	}
// 	// if all statement ok, transaction will commit/save to db
// 	repo.CommitTransaction(&Pe.repositoryManager)

// 	return nil, &models.ResponseError{
// 		Message: "Data has been created",
// 		Status:  http.StatusOK,
// 	}
// }
