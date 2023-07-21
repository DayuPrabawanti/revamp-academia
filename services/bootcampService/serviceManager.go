package bootcampService

import "codeid.revampacademy/repositories/bootcampRepository"

type ServiceManager struct {
	BatchService
}

// constructor
func NewServiceManager(repoMgr *bootcampRepository.RepositoryManager) *ServiceManager {
	return &ServiceManager{
		BatchService: *NewBatchService(&repoMgr.BatchRepository),
		// ProductService: *NewProductService(&repoMgr.ProductRepository),
	}
}
