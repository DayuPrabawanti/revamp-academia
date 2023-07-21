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
}

// constructor
func NewServiceManager(repoMgr *salesRepositories.RepositoryManager) *ServiceManager {
	return &ServiceManager{
		CartItemsService:    *NewCartItemsRepository(&repoMgr.CartItemsRepository),
		ServiceMock:         *NewServiceMock(&repoMgr.RepositoryMock),
		FintechService:      *NewFintechService(&repoMgr.FintechRepository),
		SpecialOfferService: *NewSpecialOfferService(&repoMgr.SpecialOfferRepository),
		EducationService:    *NewEducationService(&repoMgr.EducationRepository),
	}
}
