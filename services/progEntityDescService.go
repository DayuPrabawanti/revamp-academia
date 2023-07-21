package services

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories"
	"github.com/gin-gonic/gin"
)

type ProgEntityDescService struct {
	progEntityDescRepository *repositories.ProgEntityDescRepository
}

func NewProgEntityDescService(progEntityDescRepository *repositories.ProgEntityDescRepository) *ProgEntityDescService {
	return &ProgEntityDescService{
		progEntityDescRepository: progEntityDescRepository,
	}
}

func (ped ProgEntityDescService) GetListProgEntityDesc(ctx *gin.Context) ([]*models.CurriculumProgramEntityDescription, *models.ResponseError) {
	return ped.progEntityDescRepository.GetListProgEntityDesc(ctx)
}

func (ped ProgEntityDescService) GetProgEntityDesc(ctx *gin.Context, id int64) (*models.CurriculumProgramEntityDescription, *models.ResponseError) {
	return ped.progEntityDescRepository.GetProgEntityDesc(ctx, id)
}
