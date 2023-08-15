package bootcampService

import (
	"codeid.revampacademy/repositories/bootcampRepository"
)

type ServiceManager struct {
	BatchService
	BatchTraineeEvaluationService
	BatchTraineeService
	InstructorProgramService
	ProgramApplyService
	ProgramApplyProgressService
	BootcampBatchEvaluationService
	Gabung
	ProgAppllyService
}

// constructor
func NewServiceManager(repoMgr *bootcampRepository.RepositoryManager) *ServiceManager {
	return &ServiceManager{
		*NewBatchService(&repoMgr.BatchRepository),
		*NewBatchTraineeEvaluationService(&repoMgr.BatchTraineeEvaluationRepository),
		*NewBatchTraineeService(&repoMgr.BatchTraineeRepository),
		*NewInstructorProgramService(&repoMgr.InstructorProgramRepository),
		*NewProgramApplyService(&repoMgr.ProgramApplyRepository),
		*NewProgramApplyProgressService(&repoMgr.ProgramApplyProgressRepository),
		*NewBootcampBatchEvaluationService(&repoMgr.BootcampBatchEvaluationRepository),
		*NewGabung(&repoMgr.GabungRepository),
		*NewProgAppllyRepository(&repoMgr.ProgAppllyRepository),
	}
}
