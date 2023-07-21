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
		ProgEntityService:            *NewProgEntityService(&repoMgr.ProgEntityRepository),
		ProgEntityDescService:        *NewProgEntityDescService(&repoMgr.ProgEntityDescRepository),
		ProgReviewService:            *NewProgReviewsService(&repoMgr.ProgReviewsRepository),
		SectionDetailMaterialService: *NewSectionDetailMaterialService(&repoMgr.SectionDetailMaterialRepository),
	}
}
