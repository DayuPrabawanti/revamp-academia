package services

import "codeid.revampacademy/repositories"

type ServiceManager struct {
	JobPostService
	MasterIndustryService
}

// Constructor
func NewServiceManager(repoMgr *repositories.RepositoriesManager) *ServiceManager {
	return &ServiceManager{
		JobPostService: *NewJobPostService(&repoMgr.JobPostRepository),
		MasterIndustryService: *NewMasterIndustryService(&repoMgr.MasterIndustryRepository),

	}
}
