package curriculum

import "database/sql"

type RepositoryManager struct {
	ProgEntityRepository
}

func NewRepositoryManager(dbHandler *sql.DB) *RepositoryManager {
	return &RepositoryManager{
		*NewProgEntityRepository(dbHandler),
	}
}
