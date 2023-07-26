package salesServices

import (
	"codeid.revampacademy/repositories/salesRepositories"
)

type ServiceManager struct {
	CartItemsService
	ServiceMock
	FintechService
	SpecialOfferService
	EducationService
	ServiceMock3
	ServiceMock2
	ServiceMock4
	ServiceMock8
}

// constructor
func NewServiceManager(repoMgr *salesRepositories.RepositoryManager) *ServiceManager {
	return &ServiceManager{
		CartItemsService:    *NewCartItemsRepository(&repoMgr.CartItemsRepository),
		ServiceMock:         *NewServiceMock(&repoMgr.RepositoryMock),
		FintechService:      *NewFintechService(&repoMgr.FintechRepository),
		SpecialOfferService: *NewSpecialOfferService(&repoMgr.SpecialOfferRepository),
		EducationService:    *NewEducationService(&repoMgr.EducationRepository),
		ServiceMock3:        *NewMockupApplyService(&repoMgr.RepoMockup3),
		ServiceMock2:        *NewMockupApplyService2(&repoMgr.RepoMockup2),
		ServiceMock4:        *NewMockupApplyService4(&repoMgr.RepoMockup4),
		ServiceMock8:        *NewMockupApplyService8(&repoMgr.RepoMockup8),
	}
}
