package curriculumServices

import (
	repositories "codeid.revampacademy/repositories/curriculum"
)

type ServiceManager struct {
	ProgEntityService
	//ProductService
}

// constructor
func NewServiceManager(repoMgr *repositories.RepositoryManager) *ServiceManager {
	return &ServiceManager{
		ProgEntityService: *NewProgEntityService(&repoMgr.ProgEntityRepository),
		//ProductService: *NewCategoryService(&repoMgr.ProductRepository),
	}
}
