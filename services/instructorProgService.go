package services

import (
	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories"
	"github.com/gin-gonic/gin"
)

type InstructorProgramService struct {
	instructorProgramRepository *repositories.InstructorProgramRepository
}

func NewInstructorProgramService(instructorProgramRepository *repositories.InstructorProgramRepository) *InstructorProgramService {
	return &InstructorProgramService{
		instructorProgramRepository: instructorProgramRepository,
	}
}

func (ips InstructorProgramService) GetListInstructorProgram(ctx *gin.Context) ([]*models.BootcampInstructorProgram, *models.ResponseError) {
	return ips.instructorProgramRepository.GetListInstructorPrograms(ctx)
}
