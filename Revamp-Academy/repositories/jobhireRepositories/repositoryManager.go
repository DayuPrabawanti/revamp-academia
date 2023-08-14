package jobhireRepositories

import "database/sql"

type RepositoryManager struct {
	// CategoryRepo
	JobHirePostRepo
	// MasterRepo
}

func NewRepositoryManager(dbHandler *sql.DB) *RepositoryManager {
	return &RepositoryManager{
		// *NewCategoryRepo(dbHandler),
		*NewJobPostRepo(dbHandler),
		// *NewMasterRepo(dbHandler),
	}
}
