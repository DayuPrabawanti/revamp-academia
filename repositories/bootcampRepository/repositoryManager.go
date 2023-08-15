package bootcampRepository

import "database/sql"

type RepositoryManager struct {
	BatchRepository
	BatchTraineeEvaluationRepository
	BatchTraineeRepository
	InstructorProgramRepository
	ProgramApplyRepository
	ProgramApplyProgressRepository
	BootcampBatchEvaluationRepository
	GabungRepository
	ProgAppllyRepository
}

// constructor
func NewRepositoryManager(dbHandler *sql.DB) *RepositoryManager {
	return &RepositoryManager{
		*NewBatchRepository(dbHandler),
		*NewBatchTraineeEvaluationRepository(dbHandler),
		*NewBatchTraineeRepository(dbHandler),
		*NewInstructorProgramRepository(dbHandler),
		*NewProgramApplyRepository(dbHandler),
		*NewProgramApplyProgressRepository(dbHandler),
		*NewBootcampBatchEvaluationRepository(dbHandler),
		*NewGabungRepository(dbHandler),
		*NewProgAppllyRepository(dbHandler),
	}
}
