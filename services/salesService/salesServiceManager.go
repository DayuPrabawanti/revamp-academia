package salesService

import salesrepositories "codeid.revampacademy/repositories/salesRepositories"

type ServiceManager struct {
	ServiceMock
	ServiceMock3
}

func NewServiceManager(repoMgr *salesrepositories.RepositoryManager) *ServiceManager {
	return &ServiceManager{
		ServiceMock:  *NewServiceMock(&repoMgr.RepositoryMock),
		ServiceMock3: *NewMockupApplyService(&repoMgr.RepoMockup3),
	}
}
