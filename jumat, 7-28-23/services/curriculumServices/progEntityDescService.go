package curriculumServices

import (
	mod "codeid.revampacademy/models"
	repo "codeid.revampacademy/repositories/curriculumRepositories"
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
