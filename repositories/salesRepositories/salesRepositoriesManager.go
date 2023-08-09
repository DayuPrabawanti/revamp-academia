package salesRepositories

import "database/sql"

type RepositoryManager struct {
	RepositoryMock
	RepoMockup3
	RepoMockup6
	RepoMock7
	RepoMockup8
	RepoMockup4
}

func NewRepositoryManager(dbHandler *sql.DB) *RepositoryManager {
	return &RepositoryManager{
		*NewRepositoryMock(dbHandler),
		*NewMockupApplyRepo(dbHandler),
		*NewRepoShoppingCart1(dbHandler),
		*NewRepoShoppingCart2(dbHandler),
		*NewRepoShoppingCart3(dbHandler),
		*NewMockupApplyRepo4(dbHandler),
	}
}
