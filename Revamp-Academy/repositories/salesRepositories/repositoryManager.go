package salesRepositories

import "database/sql"

type RepositoryManager struct {
	CartItemsRepository
	RepositoryMock
	FintechRepository
	SpecialOfferRepository
	EducationRepository
	RepoMockup3
	RepoMockup2
	RepoMockup4
	RepoMockup8
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
		*NewMockupApplyRepo2(dbHandler),
		*NewMockupApplyRepo4(dbHandler),
		*NewMockupApplyRepo8(dbHandler),
	}
}
