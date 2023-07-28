package bootcampRepository

import "database/sql"

type RepositoryManager struct {
	BatchRepository
}

// constructor
func NewRepositoryManager(dbHandler *sql.DB) *RepositoryManager {
	return &RepositoryManager{
		*NewBatchRepository(dbHandler),
	}
}
