package repositories

import "database/sql"

type RepositoriesManager struct {
	JobPostRepository
	MasterIndustryRepository
	JobPostingRepository
	ApplyProfRepository
}

// Constructor
func NewRepositoriesManager(dbHandler *sql.DB) *RepositoriesManager {
	return &RepositoriesManager{
		*NewJobPostRepository(dbHandler),
		*NewMasterIndustryRepository(dbHandler),
		*NewJobPostingRepository(dbHandler),
		*NewApplyProfRepository(dbHandler),
	}
}
