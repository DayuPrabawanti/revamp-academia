package jobhireService

import "codeid.revampacademy/repositories/jobhireRepositories"

type ServiceManager struct {
	// CategoryService
	JobService
	// MasterService
}

func NewServiceManager(repositoryManager *jobhireRepositories.RepositoryManager) *ServiceManager {
	return &ServiceManager{
		// CategoryService: *NewCategoryService(&repositoryManager.CategoryRepo),
		JobService: *NewJobService(repositoryManager),
		// MasterService:   *NewMasterService(&repositoryManager.MasterRepo),
	}
}
