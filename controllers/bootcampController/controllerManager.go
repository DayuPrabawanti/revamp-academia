package bootcampController

import "codeid.revampacademy/services/bootcampService"

type ControllerManager struct {
	BatchController
	BootcampBatchEvaluationController
	EvaluationCandidateController
}

// constructor
func NewControllerManager(serviceMgr *bootcampService.ServiceManager) *ControllerManager {
	return &ControllerManager{
		*NewBatchController(&serviceMgr.BatchService),
		*NewBootcampBatchEvaluationController(&serviceMgr.BootcampBatchEvaluationService),
		*NewEvaluationCandidateController(&serviceMgr.EvaluationCandidateService),
	}
}
