package salesRepositories

import "database/sql"

type RepositoryManager struct {
	CartItemsRepository
	RepositoryMock
	FintechRepository
	SpecialOfferRepository
	EducationRepository
	RepoMockup3
}

// constructor
func NewRepositoryManager(dbHandler *sql.DB) *RepositoryManager {
	return &RepositoryManager{
		*NewCartItemsRepository(dbHandler),
		*NewRepositoryMock(dbHandler),
		*NewFintechRepository(dbHandler),
		*NewSpecialOfferRepository(dbHandler),
		*NewEducationRepository(dbHandler),
		*NewMockupApplyRepo(dbHandler),
	}
}
