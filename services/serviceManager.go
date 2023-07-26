package services

import "codeid.revampacademy/repositories"

type ServiceManager struct {
	JobPostService
	MasterIndustryService
	JobPostingService
	ApplyProfService
}

// Constructor
func NewServiceManager(repoMgr *repositories.RepositoriesManager) *ServiceManager {
	return &ServiceManager{
		JobPostService:        *NewJobPostService(&repoMgr.JobPostRepository),
		MasterIndustryService: *NewMasterIndustryService(&repoMgr.MasterIndustryRepository),
		JobPostingService:     *NewJobPostingService(&repoMgr.JobPostingRepository),
		ApplyProfService:      *NewApplyProfService(&repoMgr.ApplyProfRepository),
	}
}
