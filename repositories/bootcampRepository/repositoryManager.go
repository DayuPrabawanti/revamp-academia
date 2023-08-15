package bootcampRepository

import "database/sql"

type RepositoryManager struct {
	BatchRepository
	BootcampBatchEvaluationRepository
	EvaluationCandidateRepository
}

// constructor
func NewRepositoryManager(dbHandler *sql.DB) *RepositoryManager {
	return &RepositoryManager{
		*NewBatchRepository(dbHandler),
		*NewBootcampBatchEvaluationRepository(dbHandler),
		*NewEvaluationCandidateRepository(dbHandler),
	}
}
