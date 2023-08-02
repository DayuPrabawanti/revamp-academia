package curriculumServices

import (
	repositories "codeid.revampacademy/repositories/curriculumRepositories"
)

type ServiceManager struct {
	ProgEntityService
	ProgEntityDescService
	ProgReviewService
	SectionDetailMaterialService
}

// constructor
func NewServiceManager(repoMgr *repositories.RepositoryManager) *ServiceManager {
	return &ServiceManager{
		ProgEntityService:            *NewProgEntityService(repoMgr),
		ProgEntityDescService:        *NewProgEntityDescService(repoMgr),
		ProgReviewService:            *NewProgReviewsService(repoMgr),
		SectionDetailMaterialService: *NewSectionDetailMaterialService(repoMgr),
	}
}
