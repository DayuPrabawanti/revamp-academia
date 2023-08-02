package curriculumServices

import (
	"net/http"

	mod "codeid.revampacademy/models"
	"codeid.revampacademy/models/features"
	repo "codeid.revampacademy/repositories/curriculumRepositories"
	db "codeid.revampacademy/repositories/curriculumRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type ProgEntityDescService struct {
	repositoryManager repo.RepositoryManager
}

func NewProgEntityDescService(repoMgr *repo.RepositoryManager) *ProgEntityDescService {
	return &ProgEntityDescService{
		repositoryManager: *repoMgr,
	}
}

func (ped ProgEntityDescService) GetListProgEntityDesc(ctx *gin.Context, metadata *features.Metadata) ([]*mod.CurriculumProgramEntityDescription, *mod.ResponseError) {
	return ped.repositoryManager.ProgEntityRepository.GetListProgEntityDesc(ctx, metadata)
}

func (ped ProgEntityDescService) GetProgEntityDesc(ctx *gin.Context, id int64) (*mod.CurriculumProgramEntityDescription, *mod.ResponseError) {
	return ped.repositoryManager.ProgEntityRepository.GetProgEntityDesc(ctx, id)
}

func (ped ProgEntityDescService) CreateProgEntityDesc(ctx *gin.Context, progEntityDescParams *db.CreateProgEntityDescParams) (*mod.CurriculumProgramEntityDescription, *mod.ResponseError) {
	responseErr := validateProgEntityDesc(progEntityDescParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return ped.repositoryManager.ProgEntityRepository.CreateProgEntityDesc(ctx, progEntityDescParams)
}

func (ped ProgEntityDescService) UpdateProgEntityDesc(ctx *gin.Context, progEntityDescParams *db.CreateProgEntityDescParams, id int64) *mod.ResponseError {
	responseErr := validateProgEntityDesc(progEntityDescParams)
	if responseErr != nil {
		return responseErr
	}

	return ped.repositoryManager.UpdateProgEntityDesc(ctx, (*db.UpdateProgEntityDescParams)(progEntityDescParams))
}

func (ped ProgEntityDescService) DeleteProgEntityDesc(ctx *gin.Context, id int64) *mod.ResponseError {
	return ped.repositoryManager.DeleteProgEntityDesc(ctx, id)
}

func validateProgEntityDesc(progEntityDescParams *db.CreateProgEntityDescParams) *mod.ResponseError {
	if progEntityDescParams.PredProgEntityID == 0 {
		return &mod.ResponseError{
			Message: "Invalid Program Entity Description Id",
			Status:  http.StatusBadRequest,
		}
	}

	if progEntityDescParams.PredItemLearning.String == "" {
		return &mod.ResponseError{
			Message: "Invalid Program Entity Description Item Learning",
			Status:  http.StatusBadRequest,
		}
	}

	return nil

}
