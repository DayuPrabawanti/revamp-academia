package services

import (
	"codeid.revampacademy/models"
	"github.com/gin-gonic/gin"
)

func (pe ProgramEntityService) GroupList(ctx *gin.Context) ([]*models.Group, *models.ResponseError) {
	return pe.programEntityRepository.GroupList(ctx)
}
