package salesservice

import salesrepositories "codeid.revampacademy/repositories/salesRepositories"

type ServiceManager struct {
	ServiceMock
	ServiceMock3
	ServiceMock6
	ServiceMock7
	ServiceMockup8
}

func NewServiceManager(repoMgr *salesrepositories.RepositoryManager) *ServiceManager {
	return &ServiceManager{
		ServiceMock:    *NewServiceMock(&repoMgr.RepositoryMock),
		ServiceMock3:   *NewMockupApplyService(&repoMgr.RepoMockup3),
		ServiceMock6:   *NewServiceShoppingCart1(&repoMgr.RepoMockup6),
		ServiceMock7:   *NewServiceShoppingCart2(&repoMgr.RepoMock7),
		ServiceMockup8: *NewServiceShoppingCart3(&repoMgr.RepoMockup8),
	}
}
