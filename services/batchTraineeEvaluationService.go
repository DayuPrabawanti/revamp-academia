package services

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories"
	"github.com/gin-gonic/gin"
)

type BatchTraineeService struct {
	batchTraineeRepository *repositories.BatchTraineeRepository
}

func NewBatchTraineeService(batchTraineeRepository *repositories.BatchTraineeRepository) *BatchTraineeService {
	return &BatchTraineeService{
		batchTraineeRepository: batchTraineeRepository,
	}
}

func (btes BatchTraineeService) GetListBatchTrainee(ctx *gin.Context) ([]*models.BootcampBatchTrainee, *models.ResponseError) {
	return btes.batchTraineeRepository.GetListBatchTrainee(ctx)
}
