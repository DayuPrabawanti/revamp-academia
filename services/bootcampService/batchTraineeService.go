package bootcampService

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/bootcampRepository"
	"github.com/gin-gonic/gin"
)

type BatchTraineeService struct {
	batchTraineeRepository *bootcampRepository.BatchTraineeRepository
}

func NewBatchTraineeService(batchTraineeRepository *bootcampRepository.BatchTraineeRepository) *BatchTraineeService {
	return &BatchTraineeService{
		batchTraineeRepository: batchTraineeRepository,
	}
}

func (btes BatchTraineeService) GetListBatchTrainee(ctx *gin.Context) ([]*models.BootcampBatchTrainee, *models.ResponseError) {
	return btes.batchTraineeRepository.GetListBatchTrainee(ctx)
}
