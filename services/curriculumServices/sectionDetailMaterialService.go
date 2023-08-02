package curriculumServices

import (
	"net/http"

	models "codeid.revampacademy/models"
	"codeid.revampacademy/models/features"
	repo "codeid.revampacademy/repositories/curriculumRepositories"
	dbcontext "codeid.revampacademy/repositories/curriculumRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

// SECTION DETAIL MATERIAL

type SectionDetailMaterialService struct {
	repositoryManager repo.RepositoryManager
}

func NewSectionDetailMaterialService(repoMgr *repo.RepositoryManager) *SectionDetailMaterialService {
	return &SectionDetailMaterialService{
		repositoryManager: *repoMgr,
	}
}

func (sdm SectionDetailMaterialService) GetListSectionDetailMaterial(ctx *gin.Context, metadata *features.Metadata) ([]*models.CurriculumSectionDetailMaterial, *models.ResponseError) {
	return sdm.repositoryManager.ProgEntityRepository.GetListSectionDetailMaterial(ctx)
}

func (sdm SectionDetailMaterialService) GetSectionDetailMaterial(ctx *gin.Context, id int64) (*models.CurriculumSectionDetailMaterial, *models.ResponseError) {
	return sdm.repositoryManager.ProgEntityRepository.GetSectionDetailMaterial(ctx, id)
}

func (sdm SectionDetailMaterialService) CreatesectiondetailMaterial(ctx *gin.Context, sectionDetailMaterialParams *dbcontext.CreatesectionDetailMaterialParams) (*models.CurriculumSectionDetailMaterial, *models.ResponseError) {
	responseErr := validateSectDetMaterial(sectionDetailMaterialParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return sdm.repositoryManager.ProgEntityRepository.CreatesectiondetailMaterial(ctx, sectionDetailMaterialParams)
}

func (sdm SectionDetailMaterialService) UpdateSectionDetailMaterial(ctx *gin.Context, sectionDetailMaterialParams *dbcontext.CreatesectionDetailMaterialParams, id int64) *models.ResponseError {
	responseErr := validateSectDetMaterial(sectionDetailMaterialParams)
	if responseErr != nil {
		return responseErr
	}

	return sdm.repositoryManager.UpdateSectionDetailMaterial(ctx, sectionDetailMaterialParams)
}

func (sdm SectionDetailMaterialService) DeleteSectionDetailMaterial(ctx *gin.Context, id int64) *models.ResponseError {
	return sdm.repositoryManager.DeleteSectionDetailMaterial(ctx, id)
}

func validateSectDetMaterial(sectionDetailMaterialParams *dbcontext.CreatesectionDetailMaterialParams) *models.ResponseError {
	if sectionDetailMaterialParams.SedmID == 0 {
		return &models.ResponseError{
			Message: "Invalid program secd id",
			Status:  http.StatusBadRequest,
		}
	}

	if sectionDetailMaterialParams.SedmFilename == "" {
		return &models.ResponseError{
			Message: "Invalid program secd name",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}
