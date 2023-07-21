package salesrepositories

import "database/sql"

type RepositoryManager struct {
	RepositoryMock
	RepoMockup3
}

func NewRepositoryManager(dbHandler *sql.DB) *RepositoryManager {
	return &RepositoryManager{
		*NewRepositoryMock(dbHandler),
		*NewMockupApplyRepo(dbHandler),
	}
}
