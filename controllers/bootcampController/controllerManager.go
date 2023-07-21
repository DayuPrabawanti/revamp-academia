package bootcampController

import "codeid.revampacademy/services/bootcampService"

type ControllerManager struct {
	BatchController
	BatchTraineeController
	BatchTraineeEvaluationController
	InstructorProgramController
	ProgramApplyController
	ProgramApplyProgressController
	BootcampBatchEvaluationController
}

// constructor
func NewControllerManager(serviceMgr *bootcampService.ServiceManager) *ControllerManager {
	return &ControllerManager{
		*NewBatchController(&serviceMgr.BatchService),
		*NewBatchTraineeController(&serviceMgr.BatchTraineeService),
		*NewBatchTraineeEvaluationController(&serviceMgr.BatchTraineeEvaluationService),
		*NewInstructorProgramController(&serviceMgr.InstructorProgramService),
		*NewProgramApplyController(&serviceMgr.ProgramApplyService),
		*NewProgramApplyProgressController(&serviceMgr.ProgramApplyProgressService),
		*NewBootcampBatchEvaluationController(&serviceMgr.BootcampBatchEvaluationService),
	}
}
