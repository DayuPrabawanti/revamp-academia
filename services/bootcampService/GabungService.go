package bootcampService

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/bootcampRepository"
	"github.com/gin-gonic/gin"
)

type Gabung struct {
	gabungRepository *bootcampRepository.GabungRepository
}

func NewGabung(gabungRepository *bootcampRepository.GabungRepository) *Gabung {
	return &Gabung{
		gabungRepository: gabungRepository,
	}
}

func (gb Gabung) GetListGabung(ctx *gin.Context) ([]*models.Gabung, *models.ResponseError) {
	return gb.gabungRepository.GetListGabung(ctx)
}
