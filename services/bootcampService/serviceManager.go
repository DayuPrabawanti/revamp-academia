package bootcampService

import (
	"codeid.revampacademy/repositories/bootcampRepository"
)

type ServiceManager struct {
	BatchService
	BootcampBatchEvaluationService
	EvaluationCandidateService
}

// constructor
func NewServiceManager(repoMgr *bootcampRepository.RepositoryManager) *ServiceManager {
	return &ServiceManager{
		BatchService:                   *NewBatchService(repoMgr),
		BootcampBatchEvaluationService: *NewBootcampBatchEvaluationService(repoMgr),
		EvaluationCandidateService:     *NewEvaluationCandidateService(repoMgr),
	}
}
