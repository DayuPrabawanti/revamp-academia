package bootcampService

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/bootcampRepository"
	"github.com/gin-gonic/gin"
)

type InstructorProgramService struct {
	instructorProgramRepository *bootcampRepository.InstructorProgramRepository
}

func NewInstructorProgramService(instructorProgramRepository *bootcampRepository.InstructorProgramRepository) *InstructorProgramService {
	return &InstructorProgramService{
		instructorProgramRepository: instructorProgramRepository,
	}
}

func (ips InstructorProgramService) GetListInstructorProgram(ctx *gin.Context) ([]*models.BootcampInstructorProgram, *models.ResponseError) {
	return ips.instructorProgramRepository.GetListInstructorPrograms(ctx)
}
