package services

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories"
	"github.com/gin-gonic/gin"
)

type ApplyProfService struct {
	ApplyProfRepository *repositories.ApplyProfRepository
}

func NewApplyProfService(ApplyProfRepository *repositories.ApplyProfRepository) *ApplyProfService {
	return &ApplyProfService{
		ApplyProfRepository: ApplyProfRepository,
	}
}

func (apros ApplyProfService) ListApplyProfService(ctx *gin.Context, nama string) ([]*models.ApplyProf, *models.ResponseError) {
	return apros.ApplyProfRepository.ListApplyProfRepo(ctx, nama)
}
