package curriculumServices

import (
	"net/http"

	mod "codeid.revampacademy/models"
	repo "codeid.revampacademy/repositories/curriculumRepositories"
	db "codeid.revampacademy/repositories/curriculumRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type ProgEntityDescService struct {
	progEntityDescRepository *repo.ProgEntityDescRepository
}

func NewProgEntityDescService(progEntityDescRepository *repo.ProgEntityDescRepository) *ProgEntityDescService {
	return &ProgEntityDescService{
		progEntityDescRepository: progEntityDescRepository,
	}
}

func (ped ProgEntityDescService) GetListProgEntityDesc(ctx *gin.Context) ([]*mod.CurriculumProgramEntityDescription, *mod.ResponseError) {
	return ped.progEntityDescRepository.GetListProgEntityDesc(ctx)
}

func (ped ProgEntityDescService) GetProgEntityDesc(ctx *gin.Context, id int64) (*mod.CurriculumProgramEntityDescription, *mod.ResponseError) {
	return ped.progEntityDescRepository.GetProgEntityDesc(ctx, id)
}

func (ped ProgEntityDescService) CreateProgEntityDesc(ctx *gin.Context, progEntityDescParams *db.CreateProgEntityDescParams) (*mod.CurriculumProgramEntityDescription, *mod.ResponseError) {
	responseErr := validateProgEntityDesc(progEntityDescParams)
	if responseErr != nil {
		return nil, responseErr
	}

	return ped.progEntityDescRepository.CreateProgEntityDesc(ctx, progEntityDescParams)
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
